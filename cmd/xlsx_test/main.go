package main

import (
	"fmt"
	"log"
	"time"

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

	t, _ := time.ParseInLocation("2006-01-02", "2020-01-01", time.Local)

	fmt.Println(t)
}
