package data

import (
	"context"
	"database/sql"
	"easymath/internal/validator"
	"errors"
	"fmt"
	"time"
)

type Formulas struct {
	ID        int64     `json:"Id"`
	CreatedAt time.Time `json:"Created_at"`
	Chapter   string    `json:"Chapter"`
	Level     string    `json:"Level,omitempty"`
	Withvar   Withvar   `json:"Withvar"`
}

func ValidateFormulas(v *validator.Validator, formula *Formulas) {
	v.Check(formula.Chapter != "", "Chapter", "must be provided")
	v.Check(len(formula.Chapter) <= 500, "Chapter", "must not be more than 500 bytes long")
	v.Check(formula.Withvar != 0, "Withvar", "must be provided")
	v.Check(formula.Withvar > 0, "Withvar", "must be a positive integer")
	v.Check(formula.Level != "", "Level", "must be provided")
	v.Check(len(formula.Level) <= 500, "Level", "must not be more than 500 bytes long")
}

type FormulasModel struct {
	DB *sql.DB
}

func (m FormulasModel) Insert(formulas *Formulas) error {
	query := `
		INSERT INTO formulas (Chapter, Level, Withvar)
		VALUES ($1, $2, $3)
		RETURNING Id, Created_at`
	args := []interface{}{formulas.Chapter, formulas.Level, formulas.Withvar}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&formulas.ID, &formulas.CreatedAt)
}
func (m FormulasModel) Get(id int64) (*Formulas, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT Id, Created_at, Chapter, Level, Withvar
		FROM formulas
		WHERE id = $1`
	var formulas Formulas
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&formulas.ID,
		&formulas.CreatedAt,
		&formulas.Chapter,
		&formulas.Level,
		&formulas.Withvar,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &formulas, nil
}
func (m FormulasModel) Update(formula *Formulas) error {
	query := `
		UPDATE formulas
		SET Chapter = $1, Level = $2, Withvar = $3
		WHERE id = $4 and Level = $5
		RETURNING Level`
	args := []interface{}{
		formula.Chapter,
		formula.Level,
		formula.Withvar,
		formula.ID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&formula.Level)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}
func (m FormulasModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM formulas
		WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
func (m FormulasModel) GetAll(chapter string, level string, filters Filters) ([]*Formulas, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, chapter, level, withvar
		FROM formulas
		WHERE (to_tsvector('simple', chapter) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (level @> $2 OR $2 = '{}')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{chapter, level, filters.limit(), filters.offset()}
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()
	totalRecords := 0
	formulas := []*Formulas{}
	for rows.Next() {
		var formula Formulas
		err := rows.Scan(
			&totalRecords,
			&formula.ID,
			&formula.CreatedAt,
			&formula.Chapter,
			&formula.Level,
			&formula.Withvar,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		formulas = append(formulas, &formula)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return formulas, metadata, nil
}
