package main

import (
	"flag"
	"os"

	handlerfile "github.com/Minesweeper-loader/internal"
	"github.com/sirupsen/logrus"
)

var (
	filePath string // путь к файлу
)

func init() {
	// флаг для указания пути к файлу, по умолчанию путь ./exmpl.txt
	flag.StringVar(&filePath, "file-path", "./exmpl.txt", "path to file")
}

func main() {
	flag.Parse()  // парсинг флагов

	// открываем файл, мы будем только читать файл, поэтому Open вместо OpenFile
	file, err := os.Open(filePath)
	handlererror(err, "Can't open file")
	defer file.Close()

	// старт обработки файла
	handlerfile.StartFileProcessing(file)
	handlererror(err, "Error with processing data from file")
}

// функция для обработки ошибок(стилистическая вещь, код становится более читабельным, без колбасок)
func handlererror(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s : %s", msg, err)
	}
}
