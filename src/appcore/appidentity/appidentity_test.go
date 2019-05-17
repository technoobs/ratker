package appidentity

import "testing"

func TestCreateAppIdentity(t *testing.T) {
	app := new(AppIdentity)
	if err := app.CreateAppIdentity(); err != nil {
		t.Errorf("Fail to generate identity string: %s\n", err)
	}
}

func TestGetAppIdentity(t *testing.T) {
	app := new(AppIdentity)
	app.identityString = "test-identity"

	result := app.GetAppIdentity()
	if result != app.identityString {
		t.Errorf("Expected: %s\n Actual: %s\n", app.identityString, result)
	}
}

// TODO: Add test to UpdateAppIdentity function
