package util

import "encoding/json"

func ArraysToString(dataArray []string) string {
	myValue, _ := json.Marshal(dataArray)
	return string(myValue)
}
