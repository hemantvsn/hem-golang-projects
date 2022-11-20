package main

import "fmt"

type Country struct {
	code       string
	name       string
	population int
	gdp        float64
}

func main() {
	emptyMap()
	simpleMap()
	objectMap()

}

func makeMyMap() {
	var countryMap = make(map[string]Country)

	fmt.Println(countryMap)
}

// MAP using MAKE
func emptyMap() {
	var countryMap = map[string]Country{}
	fmt.Println(countryMap)

	countryMap["IND"] = Country{name: "INDIA"}
}

func objectMap() {
	var countryMap = map[string]Country{
		"IND": Country{code: "IND", name: "India", population: 100_000_000, gdp: 10},
		"PAK": Country{code: "PAK", name: "Pakistan", population: 50_000_000, gdp: -10},
	}
	// add some variable
	countryMap["HEM"] = Country{}
	fmt.Println(countryMap)

	delete(countryMap, "HEM")
	delete(countryMap, "xxx")
	fmt.Println(countryMap)
	printMap(countryMap)

}

func simpleMap() {
	var countryPhoneNumberMap = map[string]int{
		"INDIA":    91,
		"PAKISTAN": 92,
		"UK":       44,
	}

	fmt.Println(countryPhoneNumberMap)
}
