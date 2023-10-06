package intern

type intern struct {
	position string
	salary   float64
	address  string
}

func NewIntern() intern {
	return intern{}
}

func (i *intern) GetPosition() string {
	return i.position
}

func (i *intern) SetPosition(position string) {
	i.position = position
}
