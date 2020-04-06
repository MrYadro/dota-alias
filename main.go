package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	playersURL = "https://api.opendota.com/api/proPlayers"
	teamsURL   = "https://api.opendota.com/api/teams"
	leaguesURL = "https://api.opendota.com/api/leagues"
	playerTag  = "игрок"
	teamTag    = "команда"
	leagueTag  = "лига"
)

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func loadJSON(path string, target interface{}) error {
	file, _ := ioutil.ReadFile(path)
	_ = json.Unmarshal([]byte(file), target)

	return nil
}

func main() {
	var players ODotaProPlayers
	var teams ODotaTeams
	var leagues ODotaLeagues

	var dict string

	getJSON(playersURL, &players)
	getJSON(teamsURL, &teams)
	getJSON(leaguesURL, &leagues)

	for _, player := range players {
		dictEl := fmt.Sprintf("%s,%s\n", player.Personaname, playerTag)
		dict += dictEl
	}

	for _, team := range teams {
		dictEl := fmt.Sprintf("%s,%s\n", team.Name, teamTag)
		dict += dictEl
	}

	for _, league := range leagues {
		dictEl := fmt.Sprintf("%s,%s\n", league.Name, leagueTag)
		dict += dictEl
	}

	var playersAliases Aliases
	var teamsAliases Aliases
	var leaguesAliases Aliases

	loadJSON("./aliases/players.json", &playersAliases)
	loadJSON("./aliases/teams.json", &teamsAliases)
	loadJSON("./aliases/leagues.json", &leaguesAliases)

	for _, player := range playersAliases {
		for _, playerAlias := range player.Aliases {
			dictEl := fmt.Sprintf("%s,%s\n", playerAlias, playerTag)
			dict += dictEl
		}
	}

	for _, team := range teamsAliases {
		for _, teamAlias := range team.Aliases {
			dictEl := fmt.Sprintf("%s,%s\n", teamAlias, teamTag)
			dict += dictEl
		}
	}

	for _, league := range leaguesAliases {
		for _, leagueAlias := range league.Aliases {
			dictEl := fmt.Sprintf("%s,%s\n", leagueAlias, leagueTag)
			dict += dictEl
		}
	}

	f, err := os.Create("dict.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(dict)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
