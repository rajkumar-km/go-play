/*
Package whitebox demonstrates white box testing in Go

  - Whitebox testing knows the internals of the system and performs testing for its internal units
    and stats. This can improve the testing coverage to ensure that all scenarios are tested.
    Typically the the test package resides in the same package name so that it can access the
    internals.
  - In contrast, Blackbox testing is more robust and focuses only on the outside APIs and test
    with various possible inputs. It requires minimal changes as the software evolves.
  - Let's consider an example for whitebox testing. A system healthcheck send an email to admin
    when there is an issue. But, we don't want to trigger actual emails during the test.
  - Since whitebox testing knows the internals, it can supply a fake function to avoid sending the
    actual emails. But we need to structure the code and separate the email part.
*/
package whitebox

const (
	adminEmail = "admin@example.com"
)

var alertAdmin = func(email string, content string) {
	// Actual implementation to send the email
}
var dummyError error = nil

func CheckSystemHealth() {
	// Implementation of system healthcheck goes here
	var err error = dummyError

	// Send an email on health check failure
	if err != nil {
		alertAdmin(adminEmail, err.Error())
	}
}