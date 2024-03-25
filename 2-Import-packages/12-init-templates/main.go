package main

import (
	"os"
	"text/template"
)

type Class struct {
	name string
	hours int
}

func main()  {
	class := Class{"Go", 40}
	tmp := template.New("ClassTemplate")
	tmp, _ = tmp.Parse("Class: {{.nome}} - Hours: {{.hours}}")
	err := tmp.Execute(os.Stdout, class)
	if err != nil {
		panic(err)
	}

}
