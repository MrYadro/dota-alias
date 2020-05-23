package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func findInJSON(aliases Aliases, id int) []string {
	var out []string
	for _, alias := range aliases {
		if alias.ID == id {
			out = alias.Aliases
		}
	}
	return out
}

func main() {
	var players ODotaProPlayers
	var teams ODotaTeams
	var leagues ODotaLeagues

	getJSON(playersURL, &players)
	getJSON(teamsURL, &teams)
	getJSON(leaguesURL, &leagues)

	var playersAliases Aliases
	var teamsAliases Aliases
	var leaguesAliases Aliases

	loadJSON("./aliases/players.json", &playersAliases)
	loadJSON("./aliases/teams.json", &teamsAliases)
	loadJSON("./aliases/leagues.json", &leaguesAliases)

	textAliases := AliasesList{}

	for _, player := range players {
		pID := player.AccountID
		pName := strings.ToLower(player.Name)
		pAliases := findInJSON(playersAliases, pID)
		pAliases = append(pAliases, pName)
		alias := AliasesListItem{
			ID:         pID,
			EntityType: playerTag,
			Aliases:    pAliases,
		}
		textAliases = append(textAliases, alias)
	}

	for _, team := range teams {
		tID := team.TeamID
		tName := strings.ToLower(team.Name)
		tAliases := findInJSON(teamsAliases, tID)
		tAliases = append(tAliases, tName)
		alias := AliasesListItem{
			ID:         tID,
			EntityType: teamTag,
			Aliases:    tAliases,
		}
		textAliases = append(textAliases, alias)
	}

	for _, league := range leagues {
		lID := league.Leagueid
		lName := strings.ToLower(league.Name)
		lAliases := findInJSON(leaguesAliases, lID)
		lAliases = append(lAliases, lName)
		alias := AliasesListItem{
			ID:         lID,
			EntityType: leagueTag,
			Aliases:    lAliases,
		}
		textAliases = append(textAliases, alias)
	}

	jsonString, _ := json.Marshal(textAliases)
	ioutil.WriteFile("dict.json", jsonString, os.ModePerm)
}
