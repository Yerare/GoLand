package DepartmentHead

type DepartmentHead struct {
	position string
	address  string
	salary   float64
}

func NewDepartmentHead() DepartmentHead {
	return DepartmentHead{}
}

func (d *DepartmentHead) GetPosition() string {
	return d.position
}

func (d *DepartmentHead) SetPosition(position string) {
	d.position = position
}
