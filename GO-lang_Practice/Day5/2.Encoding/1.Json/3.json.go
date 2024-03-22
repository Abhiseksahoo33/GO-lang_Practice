package main

import (
	"encoding/json"
	"fmt"
)

var jsonContent = `{"permission": {"resourceURL": "https://ndc1-testbed123.ftcontentserver.rcs.mnc001.mcc262.pub.3gppnetwork.org/upta3/download/8385a4c2-19d9-49f5-b3a8-a12631e19060?t=audio%2Famr&amp;s=11334&amp;e=20200119T154306Z","rcptCount":"2","granteeList": [{"grantee": "19663136600","archivalEnabled": "true"},{"grantee": "19663136601","archivalEnabled": "true"}]}}`

func main() {
	fmt.Println("unknown json content")

	mp := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonContent), &mp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mp)
	}

	fmt.Println("Contents of Map")
	for _, value := range mp {
		//fmt.Println(key, ":", value)

		switch value.(type) {
		case map[string]interface{}:
			mp1 := value.(map[string]interface{})
			for k, v := range mp1 {
				fmt.Println(k, ":", v)
			}

		}
	}
}

//Assgn1: Array of json received marshal it
