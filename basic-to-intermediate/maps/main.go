package main

import (
	"fmt"
)

func main() {
	ncrPopulation := map[string]int{
		"Quezon City": 2936116,
		"Manila":      1780148,
		"Caloocan":    1583978,
		"Taguig":      804915,
		"Pasig":       755300,
		"Paranaque":   665822,
		"Valenzuela":  620422,
		"Las Pinas":   588894,
		"Makati":      582602,
		// "Muntinlupa":  504509,
	}
	ncrPopulation["Muntinlupa"] = 504509
	fmt.Println("ncrPopulation")
	fmt.Println(ncrPopulation)
	fmt.Printf("Length of ncrPopulation: %v\n", len(ncrPopulation))
	ncr := ncrPopulation // Passing using := will pass the address(be carefull)
	ncr["test"] = 12
	fmt.Printf("Length of ncrPopulation: %v\n", len(ncrPopulation))
	fmt.Println(ncrPopulation)
	fmt.Printf("The population of Manila is %v\n", ncrPopulation["Manila"])
	fmt.Printf("Before delete in map %v\n", ncrPopulation)
	delete(ncr, "test")
	fmt.Printf("After delete in map %v\n", ncrPopulation)

	population, isCount := ncrPopulation["Manila"]
	fmt.Println(population, isCount)
	fmt.Println("isCount.........", isCount)
	_, haveCount := ncrPopulation["Makati"]
	fmt.Println(haveCount)

	// Other way to create map
	mmm := map[[3]int]string{}
	// mmm := map[int]string{}
	mmm[[3]int{1, 2, 3}] = "test"
	mmm[[3]int{1, 2, 3}] = "test"
	// delete(mmm, [3]int{1, 2, 3})
	fmt.Println("mmm.................", mmm)
	fmt.Println("mmm length.................", len(mmm))

	ncrPopulation2 := make(map[string]int)
	ncrPopulation2 = map[string]int{
		"Quezon City": 2936116,
		"Manila":      1780148,
		"Caloocan":    1583978,
		"Taguig":      804915,
		"Pasig":       755300,
		"Paranaque":   665822,
		"Valenzuela":  620422,
		"Las Pinas":   588894,
		"Makati":      582602,
		"Muntinlupa":  504509,
	}
	fmt.Println(ncrPopulation2)
	// Find what second argument
	ncrPopulation3 := make(map[string]int, 2)
	ncrPopulation3["first"] = 1
	ncrPopulation3["second"] = 2
	ncrPopulation3["third"] = 3
}
