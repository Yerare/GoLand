CREATE INDEX IF NOT EXISTS formulas_chapter_idx ON formulas USING GIN (to_tsvector('simple', chapter));
