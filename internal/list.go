package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetQueues(queue string, env string, greater int) map[string]int {

	params := LoadEnvVariables(env)
	client := &http.Client{}
	req, err := http.NewRequest("GET", params.Url, nil)
	req.SetBasicAuth(params.User, params.Pass)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	var body []map[string]any
	var showQueue bool
	if len(queue) > 0 && queue != "all" {
		showQueue = true
	} else {
		showQueue = false
	}
	json.Unmarshal(bodyText, &body)
	result := make(map[string]int)
	for _, el := range body {
		queueName := el["name"].(string)
		if showQueue != false {
			if strings.Contains(queueName, queue) {
				floatMsgRd := el["messages_ready"].(float64)
				result[queueName] = int(floatMsgRd)
			}
		} else {
			floatMsgRd := el["messages_ready"].(float64)
			result[queueName] = int(floatMsgRd)
		}
	}
	if greater > 0 {
        filterResult := make(map[string]int)
		for key, value := range result {
			if value >= greater {
                filterResult[key] = value
			}
		}
        return filterResult
	}

	return result
}
