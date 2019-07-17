package main

import (
	"bytes"
	"encoding/json"
	"excel_to_json/parseConfig"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	println("**** START ****")
	var config = parseConfig.New("yf.json")
	var plat = config.Get("data").(map[string]interface{})
	inPath := plat["inPath"].(string)
	serverOutPath := plat["serverOutPath"].(string)

	processYf(inPath, serverOutPath)

	fmt.Println("\n **** DONE ****")
	fmt.Print("\n Press 'Enter' to continue...\n")
	fmt.Scanln()
}

func processYf(inpath, serverPath string) {
	files, _ := ioutil.ReadDir(inpath)
	var buf bytes.Buffer
	buf.WriteString("package sdata\n")
	for _, file := range files {
		excelYf(inpath, file.Name(), serverPath)

	}
}

func excelYf(path string, fileName string, serverPath string) {
	if strings.HasPrefix(fileName, "~") {
		return
	}
	if !strings.HasSuffix(fileName, ".xlsx") {
		return
	}
	println("process " + path + fileName)
	xlFile, err := xlsx.OpenFile(path + fileName)
	if err != nil {
		fmt.Println("open file error")
	}
	for _, item := range xlFile.Sheets {
		processOnce(serverPath, item)
	}
}

func processOnce(serverPath string, sheet *xlsx.Sheet) {
	sheetName := sheet.Name

	if !strings.Contains(sheetName, "|") {
		return
	}

	jsonName := strings.Split(sheetName, "|")[1] + ".json"

	celLen := len(sheet.Rows[0].Cells)
	rowLen := 0

	for _, row := range sheet.Rows {

		if len(row.Cells) > 0 && row.Cells[0].String() != "" {
			rowLen++
		}
	}

	colArr := []int{}
	colType := make(map[int]string)
	colField := make(map[int]string)
	for i := 0; i < celLen; i++ {
		if sheet.Rows[0].Cells[i].String() != "" {
			colArr = append(colArr, i)
			colType[i] = sheet.Rows[1].Cells[i].String()
			colField[i] = sheet.Rows[2].Cells[i].String()
		}
	}
	rbody := []interface{}{}
	for idxRow := 0; idxRow < rowLen; idxRow++ {
		row := sheet.Rows[idxRow]
		if idxRow < 4 {
			continue
		}

		d := make(map[string]interface{})

		for _, c := range colArr {
			var data interface{}
			if colType[c] == "String" || colType[c] == "string" {
				data = row.Cells[c].String()
			}
			if colType[c] == "int" || colType[c] == "Int" {
				data, _ = row.Cells[c].Int()
			}
			if colType[c] == "Array" || colType[c] == "array" {
				s := row.Cells[c].String()
				arrResult := []int{}
				if s != "" {
					arr := strings.Split(s, "~")
					for _, item := range arr {
						arrResult = append(arrResult, StringToInt(item))
					}
				}
				data = arrResult
			}
			d[colField[c]] = data
		}
		rbody = append(rbody, d)
	}
	sbyte, _ := json.Marshal(rbody)
	ioutil.WriteFile(serverPath+jsonName, sbyte, 0666)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
