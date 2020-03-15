package main

import "time"

type ODotaProPlayers []struct {
	AccountID       int         `json:"account_id"`
	Steamid         string      `json:"steamid"`
	Avatar          string      `json:"avatar"`
	Avatarmedium    string      `json:"avatarmedium"`
	Avatarfull      string      `json:"avatarfull"`
	Profileurl      string      `json:"profileurl"`
	Personaname     string      `json:"personaname"`
	LastLogin       interface{} `json:"last_login"`
	FullHistoryTime time.Time   `json:"full_history_time"`
	Cheese          int         `json:"cheese"`
	FhUnavailable   bool        `json:"fh_unavailable"`
	Loccountrycode  string      `json:"loccountrycode"`
	LastMatchTime   time.Time   `json:"last_match_time"`
	Plus            bool        `json:"plus"`
	Name            string      `json:"name"`
	CountryCode     string      `json:"country_code"`
	FantasyRole     int         `json:"fantasy_role"`
	TeamID          int         `json:"team_id"`
	TeamName        string      `json:"team_name"`
	TeamTag         string      `json:"team_tag"`
	IsLocked        bool        `json:"is_locked"`
	IsPro           bool        `json:"is_pro"`
	LockedUntil     interface{} `json:"locked_until"`
}

type ODotaTeams []struct {
	TeamID        int     `json:"team_id"`
	Rating        float64 `json:"rating"`
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	LastMatchTime int     `json:"last_match_time"`
	Name          string  `json:"name"`
	Tag           string  `json:"tag"`
	LogoURL       string  `json:"logo_url"`
}

type ODotaLeagues []struct {
	Leagueid int         `json:"leagueid"`
	Ticket   interface{} `json:"ticket"`
	Banner   interface{} `json:"banner"`
	Tier     string      `json:"tier"`
	Name     string      `json:"name"`
}
