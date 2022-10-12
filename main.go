package main

import (
	"flag"
	"fmt"

	cyoa "github.com/mouad-eh/cyoa/cyoa"

	// "html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port the CYOA application will be running on.")
	filename := flag.String("file", "gopher.json", "the JSON file containing the CYOA story.")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}
	// template to test options
	// t := template.Must(template.New("").Parse("hello world!"))
	// we can create a new pathFunc to test pathFunc option, we have to change the template so it can work with it
	// /story/ is the new /
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting serer on port: %d.\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
