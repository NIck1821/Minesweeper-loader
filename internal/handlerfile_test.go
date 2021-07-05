package handlerfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Тест на обработку по индексам мин
*/
func TestIndx(t *testing.T) {
	// объявление и инициализация переменных
	var (
		line     []string
		indxMine []int
	)
	line = append(line, "O", "X", "O", "O", "O", "X", "O") // строка
	indxMine = append(indxMine, 1, 3)                      // индексы мин

	// обработка по индексам мин
	line = indxProcessing(line, indxMine)

	// сравнение результатов
	assert.Equal(t, line[0], "1")
	assert.Equal(t, line[2], "2")
	assert.Equal(t, line[3], "1")
	assert.Equal(t, line[4], "1")
}
