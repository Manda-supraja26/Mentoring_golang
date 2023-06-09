package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	Name      string
	ID        int
	Salary    float64
	Techstack []string
	Project   string
}

func main() {
	Employes := map[string]Employee{}
	// emp =0
	var emp Employee

	// emp {name "" i 0 sla0 tec [] project ""}

	// map emp1 = {}
	// emp2 ={}
	// emp3 = {}
	// string(1)

	for i := 0; i < 2; i++ {
		name := "emp" + strconv.Itoa(emp.ID)
		Employes[name] = UserInput(emp)
		emp.ID += 1

	}
	fmt.Println(Employes)

}

func UserInput(em Employee) Employee {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("Enter the salary: ")
	salValue, _ := reader.ReadString('\n')

	fmt.Println("Enter the technologies: ")
	techstack, err := reader.ReadSlice('\n')

	for i, v := range strings.Split(string(techstack), " ") {
		em.Techstack = append(em.Techstack, string(v))
		em.Techstack[i] = strings.TrimSpace(string(v))
	}

	fmt.Println("Enter the Project: ")
	project, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	em.Name = strings.TrimSpace(name)
	em.Salary, _ = strconv.ParseFloat(strings.TrimSpace(salValue), 64)
	em.Project = strings.TrimSpace(project)

	return em

}

// UpdateSalary Method to update the salary
func (e *Employee) UpdateSalary(sal float64) {
	e.Salary = sal
}

// UpdateProject Method to update the Project
func (e *Employee) UpdateProject(project string) {
	e.Project = project
}
