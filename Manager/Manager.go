package Manager

type Manager struct {
	position string
	salary   float64
	address  string
}

func NewManager() Manager {
	return Manager{}
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}
