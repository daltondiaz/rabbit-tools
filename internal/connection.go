package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func GetConnection(filter string, env string){
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
    result.Header = table.Row{"Name", "Virtual Host", "User", "State" } 
    hasFilter := false
    if filter != "" {
        hasFilter = true
    }
    for _, el := range body { 
        queue := el["name"].(string)
        vhost := el["vhost"].(string)
        user := el["user"].(string)
        state := el["state"].(string)
        if hasFilter {
            if strings.Contains(queue, filter) {
                row := table.Row{queue, vhost, user, state}
                result.Contents = append(result.Contents, row)
                count ++
            }
        } else {
            row := table.Row{queue, vhost, user, state}
            result.Contents = append(result.Contents, row)
            count ++
        }
    }
    if hasFilter {
        result.Footers = append(result.Footers, table.Row{"Total Items Filtered", count})
    }
    result.Footers = append(result.Footers, table.Row{"Total of Connections", len(body)})
    PrettyResult(result)
}
