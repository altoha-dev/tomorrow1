package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Твоя задача: проверить, сколько аргументов передал пользователь.
	// Если их меньше 3 (программа + 2 файла), нужно вывести ошибку и выйти.

	if len(os.Args) != 3 {

		fmt.Println("Ошибка недостаточно аргументов")

		return
	}

	In := os.Args[1]
	Out := os.Args[2]

	fmt.Println("Данные приняты:", In, "", "и ", Out)

	// переменная_с_данными, переменная_ошибки := функция("аргумент")
	content, error := os.ReadFile(In)

	if error != nil {
		fmt.Println("Error not correct inputFile")
		return
	}

	// var result string

	word := strings.Fields(string(content))

	for i := 0; i < len(word); i++ {

		// 1. ПРОВЕРКА НА (up)
		if word[i] == "(up)" && i > 0 {
			// ТВОЙ КОД:
			// Примени strings.ToUpper к word[i-1]
			// Очисти word[i], чтобы маркер исчез
			word[i-1] = strings.ToUpper(word[i-1])
			word[i] = ""

		}

		// 2. ПРОВЕРКА НА (low)
		if word[i] == "(low)" && i > 0 {
			// ТВОЙ КОД:
			// Примени strings.ToLower к word[i-1]
			// Очисти word[i]
			word[i-1] = strings.ToLower(word[i-1])
			word[i] = ""
		}

		// 3. ПРОВЕРКА НА (cap)
		if word[i] == "(cap)" && i > 0 {
			// ТВОЙ КОД:
			// Примени strings.Title к word[i-1]
			// Очисти word[i]
			word[i-1] = strings.Title(word[i-1])
			word[i] = ""

		}
	}

	// 1. Склеиваем слова обратно в одну строку через пробел
	// ВАЖНО: strings.Join вернет строку, где на месте (up) будут лишние пробелы
	tempResult := strings.Join(word, " ")

	// 2. Делаем "чистку" от лишних пробелов:
	// strings.Fields уберет все пустые строки, которые мы создали через word[i] = ""
	cleanWords := strings.Fields(tempResult)
	finalText := strings.Join(cleanWords, " ")

	// 3. ЗАПИСЫВАЕМ в файл (тот самый Out, который мы получили из os.Args[2])
	//  os.WriteFile(имя_файла, данные_в_байтах, 0644)
	err := os.WriteFile(Out, []byte(finalText), 0o644)
	if err != nil {
		fmt.Println("Ошибка при записи:", err)
		return
	}

	fmt.Println("Готово! Проверь файл:", Out)
}
