package validator

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func ReturnValidCombos(rCombos []string) []string {
	vCombos := []string{}
	for i := 0; i < len(rCombos); i++ {
		if isNameValid(rCombos[i]) {
			vCombos = append(vCombos, rCombos[i])
		}
	} 
	return vCombos
}

func JsonDecode(raw string) map[string]interface{} {
	byt := []byte(raw)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return dat
}

func isNameValid(name string) bool {
	url := "https://auth.roblox.com/v2/usernames/validate?request.username=" + name + "&request.birthday=2000-05-05T05%3A00%3A00.000Z"
	request := sendGetRequest(url)
	if request == nil {
		fmt.Println("Failed get request.")
		return false
	}
	code := request["code"]
	codeInt := int(code.(float64))
	if codeInt == 0 {
		return true
 	} else {
		return false
	}
}

func sendGetRequest(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//Convert the body to type string
	sb := string(body)
	return JsonDecode(sb)
}