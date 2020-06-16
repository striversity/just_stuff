package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/sirupsen/logrus"
)

var (
	counter = 0
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/sse/dashboard", dashboardHandler)
	addr := ":8080"
	err := http.ListenAndServe(addr, nil)
	logrus.Fatalf("unable to start server on addr %v: %v", err)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	counter++

	fmt.Fprintf(w, "data: %v\n\n", counter) // for EventSource.onmessage handler

	t, err := template.New("foo").Parse(`{{define "T"}}<div><h3>Counter: <span>{{.}}</span></div>{{end}}`)
	if err != nil {
		logrus.Fatalf("unable to parse simple template: %v", err)
		return
	}

	fmt.Fprint(w, "event: tmpl\n") // for EventSource.addEventListener("tmpl")
	fmt.Fprint(w, "data: ")
	err = t.ExecuteTemplate(w, "T", counter)
	if err != nil {
		logrus.Fatalf("unable to write template output: %v", err)
		return
	}

	fmt.Fprint(w, "\n\n", counter)

	fmt.Println("Writing text to client:")
	t.ExecuteTemplate(os.Stdout, "T", counter)
	fmt.Println()
}
