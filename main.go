// ** 0.2

// !!! Разбить на модули

package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	//"path/filepath"

	//"path/filepath"
	"encoding/json"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2"
	configini "github.com/PulsarG/ConfigManager"
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

var isWindows bool = false

func main() {

	//configini.SaveToIni("file", "link", "G:/Barony/savegames")
	myApp := app.New()
	myWindow := myApp.NewWindow("Barony Trainer")

	dataName := []string{}
	inputPath := LinkFromSistem(isWindows)
	// Читаем содержимое родительской папки
	FileHaler(inputPath, &dataName)

	// Создаём список
	/* list := widget.NewList(
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
	) */

	// Массив для хранения полей ввода
	var entries []*widget.Entry

	// Контейнер с вертикальной компоновкой
	formContainer := container.NewVBox()

	// Заполняем поля ввода значениями из слайса
	for _, val := range dataName {
		entry := widget.NewEntry()
		entry.SetText(val)
		entries = append(entries, entry)
		formContainer.Add(entry)
	}

	// Кнопка для печати текущих значений в консоль
	saveButton := widget.NewButton("Сохранить", func() {
		for i, e := range entries {
			println("Поле", i, ":", e.Text)
		}
	})

	mainContent := container.NewVBox(formContainer, saveButton)

	myWindow.SetContent(mainContent)

	//myWindow.SetContent(container.NewMax(list))
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()
}

func FileHaler(inputPath string, dataName *[]string) {

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
				*dataName = append(*dataName, save.Players[0].Stats.Name)
				*dataName = append(*dataName, strconv.Itoa(save.Players[0].Stats.LVL))
				*dataName = append(*dataName, strconv.Itoa(save.Players[0].Stats.GOLD))

				fmt.Println("Имя персонажа:", save.Players[0].Stats.Name)
				fmt.Println("Имя персонажа:", save.Players[0].Stats.LVL)
				fmt.Println("Имя персонажа:", save.Players[0].Stats.GOLD)
			} else {
				fmt.Println("Нет игроков в JSON.")
			}
		}
	}
}

func LinkFromSistem(isWindows bool) string {
	var inputPath string
	if isWindows {
		inputPath = configini.GetFromIni("file", "link")
	} else {
		inputPath, _ = os.Getwd()
		// inputPath = filepath.Join(wd, "savegame0.baronysave")
	}

	return inputPath
}
