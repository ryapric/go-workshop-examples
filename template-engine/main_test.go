package main

import "testing"

// Test a render of input data to a DB config in a YAML file
func TestRender(t *testing.T) {
	tplText := `
database:
	address:  {{.database.address}}
	port:     {{.database.port}}
	user:     {{.database.user}}
	password: {{.database.password}}`

	data := map[string]any{
		"database": map[string]any{
			"address":  "10.0.0.1",
			"port":     5432,
			"user":     "admin",
			"password": "password",
		},
	}

	want := `
database:
	address:  10.0.0.1
	port:     5432
	user:     admin
	password: password`
	got := render(tplText, data)
	if want != got {
		t.Errorf("\nwant: %s\ngot:  %s", want, got)
	}
}
