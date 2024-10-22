package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ensureFiles(importantFiles map[int]string) bool {
	for _, fileName := range importantFiles {
		if _, err := os.Stat(fileName); err == nil {
			continue
		} else if err == os.ErrNotExist {
			return false
		}
		return false
	}
	return true
}

func GetInfo[T any](url string) (*T, error) {
	fmt.Println("sent to url:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Convert response body to Todo struct
	var data T

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func getDetailedInfo[T any](obj T, idNum int) (*T, error) {
	if idNum < 1 || idNum > 52 {
		return nil, fmt.Errorf("id num is not present")
	}
	return nil, nil
}

func filterArtist(a *[]Artist, f func(artist Artist) bool) (*Artist, error) {
	for _, v := range *a {
		if f(v) {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
func filterDate(a *Dates, f func(artist Date) bool) (*Date, error) {
	for _, v := range (*a).Index {
		if f(v) {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func filterLocation(a *Locations, f func(artist Location) bool) (*Location, error) {
	for _, v := range (*a).Index {
		if f(v) {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func filterRelation(a *Relations, f func(artist Relation) bool) (*Relation, error) {
	for _, v := range (*a).Index {
		if f(v) {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func GetConcertDates(locations *Location, relation *Relation) []string {
	placeholder := ""
	toReturn := []string{}

	for _, location := range locations.Locations {
		placeholder += location + ": "
		dates := relation.DatesLocations[location]
		for i, date := range dates {
			placeholder += date
			if i != len(dates)-1 {
				placeholder += " and "
			}
		}
		toReturn = append(toReturn, placeholder)
		placeholder = ""
	}
	fmt.Println(toReturn)
	return toReturn
}
func GetConcertDates2(locations *Location, relation *Relation) []string {
	placeholder := ""
	toReturn := []string{}
	for _, location := range locations.Locations {
		placeholder += location + ": "
		dates := relation.DatesLocations[location]
		for i, date := range dates {
			placeholder += date
			if i != len(dates)-1 {
				placeholder += " and "
			}
		}
		toReturn = append(toReturn, placeholder)
		placeholder = ""
	}
	fmt.Println(toReturn)
	return toReturn
}
