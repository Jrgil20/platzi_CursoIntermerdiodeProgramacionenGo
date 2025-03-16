package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

type Manager struct {
	id        int
	name      string
	vacation  bool
	employees []Employee
}

func NewManager(id int, name string, vacation bool, employees []Employee) *Manager {
	return &Manager{
		id:        id,
		name:      name,
		vacation:  vacation,
		employees: employees,
	}
}

func (m *Manager) SetId(id int) {
	m.id = id
}

func (m *Manager) SetName(name string) {
	m.name = name
}

func (m *Manager) GetId() int {
	return m.id
}

func (m *Manager) GetName() string {
	return m.name
}

func (m *Manager) AddEmployee(e Employee) {
	m.employees = append(m.employees, e)
}

func (m *Manager) RemoveEmployee(id int) {
	for i, e := range m.employees {
		if e.id == id {
			m.employees = append(m.employees[:i], m.employees[i+1:]...)
			break
		}
	}
}

func (m *Manager) GetEmployees() []Employee {
	return m.employees
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	// 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	// 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "Name"
	fmt.Printf("%v\n", *e3)
	// 4
	e4 := NewEmployee(1, "Name 2", true)
	fmt.Printf("%v\n", *e4)

	// Usando los métodos Get y Set
	e4.SetId(5)
	e4.SetName("Name 3")
	fmt.Println(e4.GetId())
	fmt.Println(e4.GetName())
}
