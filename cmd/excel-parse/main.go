package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("hello world")
	inFile, err := excelize.OpenFile("/Users/shane/Downloads/Qeubee功能列表.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := inFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetName := "手机功能列表"
	columns := []byte("ABCDE")
	var tasks []Task
	for i := 2; i <= 100; i++ {
		var row []string = nil
		for _, column := range columns {
			cell := fmt.Sprintf("%c%d", column, i)
			value, err := inFile.GetCellValue(sheetName, cell)
			if err != nil {
				fmt.Println(err)
				return
			}
			row = append(row, value)
		}
		task := Task{
			Title:          fmt.Sprintf("%s/%s/%s", row[1], row[2], row[3]),
			TaskFlowStatus: "未完成",
			Note:           row[4],
		}
		tasks = append(tasks, task)
	}

	outFile, err := excelize.OpenFile("/Users/shane/Downloads/qeubee手机端新项目群_任务_template.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, task := range tasks {
		row := i + 3
		if err := outFile.SetCellValue("template", fmt.Sprintf("A%d", row), task.Title); err != nil {
			fmt.Println(err)
			return
		}
		if err := outFile.SetCellValue("template", fmt.Sprintf("C%d", row), task.TaskFlowStatus); err != nil {
			fmt.Println(err)
			return
		}
		if err := outFile.SetCellValue("template", fmt.Sprintf("J%d", row), task.Note); err != nil {
			fmt.Println(err)
			return
		}
		if err := outFile.SetCellValue("template", fmt.Sprintf("L%d", row), "功能点"); err != nil {
			fmt.Println(err)
			return
		}

	}
	if err := outFile.Save(); err != nil {
		fmt.Println(err)
		return
	}
}

type Task struct {
	Title          string
	TaskFlowStatus string
	Note           string
}
