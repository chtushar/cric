package main

import (
	"bytes"
	"cricli/model"
	"cricli/utils"
	"encoding/json"
	"fmt"

	xj "github.com/basgys/goxml2json"
)

func main() {
	xmlBytes, _ := utils.GetXML("http://static.espncricinfo.com/rss/livescores.xml")
	data := bytes.NewReader(xmlBytes)
	js, err := xj.Convert(data)

	page := &model.Page{}
	err = json.Unmarshal(js.Bytes(), page)

	items := []model.Item{}
	items = utils.GetAllItems(page)

	if err != nil {
		fmt.Println(err)
	} else {
		// _ = ioutil.WriteFile("test,json", json.Bytes(), 0644)
		// fmt.Println(items)
		for _, match := range items {
			fmt.Println(match.Title)
		}
	}
}
