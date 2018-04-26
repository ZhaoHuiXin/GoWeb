package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"accessToken":"507b5e08ee444dck887b66bd08672905",
"clientToken":"64e3a5415bfe405d9485f1jf2ea5c68e",
"selectedProfile":{"id":"selID","name":"Bluek404"},
"availableProfiles":[{"id":"测试ava","name":"Bluek404"}]}`
	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
		fmt.Println(dat["selectedProfile"])
		mapTmp := dat["selectedProfile"].(map[string]interface {})
		fmt.Println(mapTmp)
		fmt.Println(mapTmp["id"])

		///*
		//var dat2 map[string]interface{}
		//if err := json.Unmarshal([]byte(jsonStr), &dat2); err == nil {
		//	fmt.Println( dat2["firstName"])
		//}
		//*/
		//
		//mapTmp2 := (dat["availableProfiles"].([]interface {}))[0].(map[string]interface {})
		////mapTmp3 := mapTmp2[0].(map[string]interface {})
		//fmt.Println(mapTmp2["id"])
	}
}

