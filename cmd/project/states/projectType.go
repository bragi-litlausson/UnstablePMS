package states

type ProjectType struct {
	name string
	pkgs string
}

var ProjectTypes = map[string]ProjectType{
	GoCli.name: GoCli,
}

var (
	GoCli ProjectType = ProjectType{
		name: "Go CLI",
		pkgs: `
		`,
	}
)
