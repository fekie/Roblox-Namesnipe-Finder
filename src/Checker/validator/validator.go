package validator

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

func ReturnValidCombos(rCombos []string) []string {
	vCombos := []string{}
	for i := range rCombos {
		if isNameValid(rCombos[i]) {
			vCombos = append(vCombos, rCombos[i])
		}
	} 
	return vCombos
}


func isNameValid(name string) bool {
	url := "https://auth.roblox.com/v2/usernames/validate?request.username=" + name + "&request.birthday=2000-05-05T05%3A00%3A00.000Z"
	request := sendGetRequest(url)
	if request == nil {
		log.Fatalf("Failed to get request")
	}
	codeInt := int(request["code"].(float64))
	if codeInt == 0 {
		return true
 	} else {
		return false
	}
}

func sendGetRequest(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error occurred while sending get request, %v", err)
		
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occurred while reading body %v", err)
	} 
	byt := []byte(string(body))
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return dat
}
