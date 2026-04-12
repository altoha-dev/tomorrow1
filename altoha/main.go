package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Работа с аргументами 
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// 2. Чтение файла
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return
	}

	// 3. Обработка
	result := processText(string(data))

	// 4. Запись результата
	err = os.WriteFile(outputFile, []byte(result), 0o644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}

func processText(input string) string {
	// Подготовка: разлепляем скобки, чтобы strings.Fields их увидел
	input = strings.ReplaceAll(input, "(", " (")
	input = strings.ReplaceAll(input, ")", ") ")

	words := strings.Fields(input)

	for i := 0; i < len(words); i++ {
		switch {
		// --- Системы счисления ---
		case words[i] == "(hex)":
			if i > 0 {
				val, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = fmt.Sprint(val)
					words[i] = ""
				}
			}

		case words[i] == "(bin)":
			if i > 0 {
				val, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = fmt.Sprint(val)
					words[i] = ""
				}
			}

		// --- Регистр (Одиночный) ---
		case words[i] == "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words[i] = ""
			}
		case words[i] == "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words[i] = ""
			}
		case words[i] == "(cap)":
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
				words[i] = ""
			}

		// --- Регистр (С числом n) ---
		case strings.HasPrefix(words[i], "(up,"):
			applyToMultiple(words, i, "up")
		case strings.HasPrefix(words[i], "(low,"):
			applyToMultiple(words, i, "low")
		case strings.HasPrefix(words[i], "(cap,"):
			applyToMultiple(words, i, "cap")
		}
	}

	// Сборка обратно и финальная чистка пунктуации
	return finalizePunctuation(words)
}

// Вспомогательная функция для сложных тегов (up, n)
func applyToMultiple(words []string, index int, action string) {
	if index+1 >= len(words) {
		return
	}
	// Используем Trim для извлечения числа
	numStr := strings.Trim(words[index+1], "(),")
	n, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}

	for j := 1; j <= n; j++ {
		target := index - j
		if target >= 0 {
			switch action {
			case "up":
				words[target] = strings.ToUpper(words[target])
			case "low":
				words[target] = strings.ToLower(words[target])
			case "cap":
				words[target] = capitalize(words[target])
			}
		}
	}
	words[index], words[index+1] = "", ""
}

// Функция для (cap)
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// Финальная сборка и логика знаков препинания
func finalizePunctuation(words []string) string {
	// Сначала склеиваем всё, убирая пустые строки, которые мы оставили вместо тегов
	var filtered []string
	for _, w := range words {
		if w != "" {
			filtered = append(filtered, w)
		}
	}

	result := strings.Join(filtered, " ")

	// Тут можно добавить логику замены " , " на ", " и так далее
	// Для идеального результата ТЗ требует аккуратной работы со знаками
	puncs := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range puncs {
		result = strings.ReplaceAll(result, " "+p, p)
	}

	// Логика для артикля "a/an"
	finalWords := strings.Fields(result)
	for i := 0; i < len(finalWords)-1; i++ {
		lowered := strings.ToLower(finalWords[i])
		if lowered == "a" {
			nextWord := strings.ToLower(finalWords[i+1])
			if strings.ContainsAny(string(nextWord[0]), "aeiouh") {
				if finalWords[i] == "A" {
					finalWords[i] = "An"
				} else {
					finalWords[i] = "an"
				}
			}
		}
	}

	return strings.Join(finalWords, " ")
}
