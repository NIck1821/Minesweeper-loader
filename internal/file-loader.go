package handlerfile

import (
	"bufio"
	"io"
	"os"
	"strings"
)

/*
	Начало обрабоки, читаем строки и передаем в функцию
	processingLine для обрабоки.
*/
func StartFileProcessing(file *os.File) error {
	// объявляю переменные
	var (
		line  string
		cells []string
		err   error
		buf   *bufio.Reader
	)

	/*
		Читать файлы можно тремя способами:

		Прочитать весь файл в память
		Читать в байтах
		Читать по строке

		Все в память я выгружать не собираюсь, файл может быть большим,
		также мне необходимо разделение на строки,
		поэтому читаю построчно.
	*/
	buf = bufio.NewReader(file)
	for {
		// использую ReadString из пакета bufio,
		// Поскольку Scanner.Scan () ограничен размером буфера в 4096(~65 500 символов) байт на строку.
		// ReadBytes не ограничен размером буфера, ReadString оболочка над ReadBytes(но возвр string)
		line, err = buf.ReadString('\n')
		// Обрезаю \n в конце строки
		line = strings.TrimSpace(line)
		// получаю срез символов (cells = ячейки)
		cells = strings.Split(line, " ")

		// передаю обрабатывающей функции символы первой строки
		processingLine(cells)

		// обработка ошибок
		if err != nil {
			// если чтение файла закончено, то выхожу из цикла
			if err == io.EOF {
				break
			} else {
				// если ошибка, то передаю вышестоящему обработчику
				// мы не можем прочитать строку, это критично,
				// просто пропустить или вывести лог нельзя
				// поэтому в main прога остановится с выводом ошибки
				return err
			}
		}

	}
	stdoutLine(lineBuffer) // очистка буфера со строкой

	return nil
}
