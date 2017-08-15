package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	
	/*data := struct {
		Name string
	}{"John Smith"}*/

	data := struct {
		Name string
	}{"<script>alert('Howdy!');</script>"}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}


/*When you run your program, the html/template package will look at
every variable you want to print out and will adjust the value of each of these
based on where the variable is being used in your template. In this example
we are using the variable inside of an HTML section, so the html/template
package replaces the < character with &lt;. This is done to ensure that when
your template is viewed in a web browser, you will see the < character on the
screen.*/