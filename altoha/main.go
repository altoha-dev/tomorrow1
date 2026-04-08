package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {

		fmt.Println("error")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	text, error := os.ReadFile(input)

	if error != nil {
		return
	}

	text = []byte(strings.ReplaceAll(string(text), "(hex)", " (hex) "))
	text = []byte(strings.ReplaceAll(string(text), "(bin)", " (bin) "))
	text = []byte(strings.ReplaceAll(string(text), "(hex )", " (hex) "))
	text = []byte(strings.ReplaceAll(string(text), "( hex)", " (hex) "))
	text = []byte(strings.ReplaceAll(string(text), "(,hex)", " (hex) "))
	text = []byte(strings.ReplaceAll(string(text), "(hex,)", " (hex) "))
	text = []byte(strings.ReplaceAll(string(text), "(,hex,)", " (hex) "))
	text = []byte(strings.ReplaceAll(string(text), "(bin )", " (hex) "))
	text = []byte(strings.ReplaceAll(string(text), "( bin)", " (bin) "))
	text = []byte(strings.ReplaceAll(string(text), "(,bin)", " (bin) "))
	text = []byte(strings.ReplaceAll(string(text), "(bin,)", " (bin) "))
	text = []byte(strings.ReplaceAll(string(text), "(,bin,)", " (bin) "))

	words := strings.Fields(string(text))

	fmt.Println(words)
	UpLow(words)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			// (hex)
			res, err := strconv.ParseInt(words[i-1], 16, 64)
			// 1E
			if err == nil {

				words[i-1] = fmt.Sprint(res) // 1E ->30 перезапись
				words[i] = ""
			}

		}
		if words[i] == "(bin)" && i > 0 {
			// (bin)
			res, err := strconv.ParseInt(words[i-1], 2, 64)
			// (10)
			if err == nil {

				words[i-1] = fmt.Sprint(res) // 10 ->2
				words[i] = ""
				// (bin del)
			}

		}

	}
	UpLow(words)

	// 1. Собираем слова в одну строку через пробел
	finalText := strings.Join(words, " ")

	// 2. Убираем лишние пробелы, которые остались от пустых ""
	// (Снова разбиваем и склеиваем — это самый простой способ почистить текст)
	cleanText := strings.Join(strings.Fields(finalText), " ")

	// 3. Записываем результат в файл
	os.WriteFile(output, []byte(cleanText), 0o644)

	fmt.Println(cleanText)
}
