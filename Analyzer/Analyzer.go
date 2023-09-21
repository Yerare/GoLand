package Analyzer

type Analyzer struct {
	position string
	salary   float64
	address  string
}

func NewAnalyzer() Analyzer {
	return Analyzer{}
}

func (a *Analyzer) GetPosition() string {
	return a.position
}

func (a *Analyzer) SetPosition(position string) {
	a.position = position
}
