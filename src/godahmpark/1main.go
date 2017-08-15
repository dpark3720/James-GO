package main

import  (
	"fmt"
	"net/http"
	

func handlerFunc (w http.ResponseWriter, r *http.Request)	{
	w.Header() .Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Look Christina it worked!</h1>")
	} else if r.URL.Path == "/contact"	{
		fmt.Fprint (w, "toget in touch, please send an email "+
			"to <a href=\"mailto:support@usegolang.com\">"+"
			"support@usegolang.com</a>.")
	}
}
fnc main () {
	
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe (":3000", nil)
}
