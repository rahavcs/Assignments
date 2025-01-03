package main

import "fmt"

//Department names
const (
	HR       = "HR"
	IT       = "IT"
	Finance  = "Finance"
)

//Details about employees
type Employee struct {
	ID        int
	Name      string
	Age       int
	Department string
}

var employees []Employee // To store employee details

// To add an employee
func addEmployee(id int, name string, age int, dept string) string {
	// Checking age
	if age <= 18 {
		return "Age must be above 18"
	}
	// Checking if ID exists
	for _, e := range employees {
		if e.ID == id {
			return "ID must be unique"
		}
	}
	// Add employee to the list
	employees = append(employees, Employee{id, name, age, dept})
	return "Employee added successfully"
}

// Looking for an employee by ID or Name
func searchEmployee(searchTerm string) {

	for _, employee := range employees {
		if fmt.Sprintf("%d", employee.ID) == searchTerm || employee.Name == searchTerm {
			fmt.Printf("Employee found: ID: %d, Name: %s, Age: %d, Department: %s\n", employee.ID, employee.Name, employee.Age, employee.Department)
			return
		}
	}

	fmt.Println("Employee not found")
}

// Counting list of employees
func listAndCountByDepartment(dept string) {
	var result string
	count := 0
	for _, e := range employees {
		if e.Department == dept {
			result += fmt.Sprintf("ID: %d, Name: %s, Age: %d\n", e.ID, e.Name, e.Age)
			count++
		}
	}

	// employees found in department or no employees
	if count == 0 {
		fmt.Println("No employees found in this department")
	} else {
		fmt.Printf("Found %d employee(s) in %s department:\n%s", count, dept, result)
	}
}

func main() {
	// 1. Add Employees
	fmt.Println(addEmployee(1, "Alice", 25, HR))    
	fmt.Println(addEmployee(2, "Bob", 30, IT))      
	fmt.Println(addEmployee(3, "Charlie", 22, Finance)) 
	fmt.Println(addEmployee(1, "Diana", 29, HR))    

	// 2. Search Employees by ID or Name
	searchEmployee("2")  // 2 is ID
	searchEmployee("Alice") // Search by Name
	searchEmployee("David") // David is non existant employee

	// 3. Counting list of employees by department
	fmt.Println("\nListing HR Department Employees:")
	listAndCountByDepartment(HR)

	// 4. Count of Employees in IT
	fmt.Println("\nListing IT Department Employees:")
	listAndCountByDepartment(IT)
}
