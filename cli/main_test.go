package main

import "testing"

func TestGetHello(t *testing.T) {
	want := "Hello, World!"
	got := getHello("World")

	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

// You could also write this & some other following tests as what's called a
// "table test", but I left it explicit here to be more readable
func TestGetWeatherReport(t *testing.T) {
	t.Run("true gets a report", func(t *testing.T) {
		want := "The weather today is perfect for learning a new programming language!"
		got := getWeatherReport(true)
		if want != got {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("false does NOT get a report", func(t *testing.T) {
		want := "No weather report requested -- you must be too busy!"
		got := getWeatherReport(false)
		if want != got {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})
}

func TestGetCLIArgsMessage(t *testing.T) {
	t.Run("no args provided says so", func(t *testing.T) {
		want := "You provided no command-line args besides flags"
		got := getCLIArgsMessage([]string{})
		if want != got {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("args provided tells you the args", func(t *testing.T) {
		want := `You provided 3 command-line arg(s), and they were: ["a" "b" "c"]`
		got := getCLIArgsMessage([]string{"a", "b", "c"})
		if want != got {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})
}

// IRL, try not to hit external URLs in unit tests
func TestHitSomePages(t *testing.T) {
	want := `You hit 2 URL(s), and they responded as follows:
	https://httpbin.org/status/200 -- 200 OK
	https://httpbin.org/status/403 -- 403 Forbidden
`
	got := hitSomePages([]string{"https://httpbin.org/status/200", "https://httpbin.org/status/403"})

	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q", want, got)
	}
}

// As you think about how you would improve/add to these tests, be sure to check
// out this article:
// https://medium.com/swlh/unit-testing-cli-programs-in-go-6275c85af2e7
