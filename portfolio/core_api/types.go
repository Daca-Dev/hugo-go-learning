package coreapi

type StronglifyProject struct {
	ID              int64           `json:"id"`
	Name            string          `json:"name"`
	Domain          Domain          `json:"domain"`
	Description     string          `json:"description"`
	Language        string          `json:"language"`
	AdvanceSettings AdvanceSettings `json:"advance_settings"`
	Params          Params          `json:"params"`
	Favicon         string          `json:"favicon"`
	Theme           Theme           `json:"theme"`
	FolderName      string          `json:"folder_name"`
	Version         string          `json:"version"`
	Templates       Templates       `json:"templates"`
}

type AdvanceSettings struct {
}

type Domain struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Params struct{}

type Templates struct {
	Collection []string `json:"collection"`
}

type Theme struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`
}

type ProjectContext struct {
	ProjectId      int    `json:"project_id"`
	ProjectVersion string `json:"project_version"`
	GenerationId   int    `json:"generation_id"`
}
