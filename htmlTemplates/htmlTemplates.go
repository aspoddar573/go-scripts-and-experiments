package htmlTemplates

import (
	"html/template"
	"log"
	"os"
)

func TemplateToImage() {
	//const text = "<head>{{.Greeting}} {{.Name}}! Good day!</head>"
	const text = "<head>{{.Greeting}} {{.Name}}! Good {{.Time}}!</head>"

	data := struct {
		Greeting string
		Name     string
		Time     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
		Time:     "day",
	}

	t := template.Must(template.New("tpl").Parse(text))
	log.Println(t.Parse(text))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildingATemplate() {

	//const text = "<head>{{.Greeting}} {{.Name}}! Good day!</head>"
	const text = "<head>{{.Greeting}} $.Name$! Good {{.Time}}!</head>"
	data := struct {
		Greeting string
		Name     string
		Time     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
		Time:     "day",
	}

	t := template.Must(template.New("tpl").Parse(text))
	log.Println(t.Parse(text))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

}
