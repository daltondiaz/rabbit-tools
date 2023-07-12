package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PurgeQueue(queue string) {
    params := LoadEnvVariables()
    client := &http.Client{}
    queues := GetQueues(queue)
    sum := 0.0
    for queueName, value := range queues {
        url := (params.Url + "/" + queueName + "/contents")
        sum += value
        req, err := http.NewRequest("DELETE", url, nil)
        req.SetBasicAuth(params.User, params.Pass)
        resp, err := client.Do(req)
        if err != nil {
            log.Fatal(err)
        }
        bodyText, err := ioutil.ReadAll(resp.Body)
        fmt.Println("Purge ", value, " from queue ", queueName, " resp: ", bodyText)
        if err != nil {
            log.Fatal(err)
        }
    }
    fmt.Println("Total of purge items: ", sum)
}

