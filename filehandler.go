package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	//"path/filepath"
	"strings"
)

func FileHandler() {
	// Читаем содержимое родительской папки
	files, err := os.ReadDir(inputPath)
	if err != nil {
		log.Fatalf("Ошибка чтения директории: %v", err)
	}

	nameFilter := strings.TrimSpace("save")
	nameFilterWithout := strings.TrimSpace("_mp")

	for _, file := range files {
		if strings.Contains(file.Name(), nameFilter) && !strings.Contains(file.Name(), nameFilterWithout) {
			fmt.Println(file.Name())

			data, err := os.ReadFile(inputPath + "/" + file.Name())
			if err != nil {
				log.Fatalf("Ошибка при чтении файла: %v", err)
			}

			// Декодируем JSON из текста
			var save SaveData
			err = json.Unmarshal(data, &save)
			if err != nil {
				log.Fatalf("Ошибка при парсинге JSON: %v", err)
			}

			// Выводим имя из stats
			if len(save.Players) > 0 {
				dataName = append(dataName, save.Players)
				fmt.Println("Имя:", save.Players[0].Stats.Name)
				fmt.Println("Уровень:", save.Players[0].Stats.LVL)
				fmt.Println("Золото:", save.Players[0].Stats.GOLD)
			} else {
				fmt.Println("Нет игроков в JSON.")
			}
		}
	}
}
