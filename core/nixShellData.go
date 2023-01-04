package core

type NixShellData struct {
	ProjectName    string
	ExecName       string
	Pkgs           []string
	AdditionalPkgs []string
}

var (
	GoCliShell NixShellData = NixShellData{
		Pkgs: []string{"figlet", "go", "cobra-cli"},
	}
)
