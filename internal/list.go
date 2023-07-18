package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ParamAccess struct {
    Url string 
    Pass string
    User string
}

func GetQueues(queue string, env string) map[string]float64 {

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
    result := make(map[string]float64) 
    for _, el := range body {
        queueName := el["name"].(string)
        if showQueue != false {
            if(strings.Contains(queueName, queue)){
                result[queueName] = el["messages_ready"].(float64)
            }
        } else {
            result[queueName] = el["messages_ready"].(float64)
        }
    }
    return result 
}

