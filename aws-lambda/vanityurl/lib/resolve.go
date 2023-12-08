package lib

import (
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

// VTP Search Types

type SearchCriteria struct {
	Criteria struct {
		Key string `json:"key"`
	} `json:"criteria"`
	SelectedItems []SearchCriteriaItem `json:"selectedItems"`
	PossibleItems []SearchCriteriaItem `json:"possibleItems"`
}

type SearchCriteriaItem struct {
	Key  string `json:"key"`
	Name string
}

type SearchCars struct {
	Results struct {
		Result struct {
			Cars []struct {
				Key string `json:"key"`
			} `json:"cars"`
		} `json:"result"`
	} `json:"results"`
	Criteria struct {
		Search struct {
			Criterias []SearchCriteria `json:"criterias"`
		} `json:"search"`
	} `json:"criteria"`
}

// VTP Configuration Types

type ConfigCriteriaItem struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type ConfigCriteria struct {
	Key   string               `json:"key"`
	Name  string               `json:"name"`
	Items []ConfigCriteriaItem `json:"items"`
}

type CriteriasCustomType struct {
	Criteria ConfigCriteria `json:"criteria"`
}

type ConfigurationCarCriteria struct {
	Criterias []CriteriasCustomType `json:"criterias"`
}

// Match Types

type Match struct {
	filterKey    string
	filterOption SearchCriteriaItem
	word         string
}

var seoKeys = []string{"t_manuf", "t_model", "t_smod"}

func Resolve(url string) map[string]Match {

	// fetch all cars and searchCriteria
	carsBody := searchCars()
	search := SearchCars{}
	/*errS := */ json.Unmarshal([]byte(carsBody), &search)
	//fmt.Println(search.Results.Result.Cars, errS)
	//fmt.Println(search.Criteria.Search.Criterias, errS)

	// fetch all criteria details
	criteriasBody := configurationCarCriteria()
	criterias := ConfigurationCarCriteria{}
	/*errC := */ json.Unmarshal([]byte(criteriasBody), &criterias)
	//fmt.Println(criterias.Criterias, errC)

	// bundle up criteria keys and details
	for index, criteria := range search.Criteria.Search.Criterias {

		criteria.PossibleItems = lo.Map[SearchCriteriaItem, SearchCriteriaItem](criteria.PossibleItems, func(searchCriteriaItem SearchCriteriaItem, _ int) SearchCriteriaItem {

			filterMatch, _ := lo.Find[CriteriasCustomType](criterias.Criterias, func(configCriteria CriteriasCustomType) bool {
				return configCriteria.Criteria.Key == criteria.Criteria.Key
			})

			filterOptionsMatch, _ := lo.Find[ConfigCriteriaItem](filterMatch.Criteria.Items, func(configCriteriaItem ConfigCriteriaItem) bool {
				return configCriteriaItem.Key == searchCriteriaItem.Key
			})

			return SearchCriteriaItem{Key: searchCriteriaItem.Key, Name: filterOptionsMatch.Name}
		})

		search.Criteria.Search.Criterias[index] = criteria

	}

	// url word matching -->
	//urlParts := strings.Split(url, "/")
	urlParts := strings.FieldsFunc(url, split)

	// try to simple first-match url parts for criteria
	interpretations := make(map[string]Match)

	for _, word := range urlParts {

		var wordMatch = Match{word: word, filterKey: "no-match"}

		for _, seoKey := range seoKeys {
			filterMatch, _ := lo.Find[SearchCriteria](search.Criteria.Search.Criterias, func(searchCriteria SearchCriteria) bool {
				return searchCriteria.Criteria.Key == seoKey
			})

			filterOptionMatch, oM := lo.Find[SearchCriteriaItem](filterMatch.PossibleItems, func(searchCriteriaItem SearchCriteriaItem) bool {
				return normalize(searchCriteriaItem.Name) == word
			})

			if oM {
				wordMatch.filterKey = filterMatch.Criteria.Key
				wordMatch.filterOption = filterOptionMatch
			}

		}

		interpretations[word] = wordMatch

	}

	fmt.Println("url ", url, " has been interpreted to ", interpretations)

	return interpretations
}
