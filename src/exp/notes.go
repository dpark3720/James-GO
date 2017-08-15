package main

import (
	"html/template"
	/* Package template (html/template) implements data-driven templates 
	for generating HTML output safe against code injection. It provides 
	the same interface as package text/template and should be used instead 
	of text/template whenever the output is HTML.*/
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
		/* ParseFiles creates a new Template and parses 
		the template definitions from the named files*/
	if err != nil {
		panic(err)
		/* We then check to see if an error was returned using if err != nil {
		... }. If an error is present, the err variable will not be nil, in which case
		we panic. Otherwise we continue on with our program and should have a valid
		pointer to use assigned to the t variable.*/
	}
	
	data := struct {
		Name string
	}{"James"}
	/*Next up we create the data variable, which is just a struct with a Name
	attribute. When we instantiate data we set the Name attribute to “James”.*/

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		/*Finally we use the template created from hello.gohtml and we call t.Execute(os.Stdout,
		data). With this line we are telling our program that we want to execute the
		template and print the output to os.Stdout, which is just your terminal, and
		we are also stating that we want to pass the data variable into our template so
		that it can use its attributes when rendering the template.*/
	}
}