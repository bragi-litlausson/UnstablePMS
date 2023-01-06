package states

import (
	"github.com/bragi-litlausson/UnstablePMS/core"
	"gopkg.in/yaml.v2"
)

func RunProjectFilesState(data *core.ProjectData) {
	d, err := yaml.Marshal(data)
	if err != nil {
		panic(err)
	}

	s := string(d)
	core.WriteTextFile(".upm", s)
}
