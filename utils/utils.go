package utils

import (
	"cricli/data"
	"cricli/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	jq "github.com/antchfx/jsonquery"
)

func GetXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func GetAllItems(page *model.Page) (items []model.Item) {
	return page.Rss.Channel.Item
}

func ExitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func getTeamName(teamAlias string) (name string, err error) {
	doc, err := jq.Parse(strings.NewReader(data.Teams))
	if err != nil {
		return "", err
	}
	if n := jq.FindOne(doc, "*/" + teamAlias); n != nil {
		return n.InnerText(), nil
	} else {
		return "", fmt.Errorf("Couldn't find the team")
	}
}

func GetScore(teamAlias string, items []model.Item) {

	team, err := getTeamName(teamAlias)

	if err != nil {
		ExitGracefully(err)
	}

	fmt.Println(team)
	for _, match := range items {
		fmt.Println(match.Title)
	}
}
