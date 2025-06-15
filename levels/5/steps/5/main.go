package main

import (
	"cmp"
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	return Player{
		Name:    name,
		Goals:   goals,
		Assists: assists,
		Misses:  misses,
		Rating:  calculateRating(goals, misses, assists),
	}
}

func goalsSort(players []Player) []Player {
	sorted := slices.Clone(players)

	slices.SortFunc(sorted, func(a, b Player) int {
		if res := cmp.Compare(b.Goals, a.Goals); res != 0 {
			return res
		}
		return cmp.Compare(a.Name, b.Name)
	})

	return sorted
}

func ratingSort(players []Player) []Player {
	sorted := slices.Clone(players)

	slices.SortFunc(sorted, func(a, b Player) int {
		if res := cmp.Compare(b.Rating, a.Rating); res != 0 {
			return res
		}

		return cmp.Compare(a.Name, b.Name)
	})

	return sorted
}

func gmSort(players []Player) []Player {
	sorted := slices.Clone(players)

	slices.SortFunc(sorted, func(a, b Player) int {
		aV := average(a.Goals, a.Misses)
		bV := average(b.Goals, b.Misses)

		if res := cmp.Compare(bV, aV); res != 0 {
			return res
		}

		return cmp.Compare(a.Name, b.Name)
	})

	return sorted
}

func calculateRating(goals, misses, assists int) float64 {
	goalsToAssists := float64(goals) + float64(assists)/2

	if misses == 0 {
		return goalsToAssists
	}

	return goalsToAssists / float64(misses)
}

func average(goals, misses int) float64 {
	if misses == 0 {
		return float64(goals)
	}
	return float64(goals) / float64(misses)
}
