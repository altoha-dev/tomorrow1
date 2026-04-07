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

	result := string(content) + " " + "aboba"

	//  os.WriteFile(имя_файла, данные_в_байтах, 0644)
	os.WriteFile(Out, []byte(result), 0o644)

	fmt.Println(string(content))
	fmt.Println(result)

	word := strings.Fields(string(content))
	fmt.Println(word)
}
