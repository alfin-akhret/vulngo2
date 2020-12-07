package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/kennygrant/sanitize"
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
	str := validate(Name)
	fmt.Println(str)
	fmt.Fprintf(w, str)
}

func validate(str string) string {
	sanitizedStr := str
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(str) {
		sanitizedStr = sanitize.Accents(str)
		return sanitizedStr
	}
	return sanitizedStr
}
