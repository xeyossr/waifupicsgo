package waifupicsgo

import (
	"encoding/json"
)

func (categoryClient CategoryClient[A]) GetOne(category string) (string, error) {
	image := Image{}
	resp, err := get(categoryClient.category + "/" + category)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&image); err != nil {
		return "", err
	}

	return image.URL, nil
}

func (categoryClient CategoryClient[A]) GetMany(category string) ([]Image, error) {
	var respData struct {
		Files []string `json:"files"`
	}

	resp, err := post("many/"+categoryClient.category+"/"+category, map[string][]string{"exclude": {}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	var images []Image
	for _, url := range respData.Files {
		images = append(images, Image{URL: url})
	}

	return images, nil
}
