package client

import (
	"ais/internal/database"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func openDialogWindowAddComponent() {
	windowDialogAdd := AppFyne.NewWindow("Добавление компонента")

	form := widget.NewForm(
		widget.NewFormItem("Компонент", widget.NewEntry()),
		widget.NewFormItem("Описание", widget.NewEntry()),
		widget.NewFormItem("Цена", widget.NewEntry()),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Добавить", func() {
				err := database.CreateComponent(conn)
				if err != nil {
					log.Println(err)
				}
			}),
			widget.NewButton("Отмена", func() {
				windowDialogAdd.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogAdd.Resize(fyne.NewSize(500, 300))
	windowDialogAdd.CenterOnScreen()
	windowDialogAdd.SetContent(content)
	windowDialogAdd.Show()
}

func openDialogWindowAddCustomer() {
	windowDialogAdd := AppFyne.NewWindow("Добавление клиента")

	form := widget.NewForm(
		widget.NewFormItem("Имя", widget.NewEntry()),
		widget.NewFormItem("Телефон", widget.NewEntry()),
		widget.NewFormItem("Почта", widget.NewEntry()),
		widget.NewFormItem("Адрес", widget.NewEntry()),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Добавить", func() {
				err := database.CreateComponent(conn)
				if err != nil {
					log.Println(err)
				}
			}),
			widget.NewButton("Отмена", func() {
				windowDialogAdd.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogAdd.Resize(fyne.NewSize(500, 300))
	windowDialogAdd.CenterOnScreen()
	windowDialogAdd.SetContent(content)
	windowDialogAdd.Show()
}

func openDialogWindowAddSale() {
	windowDialogAdd := AppFyne.NewWindow("Добавление продажи")

	form := widget.NewForm(
		widget.NewFormItem("Компонент", widget.NewEntry()),
		widget.NewFormItem("Клиент", widget.NewEntry()),
		widget.NewFormItem("Количество", widget.NewEntry()),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Добавить", func() {
				err := database.CreateComponent(conn)
				if err != nil {
					log.Println(err)
				}
			}),
			widget.NewButton("Отмена", func() {
				windowDialogAdd.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogAdd.Resize(fyne.NewSize(500, 300))
	windowDialogAdd.CenterOnScreen()
	windowDialogAdd.SetContent(content)
	windowDialogAdd.Show()
}
