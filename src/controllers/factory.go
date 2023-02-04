package controllers

import (
	"github.com/matheus-osorio/go-term-solver/src/controllers/filter"
	"github.com/matheus-osorio/go-term-solver/src/controllers/getters"
)

func ControllerFactory(wordSize int, language string) WordleSolver {
	return WordleSolver{
		Getter:   getters.CreateGetter(uint(wordSize), language),
		Score:    nil,
		Filter:   filter.WordFilter{},
		WordSize: wordSize,
	}
}
