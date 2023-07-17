package internal

import "github.com/jedib0t/go-pretty/v6/table"

type PrettyModelResult struct {
	Title    string
	Header   table.Row
	Contents []table.Row
	Footers  []table.Row
}
