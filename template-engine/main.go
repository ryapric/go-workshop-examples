package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template" // this could also be "text/template", but the html package has some security features we get for free, and this program is general-purpose
	"os"
)

// render takes in the raw template string, and a map with string keys and *any*
// kind of values (this is a pretty common way to generalize template rendering,
// as well as e.g. reading unknown JSON schemas, etc.)
func render(tplText string, data map[string]any) string {
	tpl, err := template.New("tpl").Parse(tplText)
	if err != nil {
		fmt.Println("error parsing text template: ", err.Error())
		os.Exit(1)
	}

	// If you want to return this as a string vs. rendering straight to a file
	// (which is what template.Execute expects), you just need something that
	// implements the io.Writer interface (which bytes.Buffer does) and then
	// call Execute using the object's pointer:
	var rendered bytes.Buffer
	err = tpl.Execute(&rendered, data)
	if err != nil {
		fmt.Println("error executing parsed template with input data: ", err.Error())
		os.Exit(1)
	}

	return rendered.String()
}

// Currently, we're rendering an HTML web page here, but you can do anything!
func main() {
	// We need to read in two (or more) files: the template(s) themselves, and
	// the data used to render them
	tplBytes, err := os.ReadFile("./index.html.tpl")
	if err != nil {
		fmt.Println("error reading in template text file: ", err.Error())
		os.Exit(1)
	}

	configBytes, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("error reading in config.json: ", err.Error())
		os.Exit(1)
	}

	// Here's another case where the unmarshal target's schema isn't something
	// we care about, so we can just use a map[string]any
	configData := make(map[string]any)
	err = json.Unmarshal(configBytes, &configData)
	if err != nil {
		fmt.Println("error unmarshaling connfig.json: ", err.Error())
		os.Exit(1)
	}

	// Now that we have a Go type with interpolable data, we can provide it to
	// render() to output a rendered string!
	rendered := render(string(tplBytes), configData)
	fmt.Print(rendered)
}
