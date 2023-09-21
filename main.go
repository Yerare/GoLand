package main

import (
	"fmt"
	"goland/DepartmentHead"
	"goland/Manager"
	"goland/President"
	"goland/intern"
)

func main() {
	m := Manager.NewManager()
	m.SetPosition("HRManager")
	fmt.Println(m.GetPosition())
	i := intern.NewIntern()
	i.SetPosition("Starter")
	fmt.Println(i.GetPosition())
	d := DepartmentHead.NewDepartmentHead()
	d.SetPosition("Boss")
	fmt.Println(d.GetPosition())
	p := President.NewPresident()
	p.SetPosition("Boss x2")
	fmt.Println(p.GetPosition())
	a := President.NewPresident()
	a.SetPosition("Analysis")
	fmt.Println(a.GetPosition())
}
