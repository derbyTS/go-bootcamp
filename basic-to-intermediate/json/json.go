package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"address"`
}

type Employee struct {
	FirstName  string  `json:"first_name"`
	EmployeeId string  `json:"emp_id"`
	Age        int     `json:"age"`
	Address    Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{
		FirstName: "John Doe",
		Age:       32,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{
		FirstName: "John",
		Age:       32,
		Email:     "fakeemail@mail.com",
		Address: Address{
			City:  "Manila",
			State: "Metro Manila",
		},
	}

	jsonData, err = json.Marshal(person1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	jsonData1 := `{ "full_name" : "John Doe", "emp_id" : "0002", "age" : 30, "address": { "city": "Manila", "state": "Metro Manila" }
	}`

	var employeeJson Employee
	err = json.Unmarshal([]byte(jsonData1), &employeeJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Employee from json: %v\n", employeeJson)
	fmt.Println(employeeJson)
	fmt.Printf("Age: %d\n", employeeJson.Age)
	fmt.Printf("Address: %s\n", employeeJson.Address)

	listOfCityAndStates := []Address{
		{City: "New York", State: "NY"},
		{City: "San Francisco", State: "CA"},
		{City: "Dallas", State: "TX"},
	}

	fmt.Println(listOfCityAndStates)
	jsonList, err := json.Marshal(listOfCityAndStates)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonList))

	// Handling unknow JSON structure
	jsonData2 := `{"name": "John", "emp": "123"}`
	var dataMap map[string]any
	err = json.Unmarshal([]byte(jsonData2), &dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dataMap)
	fmt.Println(dataMap["nam"])
	fmt.Println(dataMap["name"])

	// Handling list of json using Slice of Maps
	listJson := `[{"city":"New York","state":"NY"},{"city":"San Francisco","state":"CA"},{"city":"Dallas","state":"TX"}]`
	var dataMapList []map[string]string
	err = json.Unmarshal([]byte(listJson), &dataMapList)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, cityMap := range dataMapList {
		fmt.Printf("%d: %s is in %s\n", i+1, cityMap["city"], cityMap["state"])
	}
}
