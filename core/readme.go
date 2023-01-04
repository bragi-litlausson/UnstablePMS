package core

import (
	"os"
	"text/template"
)

const (
	TemplateReadme = "Yes, from template"
	EmptyReadme    = "Yes, empty"
	NoReadme       = "No"
)

func CreateReadmeFile(t string, name string) {
	switch {
	case t == TemplateReadme:
		createReadmeTemplate(name)
	case t == EmptyReadme:
		WriteTextFile("README.md", "")
	case t == NoReadme:
		return
	}
}

func createReadmeTemplate(name string) {
	t := template.Must(template.New("readme").Parse(readMeTemplate))
	file, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, name)
	if err != nil {
		panic(err)
	}
}

const readMeTemplate = `# {{.}}

Short description of the project

# Description

More in-depth project description

## Built with

### Library #1

Awesome library used to create this project

# Getting started

## Dependencies

Any requirements to run this

## Installing

How to install damn thing

## Usage

Grab big hammer and...

# Known issues

FAQ it

# Contributors

Awesome people that helped

# Version history

For more detailed information check release history(Add link)

- v1.0.0
  - awesome feature
  - not so awesome feature

# License

This project is licensed under [NAME HERE] License - see LICENSE file for more details.

# Standing on shoulders of the giants

Thank you mom!`
