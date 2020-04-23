package main

import (
	"actions/actions"
	"fmt"
)

func main() {
	actionsObj := actions.NewSafeActionsList()
	testString1 := `{"action":"jump", "time":100}`
	testString2 := `{"action":"run", "time":75}`
	testString3 := `{"action":"jump", "time":200}`
	// testString4 := ``
	err := actionsObj.AddAction(testString1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = actionsObj.AddAction(testString2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = actionsObj.AddAction(testString3)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// err = actionsObj.AddAction(testString4)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	var json string
	json, err = actionsObj.Statistics()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("RETURNED", json)
}
