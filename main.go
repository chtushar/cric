package main

import (
	"bytes"
	"cricli/model"
	"cricli/utils"
	"encoding/json"
	"fmt"
	"os"

	xj "github.com/basgys/goxml2json"
)

const LIVE_SCORES = "http://static.espncricinfo.com/rss/livescores.xml"

func main() {
	xmlBytes, _ := utils.GetXML(LIVE_SCORES)
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
		team := os.Args[1]

		utils.GetScore(team, items)
	}
}
