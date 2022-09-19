package main

import (
	"flag"
	"fmt"
	"net/http"
)

// The stdlib flag package provides an out-of-the-box way to work with
// command-line flags. Flags are statically-typed by their function names, which
// provides a nice guarantee of what to expect when you parse & process them
var (
	// name will take a string value
	name = flag.String("name", "World", "who to say hello to")
	// weatherReport is a bool flag, where the presence of the flag itself means
	// 'true' (i.e. not providing it means it's false)
	weatherReport = flag.Bool("weatherReport", false, "get a weather report readout")
)

// Flag variables are accessed via their pointers

func getHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func getWeatherReport(shouldReport bool) string {
	var msg string
	if shouldReport {
		msg = "The weather today is perfect for learning a new programming language!"
	} else {
		msg = "No weather report requested -- you must be too busy!"
	}
	return msg
}

func getCLIArgsMessage(args []string) string {
	// We can grab any non-flag CLI args like this. "Non-flag" meaning anything
	// passed to the program once all flag values have been set -- e.g. in:
	//    $ go run main.go -flag1 value1 -flag2 value2 cliArg
	// "cliArg" is a non-flag arg, since value1 and value2 are passed to their
	// respective flags
	var msg string
	if len(args) > 0 {
		msg = fmt.Sprintf("You provided %d command-line arg(s), and they were: %q", len(args), args)
	} else {
		msg = "You provided no command-line args besides flags"
	}
	return msg
}

// This func actually does something besides printing stupid tutorial garbage :)
func hitSomePages(urls []string) string {
	statusMap := make(map[string]string)
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		statusMap[url] = res.Status
	}

	msg := fmt.Sprintf("You hit %d URL(s), and they responded as follows:\n", len(urls))
	for k, v := range statusMap {
		msg += fmt.Sprintf("	%s -- %s\n", k, v)
	}

	return msg
}

func main() {
	// The first thing you should do in your main function is parse the received
	// flag values -- otherwise the values won't be accessible!
	flag.Parse()

	// All these print calls look kind of yucky, but it's to show the easiest
	// way to test the functions themselves (i.e. if they return a value)
	fmt.Println(getHello(*name))
	fmt.Println(getWeatherReport(*weatherReport))

	args := flag.CommandLine.Args()
	fmt.Println(getCLIArgsMessage(args))

	fmt.Print(hitSomePages(args))
}
