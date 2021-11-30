package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type file struct {
	filename string
}

func (f *file) Read() string {
	content, err := ioutil.ReadFile(f.filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func (f *file) Append(text string) {
	fl, err := os.OpenFile(f.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	text = text + " "
	if _, err := fl.WriteString(text); err != nil {
		log.Fatal(err)
	}
}

func main() {
	myf := &file{"myfile.txt"}
	myf.Append("This is a test to append.")
	fmt.Println(myf.Read())
}