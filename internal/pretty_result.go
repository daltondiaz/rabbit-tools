package internal

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrettyResult(result PrettyModelResult){ 
    tableResult := table.NewWriter()
    tableResult.SetOutputMirror(os.Stdout)
    tableResult.AppendHeader(result.Header)
    tableResult.AppendRows(result.Contents)
    tableResult.AppendSeparator()
    for _, row := range result.Footers {
        tableResult.AppendFooter(row)
    }
    tableResult.Render()
}
