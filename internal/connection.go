package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jedib0t/go-pretty/v6/table"
)

func GetConnection(env string){
    params := LoadEnvVariables(env)
    client := &http.Client{}
    connUrl := params.Api + "/connections"
    req, err := http.NewRequest("GET", connUrl, nil)
    req.SetBasicAuth(params.User, params.Pass)
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    var body []map[string]any
    json.Unmarshal(bodyText, &body)
    count := 0
    var result PrettyModelResult
    result.Title = "Connections"
    result.Header = table.Row{"name", "vhost", "user", "state" } 
    for _, el := range body { 
        queue := el["name"].(string)
        vhost := el["vhost"].(string)
        user := el["user"].(string)
        state := el["state"].(string)
        row := table.Row{queue, vhost, user, state}
        result.Contents = append(result.Contents, row)
        count ++
    }
    result.Footers = append(result.Footers, table.Row{"Total of Connections", len(body)})
    PrettyResult(result)
}
