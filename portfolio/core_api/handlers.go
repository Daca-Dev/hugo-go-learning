package coreapi

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type ProjectTomlConfig struct {
	BaseURl      string   `toml:"baseURL"`
	LanguageCode string   `toml:"languageCode"`
	Title        string   `toml:"title"`
	Theme        string   `toml:"theme"`
	Params       struct{} `toml:"params"`
	UglyURLs     bool     `toml:"uglyurls"`
}

func BuildConfigFile(project StronglifyProject) error {
	data := ProjectTomlConfig{
		BaseURl:      project.Domain.Name,
		LanguageCode: "en",
		Title:        project.Name,
		Theme:        "basic",
		Params:       project.Params,
		UglyURLs:     true,
	}
	path := fmt.Sprintf("%s/%s", os.TempDir(), "config.toml")
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	err = toml.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}

	return nil

}
