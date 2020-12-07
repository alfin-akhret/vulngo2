package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	var portNum = flag.String("p", "80", "Specify application server listening port")
	flag.Parse()
	fmt.Println("Vulnapp server listening : " + *portNum)

	http.HandleFunc("/", sayYourName)

	err := http.ListenAndServe(":"+*portNum, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayYourName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.Form[name]", r.Form["name"])
	var Name string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		Name = strings.Join(v, ",")
	}
	fmt.Println(Name)
	err := sanitize(Name)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
	}
	fmt.Println(Name)
	fmt.Fprintf(w, Name)
}

func sanitize(str string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(str) {
		e := errors.New("Invalid input")
		return e
	}
	return nil
}
