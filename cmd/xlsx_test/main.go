package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	xlsxFile, err := xlsx.OpenFile("./bk_cmdb_export_inst_physical_server.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range xlsxFile.Sheets[0].Rows[0:10] {
		for _, cell := range row.Cells {
			text := cell.String()
			fmt.Printf("%s\n", text)
		}
	}

}
