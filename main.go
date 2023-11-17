package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type Param struct {
	DashboardTitle string
}

var PARAM *Param

func init() {
	PARAM = &Param{}

	flag.StringVar(&PARAM.DashboardTitle, "title", "[Development][Merchant Order Management] New Dashboard", "Dashboard Title")

	flag.Parse()
}

func main() {
	sTemplate, err := readFile("template.yaml")
	if err != nil {
		log.Fatalln("Err:", err)
	}

	strData, err := readCsvFile("data.csv")
	if err != nil {
		log.Fatalln("Err:", err)
	}
	data := buildStrData(strData)

	builder := &DashboardBuilder{
		DashboardTitle: PARAM.DashboardTitle,
		StrTemplate:    sTemplate,
		SourceData:     data,
	}
	builder.ProcessData()

	if buff, err := builder.Build(); err != nil {
		fmt.Println("Err:", err)
	} else {
		temp := make(map[string]interface{})
		if err := yaml.Unmarshal(buff.Bytes(), &temp); err != nil {
			log.Fatalln("Err:", err)
		}

		jResult, jErr := json.Marshal(temp)
		if jErr != nil {
			log.Fatalln("Err:", err)
		}
		fmt.Println(string(jResult))

		// fmt.Println(buff.String())
	}
}
