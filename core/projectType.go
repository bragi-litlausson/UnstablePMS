package core

type ProjectType struct {
	Name string
	Pkgs []string
}

var ProjectTypes = map[string]ProjectType{
	GoCli.Name: GoCli,
}

var (
	GoCli ProjectType = ProjectType{
		Name: "Go CLI",
		Pkgs: []string{"go", "cobra-cli"},
	}
)
