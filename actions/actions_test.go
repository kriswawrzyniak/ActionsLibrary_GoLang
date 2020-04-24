package actions

import "testing"

func TestNewSafeActionsList(t *testing.T) {
	var testObj = NewSafeActionsList()
	if testObj == nil {
		t.Error("Unable to create Safe Actions List:", testObj)
	}
}

func TestAddAction(t *testing.T) {
	var testObj = NewSafeActionsList()
	testString1 := `{"action":"jump", "time":100}`
	err := testObj.AddAction(testString1)
	if err != nil {
		t.Error("Error in AddAction:", err)
	}
	if testObj.total["jump"] != 100 {
		t.Error("Error in AddAction:", "Adding jump - 100 got", testObj.total["jump"])
	}

}
