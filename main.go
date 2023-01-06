package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Посчитать зарплату")
	w.Resize(fyne.NewSize(500, 600))

	//Создание темной темы
	a.Settings().SetTheme(theme.DarkTheme())

	//оклад
	label1 := widget.NewLabel("Введи оклад")
	entry1 := widget.NewEntry()

	//количество часов в месяце
	label2 := widget.NewLabel("Количество часов в месяце")
	entry2 := widget.NewEntry()

	//премия%
	label3 := widget.NewLabel("Премия(введи проценты)")
	entry3 := widget.NewEntry()

	//кол-во по коф. 1
	label4 := widget.NewLabel("Введи часы по кофициенту 1")
	entry4 := widget.NewEntry()

	//кол-во по коф. 1.3
	label5 := widget.NewLabel("Введи часы по кофициенту 1.3 (если нет! Введи 0)")
	entry5 := widget.NewEntry()

	//кол-во по коф. 1.5
	label6 := widget.NewLabel("Введи часы по кофициенту 1.5 (если нет! Введи 0)")
	entry6 := widget.NewEntry()

	//сумма
	answer_dirty := widget.NewLabel("")
	answer_4ist := widget.NewLabel("")

	btn := widget.NewButton("Посчитать", func() {

		num1, err1 := strconv.ParseFloat(entry1.Text, 64)

		num2, err2 := strconv.ParseFloat(entry2.Text, 64)

		num3, err3 := strconv.ParseFloat(entry3.Text, 64)

		num4, err4 := strconv.ParseFloat(entry4.Text, 64)

		num5, err5 := strconv.ParseFloat(entry5.Text, 64)

		num6, err6 := strconv.ParseFloat(entry6.Text, 64)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			answer_dirty.SetText("Ошибка ввода")
		} else {

			//Стоимость часа
			cost_per_hour := num1 / num2

			//Часы коф. 1
			coefficient_1 := cost_per_hour * num4

			//Часы коф. 1.3
			coefficient_1_3 := (cost_per_hour * num5) * 1.3

			//Часы коф. 1.5
			coefficient_1_5 := (cost_per_hour * num6) * 1.5

			//Премия
			prize := cost_per_hour * (num4 * (num3 / 100))

			//Сумма грязными
			sum_dirty := (coefficient_1 + coefficient_1_3 + coefficient_1_5 + prize)
			answer_dirty.SetText(fmt.Sprintf("%.2f      Сумма без вычета", sum_dirty))

			//Сумма чистыми
			sum_4ist := (coefficient_1 + coefficient_1_3 + coefficient_1_5 + prize) * 0.87
			answer_4ist.SetText(fmt.Sprintf("%.2f      Сумма чистыми", sum_4ist))

		}
	})

	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		label3,
		entry3,
		label4,
		entry4,
		label5,
		entry5,
		label6,
		entry6,
		btn,
		answer_dirty,
		answer_4ist,
	))

	w.ShowAndRun()
}
