package main
import (
	"fmt"
	"encoding/json"
	"github.com/tealeg/xlsx"
	"os"
	"io/ioutil"
	"excel_to_json/parseConfig"
	"strings"
	"archive/zip"
)
func main() {
	readPath()
}
func readPath(){
	println("**** START ****")
	var config = parseConfig.New("config.json")
	var itemList = config.Get("data").([]interface{})
	for _,v := range itemList {
		var plat = v.(map[string]interface{})
		name := plat["name"].(string)
		inPath := plat["inPath"].(string)
		serverOutPath := plat["serverOutPath"].(string)
		clientOutPath := plat["clientOutPath"].(string)

		serverZip := 0
		if plat["serverZip"] !=nil{
			serverZip = int(plat["serverZip"].(float64))
		}

		clientZip:=0
		if plat["clientZip"]!=nil{
			clientZip = int(plat["clientZip"].(float64))
		}

		println("\n **** PROCESS "+name+"**** \n")
		processAll(inPath,serverOutPath,clientOutPath,serverZip,clientZip)
	}
	fmt.Println("\n **** DONE ****")
	fmt.Print("\n Press 'Enter' to continue...\n")
	fmt.Scanln()
}

func processAll(inpath string,serverPath string,clientPath string,serverZip int,clientZip int){
	files,_:=ioutil.ReadDir(inpath)
	for _,file := range files{
		excelOp(inpath,file.Name(),serverPath,clientPath)
	}
	if serverZip==1{
		delZip(serverPath)
		createZip(serverPath)
	}
	if clientZip==1{
		delZip(clientPath)
		createZip(clientPath)
	}
}

func delZip(path string){
	println("clear old zip in path:"+path)
	files,_:=ioutil.ReadDir(path)
	for _,file := range files{
		if strings.HasSuffix(file.Name(),".zip"){
			os.Remove(path+file.Name())
		}
	}
}

func createZip(path string){
	println("create zipFile")
	f, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	fzip, _ := os.Create(path+"sdata.zip")
	w := zip.NewWriter(fzip)

	defer fzip.Close()
	defer w.Close()

	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(path + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		fw.Write(filecontent)
	}
}


func excelOp(path string,fileName string,serverPath string,clientPath string) {
	println("process "+path+" "+fileName)
	xlFile, err := xlsx.OpenFile(path+fileName)
	if err != nil {
		fmt.Println("open file error")
	}
	sheet := xlFile.Sheets[0]
	rowLen := len(sheet.Rows)

	celLen := len(sheet.Cols)
	var field= make([]string, celLen)
	var types= make([]string, celLen)

	var fieldClient= make([]interface{}, celLen)

	s := 0
	var rbody= make([]map[string]interface{}, rowLen-3)

	var cbody = make([][]interface{}, rowLen-2)
	cbody[0] = fieldClient

	for idxRow, row := range sheet.Rows {
		if idxRow == 0 || idxRow == 1 || idxRow == 2 {
			for cellIdx, cell := range row.Cells {
				text := strings.TrimSpace(cell.String())
				if idxRow == 0 {
					field[cellIdx] = text
					fieldClient[cellIdx] = text
					continue
				}
				if (idxRow == 1) {
					types[cellIdx] = text
					continue
				}
				if (idxRow == 2) {
					continue
				}
			}
			continue
		}
		t := make(map[string]interface{})
		var cValue = make([]interface{}, celLen)
		for cellIdx, cell := range row.Cells {
			if types[cellIdx] == "int" {
				v, _ := cell.Int64()
				t[field[cellIdx]] = v
				cValue[cellIdx] = v
			} else if types[cellIdx] == "string" {
				itemCell:=strings.TrimSpace(cell.String())
				t[field[cellIdx]] = itemCell
				cValue[cellIdx] = itemCell
			} else {
				v, _ := cell.Float()
				t[field[cellIdx]] = v
				cValue[cellIdx] = v
			}
		}

		cbody[s+1] = cValue
		rbody[s] = t
		s++
	}
	cbyte,_ := json.Marshal(cbody)
	sbyte,_ := json.Marshal(rbody)

	ioutil.WriteFile(clientPath+getOutputFileName(fileName),cbyte,0666)
	ioutil.WriteFile(serverPath+getOutputFileName(fileName),sbyte,0666)
}

func getOutputFileName(excelName string) string{
	arr := strings.Split(excelName,"-")
	var len = len(arr)
	if(len==1){
		r := strings.Split(excelName,".")
		return r[0]+".json"
	}
	r := strings.Split(arr[1],".")
	return r[0]+".json"
}