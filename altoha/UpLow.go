package main

import (
	"strconv"
	"strings"
)

func UpLow(words []string) {
	for i := 0; i < len(words); i++ {
		// Проверяем на сложный тег (up, n)
		if strings.HasPrefix(words[i], "(up,") && i > 0 {

			// 1. Извлекаем число n из строки типа "3)"
			// words[i+1] это "3)"
			nStr := strings.Trim(words[i+1], ")")
			n, _ := strconv.Atoi(nStr)

			// 2. Цикл, который идет назад на n шагов
			for j := 1; j <= n; j++ {
				target := i - j
				if target >= 0 { // Защита: не уходим в минус
					words[target] = strings.ToUpper(words[target])
				}
			}

			// 3. Стираем тег и число
			words[i] = ""
			words[i+1] = ""
		}
	}
}
