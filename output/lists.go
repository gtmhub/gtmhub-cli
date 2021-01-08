package output

import (
	"gtmhub-cli/model"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func PrintKrsFromList(columns []model.ListColumn, data []map[string]interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"id"}
	for _, column := range columns {
		header = append(header, column.FieldName)
	}
	t.AppendHeader(header)
	for _, metric := range data {
		dataRow := table.Row{metric["id"]}
		for _, column := range columns {
			dataFromRow, found := metric[column.FieldName]
			if found {
				dataRow = append(dataRow, dataFromRow)
			} else {
				dataRow = append(dataRow, "not specified")
			}
		}
		t.AppendRow(dataRow)
	}
	t.Render()
}
