package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {

		fmt.Println("LOH")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	text, error := os.ReadFile(input)

	if error != nil {
		return
	}

	words := strings.Fields(string(text))

	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			res, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {

				words[i-1] = fmt.Sprint(res)
				words[i] = ""
			}

		}
	}
	// 1. Собираем слова в одну строку через пробел
	finalText := strings.Join(words, " ")

	// 2. Убираем лишние пробелы, которые остались от пустых ""
	// (Снова разбиваем и склеиваем — это самый простой способ почистить текст)
	cleanText := strings.Join(strings.Fields(finalText), " ")

	// 3. Записываем результат в файл
	os.WriteFile(output, []byte(cleanText), 0o644)
}
