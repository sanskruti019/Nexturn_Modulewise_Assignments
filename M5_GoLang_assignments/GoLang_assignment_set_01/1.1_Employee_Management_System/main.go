package main

import (
    "errors"
    "fmt"
    "strings"
)

// Constants for departments
const (
    HR_DEPT  = "HR"
    IT_DEPT  = "IT"
    FIN_DEPT = "FINANCE"
)

// Employee struct to hold employee information
type Employee struct {
    ID         int
    Name       string
    Age        int
    Department string
}

// EmployeeManager handles all employee operations
type EmployeeManager struct {
    employees []Employee
}

// NewEmployeeManager creates a new instance of EmployeeManager
func NewEmployeeManager() *EmployeeManager {
    return &EmployeeManager{
        employees: make([]Employee, 0),
    }
}

// AddEmployee adds a new employee after validation
func (em *EmployeeManager) AddEmployee(id int, name string, age int, department string) error {
    // Validate age
    if age < 18 {
        return errors.New("employee must be at least 18 years old")
    }

    // Validate department
    department = strings.ToUpper(department)
    if department != HR_DEPT && department != IT_DEPT && department != FIN_DEPT {
        return fmt.Errorf("invalid department: %s", department)
    }

    // Check for duplicate ID
    for _, emp := range em.employees {
        if emp.ID == id {
            return fmt.Errorf("employee with ID %d already exists", id)
        }
    }

    // Create and add new employee
    newEmployee := Employee{
        ID:         id,
        Name:       name,
        Age:        age,
        Department: department,
    }

    em.employees = append(em.employees, newEmployee)
    return nil
}

// SearchByID searches for an employee by their ID
func (em *EmployeeManager) SearchByID(id int) (*Employee, error) {
    for i := range em.employees {
        if em.employees[i].ID == id {
            return &em.employees[i], nil
        }
    }
    return nil, fmt.Errorf("employee with ID %d not found", id)
}

// SearchByName searches for an employee by their name
func (em *EmployeeManager) SearchByName(name string) ([]*Employee, error) {
    var found []*Employee
    name = strings.ToLower(name)
    
    for i := range em.employees {
        if strings.Contains(strings.ToLower(em.employees[i].Name), name) {
            found = append(found, &em.employees[i])
        }
    }

    if len(found) == 0 {
        return nil, fmt.Errorf("no employees found with name containing '%s'", name)
    }
    return found, nil
}

// ListByDepartment returns all employees in a given department
func (em *EmployeeManager) ListByDepartment(department string) ([]*Employee, error) {
    department = strings.ToUpper(department)
    var deptEmployees []*Employee

    for i := range em.employees {
        if em.employees[i].Department == department {
            deptEmployees = append(deptEmployees, &em.employees[i])
        }
    }

    if len(deptEmployees) == 0 {
        return nil, fmt.Errorf("no employees found in department %s", department)
    }
    return deptEmployees, nil
}

// CountByDepartment returns the number of employees in a department
func (em *EmployeeManager) CountByDepartment(department string) int {
    department = strings.ToUpper(department)
    count := 0
    
    for _, emp := range em.employees {
        if emp.Department == department {
            count++
        }
    }
    return count
}

func main() {
    // Create new employee manager
    manager := NewEmployeeManager()

    // Example 
    fmt.Println("Adding employees...")
    
    // Add some employees
    errors := []error{
        manager.AddEmployee(1, "Rajesh hadpe", 30, "IT"),
        manager.AddEmployee(2, "Vishwa Ghuge", 25, "HR"),
        manager.AddEmployee(3, "Amar bodke", 35, "IT"),
        manager.AddEmployee(4, "Abhishek bodke", 28, "FINANCE"),
    }

    // Check for errors during addition
    for _, err := range errors {
        if err != nil {
            fmt.Printf("Error adding employee: %v\n", err)
        }
    }

    // Try to add an employee with duplicate ID
    err := manager.AddEmployee(1, "Test User", 20, "IT")
    if err != nil {
        fmt.Printf("Expected error: %v\n", err)
    }

    // Search by ID
    emp, err := manager.SearchByID(2)
    if err != nil {
        fmt.Printf("Search error: %v\n", err)
    } else {
        fmt.Printf("Found employee: %+v\n", *emp)
    }

    // Search by name
    employees, err := manager.SearchByName("john")
    if err != nil {
        fmt.Printf("Search error: %v\n", err)
    } else {
        fmt.Println("Employees found by name:")
        for _, emp := range employees {
            fmt.Printf("%+v\n", *emp)
        }
    }

    // List IT department employees
    itEmployees, err := manager.ListByDepartment("IT")
    if err != nil {
        fmt.Printf("List error: %v\n", err)
    } else {
        fmt.Println("\nIT Department employees:")
        for _, emp := range itEmployees {
            fmt.Printf("%+v\n", *emp)
        }
    }

    // Count employees by department
    fmt.Printf("\nEmployee counts by department:\n")
    fmt.Printf("IT: %d\n", manager.CountByDepartment(IT_DEPT))
    fmt.Printf("HR: %d\n", manager.CountByDepartment(HR_DEPT))
    fmt.Printf("Finance: %d\n", manager.CountByDepartment(FIN_DEPT))
}