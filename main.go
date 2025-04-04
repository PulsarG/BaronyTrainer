// ** 0.01

package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
	//"path/filepath"
	"encoding/json"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2"
)

type Stats struct {
	Name string `json:"name"`
	LVL  int    `json:"LVL"`
	GOLD int    `json:"GOLD"`
}

type Player struct {
	Stats Stats `json:"stats"`
}

type SaveData struct {
	Players []Player `json:"players"`
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Barony Trainer")

	inputPath := strings.TrimSpace("G:/Barony/savegames")

	dataName := []string{}

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

			data, err := os.ReadFile("G:/Barony/savegames/" + file.Name())
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
				dataName = append(dataName, save.Players[0].Stats.Name)
				fmt.Println("Имя персонажа:", save.Players[0].Stats.Name)
				fmt.Println("Имя персонажа:", save.Players[0].Stats.LVL)
				fmt.Println("Имя персонажа:", save.Players[0].Stats.GOLD)
			} else {
				fmt.Println("Нет игроков в JSON.")
			}
		}
	}

	// Создаём список
	list := widget.NewList(
		func() int {
			return len(dataName)
		},
		func() fyne.CanvasObject {
			// Шаблон одного элемента списка
			return widget.NewLabel("Текст")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// Отображение элемента по индексу
			o.(*widget.Label).SetText(dataName[i])
		},
	)

	myWindow.SetContent(container.NewMax(list))
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}

/* package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
	//"path/filepath"
	"encoding/json"
	"strings"
)

type Stats struct {
	Name string `json:"name"`
	LVL  int    `json:"LVL"`
	GOLD int    `json:"GOLD"`
}

type Player struct {
	Stats Stats `json:"stats"`
}

type SaveData struct {
	Players []Player `json:"players"`
}

func main() {

	inputPath := strings.TrimSpace("G:/Barony/savegames")

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

			data, err := os.ReadFile("G:/Barony/savegames/" + file.Name())
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
				fmt.Println("Имя персонажа:", save.Players[0].Stats.Name)
				fmt.Println("Имя персонажа:", save.Players[0].Stats.LVL)
				fmt.Println("Имя персонажа:", save.Players[0].Stats.GOLD)
			} else {
				fmt.Println("Нет игроков в JSON.")
			}
		}
	}
} */
