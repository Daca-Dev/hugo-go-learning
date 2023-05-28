package coreapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (pc *ProjectContext) GetProyectConfiguration() (StronglifyProject, error) {
	data := StronglifyProject{}
	client := http.Client{}
	url := fmt.Sprintf(
		"%s/projects/%d/build-config/",
		BaseUrl,
		pc.ProjectId,
	)
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return data, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", Token))

	res, err := client.Do(req)
	if err != nil {
		return data, err
	}

	err = json.NewDecoder(res.Body).Decode(&data)

	return data, err
}
