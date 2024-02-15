package internal

import (
	"encoding/json"
	"group/internal/models"
	"io"
	"net/http"
)

func GetArtists() ([]models.Artists, error) {
	url_artists := "https://groupietrackers.herokuapp.com/api/artists"

	artistsList := []models.Artists{}

	err := GetJson(url_artists, &artistsList)
	if err != nil {
		return nil, err
	}
	return artistsList, nil
}

func GetRelation() (models.Relation, error) {
	url_relation := "https://groupietrackers.herokuapp.com/api/relation"
	relation := &models.Relation{}
	err := GetJson(url_relation, &relation)
	if err != nil {
		return *relation, err
	}
	return *relation, nil
}

func GetJson(url string, object interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	dataBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(dataBody, object)
	return err
}
