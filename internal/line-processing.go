package handlerfile

import (
	"fmt"
	"strconv"
)

/*
	1. На вход поступает слайс с символами(символы строки).

	2. У нас есть индексы мин верхней строки(для первой строки нет), по этим индексам
	   увеличиваем значения соответствующих ячеек обрабатываемой строки.

	3. Итерируемся по строке, если встречаем мину - увеличиваем значения соседей, сохраняем индексы мин.

	4. Нам известны индексы мин строки, повторяем пукт 2, но наоборот(по индексам только что обработанной строки
	   обрабатывем верхнюю) и выводим в консоль верхнюю строку. Нынешняя становится верхней.
*/

var (
	indxMine   []int    // индексы мин предыдущей строки
	lineBuffer []string // буфер для предыдущей строки,
)

/*
	В памяти 2 слайса строк и 1 слайс для индексов мин прошлой строки
*/
func processingLine(line []string) {

	line = indxProcessing(line, indxMine) // обработка текущей строки по индексам мин прошлой строки
	indxMine = indxMine[:0]               // очищаем слайс с индексами мин верхней строки

	// итерация по слайсу с инкрементированием соседних ячеек при встрече мины
	for indx, character := range line {
		if character == "O" {
			line[indx] = "0"
		}
		if character == "X" {
			indxMine = append(indxMine, indx)
			if indx-1 >= 0 {
				line[indx-1] = Oplus(line[indx-1])
			}
			if indx+1 < len(line) {
				line[indx+1] = Oplus(line[indx+1])
			}
		}
	}

	// обработка прошлой строки по индексам мин текущей строки и вывод в консоль
	if len(lineBuffer) != 0 { // если буфер не пустой
		lineBuffer = indxProcessing(lineBuffer, indxMine)
		stdoutLine(lineBuffer)
	}
	lineBuffer = line // обработанную строку кладем в буфер
}

// функция обработки по индексам
func indxProcessing(line []string, indx []int) []string {
	for _, mineNumber := range indx {
		line[mineNumber] = Oplus(line[mineNumber]) // сам индекс
		if mineNumber-1 >= 0 {
			line[mineNumber-1] = Oplus(line[mineNumber-1]) // ячейка левее
		}
		if mineNumber+1 < len(line) {
			line[mineNumber+1] = Oplus(line[mineNumber+1]) // ячейка правее
		}
	} 
	return line
}

// увеличивает значение для ячейки
func Oplus(cell string) string {
	if cell == "X" || cell == " " {
		return cell
	}

	if cell == "O" || cell == "0" {
		return "1"
	}

	// увеличиваем значение
	countMines, _ := strconv.Atoi(cell)
	return strconv.Itoa(countMines + 1) // (strconv быстрее fmt.Sprint)
}

// вывод в консоль
func stdoutLine(line []string) {
	for indx, character := range line {
		if indx != 0 {
			fmt.Print(" ")
		}
		fmt.Print(character)
	}
	fmt.Print("\n")
}
