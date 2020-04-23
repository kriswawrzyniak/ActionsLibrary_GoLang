package actions

import (
	"encoding/json"
	"sync"
)

/*
** Structs
 */
//Actions List thread safe struct
type SafeActionsList struct {
	mux sync.Mutex
	m   map[string]int
	n   map[string]int
}

//Actions Decode json struct
type ActionDecode struct {
	Action string
	Time   int
}

//Actions Encode json struct
type ActionEncode struct {
	Action string
	Avg    int
}

/*
**Export Functions
 */

//Actions List Constructor
func NewSafeActionsList() *SafeActionsList {
	list := new(SafeActionsList)
	list.m = make(map[string]int)
	list.n = make(map[string]int)
	return list
}

//Adds an action to the actions list by converting the json string; thread-safe
func (a *SafeActionsList) AddAction(jsonString string) error {
	var action ActionDecode
	err := json.Unmarshal([]byte(jsonString), &action)
	if err != nil {
		return err
	}

	a.mux.Lock()
	defer a.mux.Unlock()

	if _, ok := a.m[action.Action]; ok {
		a.m[action.Action] += action.Time
		a.n[action.Action] += 1
	} else {
		a.m[action.Action] = action.Time
		a.n[action.Action] = 1
	}

	return nil
}

//Gets the average times for all the actions added to the actions list by returning a json string; threadsafe
func (a *SafeActionsList) Statistics() (string, error) {

	a.mux.Lock()
	defer a.mux.Unlock()

	encodeSlice := make([]ActionEncode, 0)
	for key, element := range a.m {
		count := a.n[key]
		average := element / count
		encode := ActionEncode{
			Action: key,
			Avg:    average,
		}
		encodeSlice = append(encodeSlice, encode)
	}
	encodeJson, err := json.Marshal(encodeSlice)
	if err != nil {
		return "", err
	}
	return string(encodeJson), nil
}
