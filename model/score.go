package model

type Score struct {
	Overs   string
	Players Players
}

type Players struct {
	Batters []string
	Bowler  string
}
