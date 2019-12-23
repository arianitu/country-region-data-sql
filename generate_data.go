package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Region struct {
	Name      string `json:"name"`
	Shortcode string `json:"shortcode"`
}

type Country struct {
	Name      string   `json:"countryName"`
	Shortcode string   `json:"countryShortCode"`
	Regions   []Region `json:"regions"`
}

func main() {
	var countries []Country
	b, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(b, &countries)
	if err != nil {
		log.Fatalln(err)
	}

	countryInsert := "INSERT INTO `countries`(`id`, `name`, `code`) VALUES "
	fmt.Println(countryInsert)

	countryValues := "  (%d, '%s', '%s')"
	cv := func(id int, country Country, isLast bool) string {
		ending := ";"
		if !isLast {
			ending = ","
		}

		return fmt.Sprintf(countryValues,
			id,
			strings.ReplaceAll(country.Name, "'", "\\'"),
			country.Shortcode) + ending
	}

	for idx, country := range countries {
		fmt.Println(cv(idx+1, country, idx == len(countries)-1))
	}
	fmt.Printf("\n")

	regionInsert := "INSERT INTO `regions`(`name`, `code`, `country_id`) VALUES "
	fmt.Println(regionInsert)

	regionValues := "  ('%s', '%s', %d)"
	rv := func(countryId int, region Region, isLast bool) string {
		ending := ";"
		if !isLast {
			ending = ","
		}

		return fmt.Sprintf(regionValues,
			strings.ReplaceAll(region.Name, "'", "\\'"),
			region.Shortcode,
			countryId) + ending
	}

	for countryIdx, country := range countries {
		for regionIdx, region := range country.Regions {
			fmt.Println(rv(countryIdx+1,
				region,
				regionIdx == len(country.Regions)-1 && countryIdx == len(countries)-1))
		}
	}
	fmt.Printf("\n")
}
