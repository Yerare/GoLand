package Employee

type Employee interface {
	GetPosition()
	SetPosition(position string)
	GetAddress()
	SetAddress(address string)
	GetSalary()
	SetSalary(salary float64)
}
