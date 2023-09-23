/*
html_template demonstrates using Go html/template

Similar to text/template, and provides additional security by context-appropriate escaping of
strings in HTML, JavaScript, CSS, and URLs. This would prevent from injection or attack.
*/
package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

// tmpl contains the HTML template for reporting Employee details.
const tmpl = 
`<h1>{{.CompanyName}}</h1><p>{{.TagLine}}</p>
<h4>New Employees joined within three months</h4>
<table>
  <tr>
    <th>Emp ID</th>
	<th>Name</th>
	<th>Programming Skills</th>
	<th>Age</th>
  </tr>
  {{range .Employees}}
  <tr>
    <td>{{.ID}}</td>
	<td>{{.Name | printf "%.64s"}}</td>
	<td>{{.Skills.Programming}}</td>
	<td>{{.DOB | yearsPassed}}</td>
  </tr>
  {{end}}
</table>`

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
		CompanyName template.HTML // allows to embed trusted HTML
		TagLine string // strings are escaped for safety
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
		CompanyName: `<p style="color: #33bb33">CompanyName</p>`, // HTML allowed since the field type is template.HTML
		TagLine: "<i>TagLine</i>", // HTML is escaped here since TagLine is a string field
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
