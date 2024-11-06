package client

import (
	"ais/internal/database"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func openDialogWindowEditComponent() {
	windowDialogEdit := AppFyne.NewWindow("Изменение компонента")

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
			widget.NewButton("Изменить", func() {
				err := database.UpdateComponent(conn)
				if err != nil {
					log.Println(err)
				}
			}),
			widget.NewButton("Отмена", func() {
				windowDialogEdit.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogEdit.Resize(fyne.NewSize(500, 300))
	windowDialogEdit.CenterOnScreen()
	windowDialogEdit.SetContent(content)
	windowDialogEdit.Show()
}

func openDialogWindowEditCustomer() {
	windowDialogEdit := AppFyne.NewWindow("Изменение клиента")

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
			widget.NewButton("Изменить", func() {
				err := database.UpdateComponent(conn)
				if err != nil {
					log.Println(err)
				}
			}),
			widget.NewButton("Отмена", func() {
				windowDialogEdit.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogEdit.Resize(fyne.NewSize(500, 300))
	windowDialogEdit.CenterOnScreen()
	windowDialogEdit.SetContent(content)
	windowDialogEdit.Show()
}

func openDialogWindowEditSale() {
	windowDialogEdit := AppFyne.NewWindow("Изменение продажи")

	form := widget.NewForm(
		widget.NewFormItem("Компонент", widget.NewEntry()),
		widget.NewFormItem("Клиент", widget.NewEntry()),
		widget.NewFormItem("Количество", widget.NewEntry()),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		form,
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Изменить", func() {
				err := database.UpdateComponent(conn)
				if err != nil {
					log.Println(err)
				}
			}),
			widget.NewButton("Отмена", func() {
				windowDialogEdit.Close()
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	windowDialogEdit.Resize(fyne.NewSize(500, 300))
	windowDialogEdit.CenterOnScreen()
	windowDialogEdit.SetContent(content)
	windowDialogEdit.Show()
}
