/*
text_template demonstrates using Go text/templates

In Go, a template is a string or file containing one or more portions enclosed in double
braces {{ <actions> }}. The actions can be values, expressions, conditional statements or
range loops.
*/
package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

// tmpl contains the text based template for reporting Employee details.
// This could be stored in a .tmpl file instead and parsed using template.ParseFiles()
// Convensions used in this template:
// - The dot (.) refers to the field value from the supplied data.
// - "range" and "end" indicates the loop for each employee
// - "if" and corresponding "end" indicates a conditional statement
// - The pipe | indicates passing the output as input to next function
const tmpl = 
`New Employees joined within three months:
{{range .Employees}}{{if isNewJoinee .JoinDate}}***
EmpID: {{.ID}}
Name: {{.Name | printf "%.64s"}}
Programming Skills: {{.Skills.Programming}}
Age: {{.DOB | yearsPassed | printf "%d years old" }}
{{end}}{{end}}`

// isNewJoinee returns true if the given join date is less than 100 days
func isNewJoinee(t time.Time) bool {
	return int(time.Since(t).Hours() / 24) < 100
}

// yearsPassed calculates the total years passed from the given date
func yearsPassed(t time.Time) int {
	return int(time.Since(t).Hours() / 24 / 365)
}

// template.Must process the errors from Parse() and simplifies it. Also, panics with the
// error so that the issues are identified before running the program
var report = template.Must(
	template.New("employee").Funcs(template.FuncMap{"isNewJoinee": isNewJoinee, "yearsPassed": yearsPassed}).Parse(tmpl),
)

func main() {
	// A company data structure for employees
	type Skill struct {
		Database string
		Programming string
	}
	type Employee struct {
		ID int
		Name string
		DOB time.Time
		JoinDate time.Time
		Skills Skill
	}
	type Company struct {
		Employees []Employee
	}

	// A helper function to prepare Time object for given date
	date := func(d string) time.Time {
		t, _ := time.Parse("2006-01-02", d)
		return t
	}

	// Prepare the input data for report
	cProfile := Skill{Programming: "C/C++/Go", Database: "PostgreSQL"}
	javaProfile := Skill{Programming: "Java", Database: "MySQL"}
	company := Company{
		Employees: []Employee{
			{ID: 1, Name: "joe", DOB: date("1988-01-01"), JoinDate: date("2023-09-01"), Skills: cProfile},
			{ID: 2, Name: "john", DOB: date("1989-02-02"), JoinDate: time.Now().AddDate(0, -1, 0), Skills: javaProfile},
			{ID: 3, Name: "jack", DOB: date("1990-03-03"), JoinDate: time.Now().AddDate(0, -2, -5), Skills: javaProfile},
		},
	}
	
	// Prepare and display the report using template
	if err := report.Execute(os.Stdout, company); err != nil {
		log.Fatal(err)
	}
}