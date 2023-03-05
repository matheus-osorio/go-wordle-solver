package controllers

import (
	"github.com/matheus-osorio/go-term-solver/src/controllers/filter"
	"github.com/matheus-osorio/go-term-solver/src/controllers/getters"
)

func ControllerFactory(wordSize int, language string) WordleSolver {
	return WordleSolver{
		getter:   getters.CreateGetter(uint(wordSize), language),
		score:    nil,
		filter:   filter.WordFilter{},
		wordSize: wordSize,
	}
}
