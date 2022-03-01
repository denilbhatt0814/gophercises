package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/denil/cyoa"
)

func main() {
	// setting the cmd flags
	port := flag.Int("port", 3000, "The port to start the CYOA web app on")
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// opening the file
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	// Decoding JSON file
	story, err := cyoa.JsonStory(f)
	if err != nil{
		panic(err)
	}

	tmpl := template.Must(template.New("").Parse(storyTmpl))


	// strating the server
	h := cyoa.NewHandler(story, cyoa.WithTemplate(tmpl), cyoa.WithPathFunc(pathFn))
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/"{
		path = "/story/intro"
	} 
	return path[len("/story/"):]	
}

var storyTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>

    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
            <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`