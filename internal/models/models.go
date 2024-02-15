package models

import (
	"strings"
)

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstalbum"`
	Related      map[string][]string
}

type Relation struct {
	Index []struct {
		Id            int                 `json:"id"`
		DatesLocation map[string][]string `json:"datesLocations"`
	}
}

type ArtistInfo struct {
	Artist  Artists
	Related map[string][]string
}

func (relMap *Relation) GetNormalMap(id int) map[string][]string {
	m := map[string][]string{}
	for loc, dates := range relMap.Index[id].DatesLocation {
		location := strings.ReplaceAll(loc, "-", ", ")
		location = strings.ReplaceAll(location, "_", " ")
		location = strings.Title(location)

		m[location] = dates
		location = ""
	}
	return m
}
