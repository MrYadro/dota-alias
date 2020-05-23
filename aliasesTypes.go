package main

type Aliases []struct {
	ID      int      `json:"id"`
	Aliases []string `json:"aliases"`
}

type AliasesList []AliasesListItem

type AliasesListItem struct {
	EntityType string   `json:"type"`
	ID         int      `json:"id"`
	Aliases    []string `json:"aliases"`
}
