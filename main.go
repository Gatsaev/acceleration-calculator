package main // Убедитесь, что это первая строка файла!

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Расчёт ускорения")

	// Загрузка изображения для фона (formula.png должен быть в папке)
	img := canvas.NewImageFromFile("formula.png")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(300, 300))

	// Поля ввода
	initialSpeedEntry := widget.NewEntry()
	initialSpeedEntry.SetPlaceHolder("Начальная скорость (u, м/с)")

	finalSpeedEntry := widget.NewEntry()
	finalSpeedEntry.SetPlaceHolder("Конечная скорость (v, м/с)")

	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("Время (t, с)")

	// Метка для результата
	resultLabel := widget.NewLabel("Результат: ")

	// Кнопка "Расчёт"
	calculateButton := widget.NewButton("Расчёт", func() {
		u, errU := strconv.ParseFloat(initialSpeedEntry.Text, 64)
		v, errV := strconv.ParseFloat(finalSpeedEntry.Text, 64)
		t, errT := strconv.ParseFloat(timeEntry.Text, 64)

		if errU != nil || errV != nil || errT != nil || t == 0 {
			resultLabel.SetText("Ошибка: Введите корректные числа (t ≠ 0)")
			return
		}

		a := (v - u) / t
		resultLabel.SetText(fmt.Sprintf("Ускорение: %.2f м/с²", a))
	})

	// Контейнер для элементов ввода
	inputContainer := container.NewVBox(
		initialSpeedEntry,
		finalSpeedEntry,
		timeEntry,
		calculateButton,
		resultLabel,
	)

	// Основной контейнер: ввод сверху, фон в центре
	mainContainer := container.NewBorder(
		inputContainer, nil, nil, nil,
		container.NewCenter(img),
	)

	// Замена устаревшего NewMax на современный подход
	myWindow.SetContent(container.NewMax(
		canvas.NewRectangle(color.RGBA{R: 240, G: 240, B: 240, A: 255}),
		mainContainer,
	))

	myWindow.Resize(fyne.NewSize(400, 500))
	myWindow.ShowAndRun()
}
