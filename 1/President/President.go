package President

type President struct {
	position string
	salary   float64
	address  string
}

func NewPresident() President {
	return President{}
}

func (p *President) GetPosition() string {
	return p.position
}

func (p *President) SetPosition(position string) {
	p.position = position
}
