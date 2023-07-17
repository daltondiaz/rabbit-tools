package internal

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/jedib0t/go-pretty/v6/table"
)

// Do purge of the queues found from queue argument
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
        if err != nil {
            log.Fatal(err)
        }
        _ = bodyText
    }
    var prettyResult PrettyModelResult
    prettyResult.Title = "Purged Queues"
    prettyResult.Header = table.Row{"Purge Queue", "Messages"}
    totalItems := 0.0
    for queueName, value := range queues {
        row := table.Row{queueName, value}
        totalItems += value
        prettyResult.Contents = append(prettyResult.Contents, row)
    }
    prettyResult.Footers = append(prettyResult.Footers, table.Row{"Total Purged Messages", sum})
    PrettyResult(prettyResult)
}

