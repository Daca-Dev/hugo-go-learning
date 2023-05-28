package main

import (
	"fmt"

	coreapi "daca-dev.com/generator/core_api"
)

func main() {
	project := coreapi.ProjectContext{
		ProjectId:      13,
		ProjectVersion: "test",
		GenerationId:   1,
	}
	stronglifyProject, err := project.GetProyectConfiguration()
	if err != nil {
		fmt.Println(err)
	}

	coreapi.BuildConfigFile(stronglifyProject)

}
