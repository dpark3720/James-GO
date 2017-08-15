package views

import (
	"net/http"
	"html/template"
	"path/filepath"
)
var ViewsDir string = "views"
var ViewsExt string = ".gohtml"
var LayoutDir string = ViewsDir + "/layouts"

func NewView(layout string, files ...string) *View {
	addViewsDirPrefix(files)
	addViewsExtSuffix(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout: layout,
	}
}
type View struct {
	Template *template.Template
	Layout string
}
func addViewsDirPrefix (files []string) {
	for i, f := range files {
		files[i] = ViewsDir + "/" + f
	}
}

func addViewsExtSuffix(files []string) {
	for i, f := range files {
		files[i] = f + ViewsExt
	}
}

func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	w.Header() .Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate (w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + ViewsExt)
	if err != nil {
		panic(err)
	}
	return files
}




/*NewView() -- variadic parameter*/
/*Variadic Parameter --> is essentially a function parameter that can be 0, 1, or any
other number of items as long as they match the correct type, and are the last
argument to the function call. This is represented by using the triple dot (...)
operator before the argument type when declaring the function NewView(). We
are saying that our function can take in any number of strings, and then Go merges
all of these strings into a string slice ([]string) named files for us to use inside
of the function.*/