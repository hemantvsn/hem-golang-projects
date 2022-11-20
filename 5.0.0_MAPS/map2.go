package main

import "fmt"

func printMap(countryMap map[string]Country) {
	for key, value := range countryMap {
		fmt.Println("KEY = ", key, " AND VALUE =", value)
	}
}
