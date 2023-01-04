package states

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bragi-litlausson/UnstablePM/core"
)

func RunNixShellState(projectData *core.ProjectData) {
	prompt := &survey.Confirm{
		Message: "Create shell.nix?",
	}
	create := false
	survey.AskOne(prompt, &create)

	if create == false {
		return
	}
	data := getNixShellData(projectData)
	projectData.ShellData = data

	createShell(projectData)
}

func getNixShellData(projectData *core.ProjectData) core.NixShellData {
	execName := getExecName(projectData.ProjectName)
	pkgs := projectData.ProjectType.Pkgs
	addPkgs := getAdditionalPkgs()

	return core.NixShellData{
		ProjectName:    projectData.ProjectName,
		ExecName:       execName,
		Pkgs:           pkgs,
		AdditionalPkgs: addPkgs,
	}
}

func getExecName(projectName string) string {
	useMsg := fmt.Sprintf("Do you want to set up custom name for executable? (current %s)", projectName)
	usePrompt := &survey.Confirm{
		Message: useMsg,
	}

	use := true
	survey.AskOne(usePrompt, &use)
	if use == false {
		return projectName
	}

	execName := ""
	execPrompt := &survey.Input{
		Message: "Exec name",
	}
	survey.AskOne(execPrompt, &execName)
	return execName
}

func getAdditionalPkgs() []string {
	usePrompt := &survey.Confirm{
		Message: "Do you want to add additional pkgs",
	}
	use := false
	survey.AskOne(usePrompt, &use)

	if use == false {
		return nil
	}

	var pkgs string
	pkgsPrompt := &survey.Input{
		Message: "Additional nix packages (Separate with space)",
	}
	survey.AskOne(pkgsPrompt, &pkgs)

	return strings.Split(pkgs, " ")
}

func createShell(projectData *core.ProjectData) {
	t := template.Must(template.New("readme").Parse(shellTemplate))
	file, err := os.Create("shell.nix")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = t.Execute(file, projectData.ShellData)
	if err != nil {
		panic(err)
	}
}

const shellTemplate = `let
  pkgs = import <nixpkgs> {};
  projectName = "{{.ProjectName}}";
  goPackagePath = "github.com/bragi-litlausson/{{.ProjectName}}";
  execName = "{{.ExecName}}";
  pwd = "$PWD";
  build = pkgs.writeScriptBin "build" ''
    if [ ! -d "./bin" ]
    then
      mkdir ./bin
    fi

    go build -o ./bin/${execName} main.go
  '';
  run = pkgs.writeScriptBin "run" ''
    build
    exec ./bin/${execName} $*
  '';
  format = pkgs.writeScriptBin "format" ''
    gofmt -l -w -s .
  '';
in pkgs.mkShell rec {
  name = projectName;
  buildInputs = with pkgs; [
      format
      build
      run

			{{range .Pkgs}}
			{{.}}
			{{end}}

			{{range .AdditionalPkgs}}
			{{.}}
			{{end}}
  ];
  shellHook= ''
    figlet -f doom "Entering goland"

    export GOPATH="${pwd}/.go"
    export GOCACHE=""
    export GO111MODULE='on'
    export PATH="$PATH:GOPATH";

    if [ ! -f "./go.mod" ]
    then
      go mod init ${goPackagePath}
    fi
    
    figlet -f doom "Happy hacking!"
  '';
}
`
