/*
healthcheck_test demonostrates white box testing

The test TestCheckSystemHealthCheck() fakes the implementation of alertAdmin(). Instead of sending
an email, it can validate the notified email and its content.
*/
package whitebox

import (
	"errors"
	"testing"
)

func TestCheckSystemHealthCheck(t *testing.T) {
	// Supply the fake implementation of alertAdmin
	alertAdminSaved := alertAdmin
	var notifiedEmail, notifiedContent string
	alertAdmin = func(email, content string) {
		// results are stored in bounded local variables notifiedEmail and notifiedContent
		notifiedEmail = email
		notifiedContent = content
	}
	// Restore the original implementation after this test
	defer func() { alertAdmin = alertAdminSaved }()

	// Invoke CheckSystemHealth with a simulated error
	dummyError = errors.New("essential procs down")
	CheckSystemHealth()

	// Validate the notified email and its content
	if notifiedEmail != adminEmail {
		t.Errorf("email = %q, want %q", notifiedEmail, adminEmail)
	}
	wantContent := dummyError.Error()
	if notifiedContent != wantContent {
		t.Errorf("content = %q, want %q", notifiedContent, wantContent)
	}
}