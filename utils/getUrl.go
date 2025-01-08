package utils

import (
	"errors"
	"url-shortner/model"
)

func GetUrl(id string) (model.URL, error) {
	url, ok := model.UrlDB[id]
	if !ok {
		return model.URL{}, errors.New("URL Not Found")
	}
	return url, nil
}
