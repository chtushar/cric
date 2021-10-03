package utils

import (
	"cric/data"
	"cric/model"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	jq "github.com/antchfx/jsonquery"
	"github.com/ttacon/chalk"
	"golang.org/x/net/html"
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
	if n := jq.FindOne(doc, "*/"+teamAlias); n != nil {
		return n.InnerText(), nil
	} else {
		return "", fmt.Errorf("Couldn't find the team.")
	}
}

func PrintLiveMatches(items []model.Item) {
	for _, match := range items {
		if strings.Contains(match.Title, "*") {
			fmt.Printf("🏏 " + match.Title + "\n\n")
		}
	}
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func GetHtmlTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}

func getPlayersAndOvers(details string, score *model.Score) {
	re := regexp.MustCompile(`\((.*?)\)`)
	playersAndOvers := re.FindString(details)

	playersAndOvers = strings.Trim(playersAndOvers, "(")
	playersAndOvers = strings.Trim(playersAndOvers, ")")

	extracted := strings.Split(playersAndOvers, ", ")
	if len(extracted) >= 3{
		*&score.Overs = extracted[0]
		if len(extracted) == 4 {
			*&score.Players.Batters = extracted[1:3]
			*&score.Players.Bowler = extracted[3]
		} else if len(extracted) == 3 {
			*&score.Players.Batters = extracted[1:2]
			*&score.Players.Bowler = extracted[2]
		}
	}
}
func PrintMatchDetails(match model.Item) {
	score := &model.Score{}

	htmlData, _ := http.Get(match.Link)
	details, _ := GetHtmlTitle(htmlData.Body)
	getPlayersAndOvers(details, score)

	fmt.Println(chalk.Yellow.Color(chalk.Bold.TextStyle("🏏💻 cric")))

	for i := 0; i < len(match.Title)+3; i++ {
		fmt.Printf("-")
	}

	fmt.Printf("\n\n" + "⚔️  " + chalk.Bold.TextStyle(match.Title) + "\n\n")

	if len(score.Overs) > 0 {
		fmt.Printf(chalk.Magenta.Color("Overs: ") + chalk.Bold.TextStyle(score.Overs) + "\n")
	}

	if len(score.Players.Batters) >= 1 {
		fmt.Printf(chalk.Magenta.Color("Batters: "))
		for i := range score.Players.Batters {
			fmt.Printf(chalk.Bold.TextStyle(score.Players.Batters[i]) + "\t")
		}
	}

	if len(score.Players.Bowler) > 0 {
		fmt.Printf("\n" + chalk.Magenta.Color("Bowler: "))
		fmt.Printf(chalk.Bold.TextStyle(score.Players.Bowler) + "\n\n")
	}

	for i := 0; i < len(match.Title)+3; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n\n")

}

func GetScore(teamAlias string, items []model.Item) {

	team, err := getTeamName(teamAlias)

	if err != nil {
		fmt.Printf(chalk.Magenta.Color("Sorry, cric couldn't find the team you are looking for but here are some of the live scores. Hope it helps. 😬\n\n\n"))

		PrintLiveMatches(items)
		os.Exit(1)
	}

	for _, match := range items {
		matchTitleInLowercase := strings.ToLower(match.Title)
		if strings.Contains(matchTitleInLowercase, team) {
			PrintMatchDetails(match)
		}
	}
}
