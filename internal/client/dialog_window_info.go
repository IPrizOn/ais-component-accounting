package client

import (
	"ais/internal/database"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func openDialogWindowError() {
	windowDialog := AppFyne.NewWindow("Ошибка")

	content := container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("У вас недостаточно прав"),
			layout.NewSpacer(),
		),
		widget.NewButton("ОК", func() {
			windowDialog.Close()
		}),
	)

	windowDialog.Resize(fyne.NewSize(250, 75))
	windowDialog.CenterOnScreen()
	windowDialog.SetContent(content)
	windowDialog.Show()
}

func openDialogWindowConfirm(selectedTab string, id int) {
	windowDialog := AppFyne.NewWindow("Подтверждение")

	content := container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Вы точно хотите удалить эту запись?"),
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Да", func() {
				switch selectedTab {
				case "component":
					err := database.DeleteComponent(conn, id)
					if err != nil {
						log.Println(err)
					}
				case "customer":
					err := database.DeleteCustomer(conn, id)
					if err != nil {
						log.Println(err)
					}
				case "sale":
					err := database.DeleteSale(conn, id)
					if err != nil {
						log.Println(err)
					}
				}
				windowDialog.Close()
			}),
			widget.NewButton("Нет", func() {
				windowDialog.Close()
			}),
			layout.NewSpacer(),
		),
	)

	windowDialog.Resize(fyne.NewSize(250, 75))
	windowDialog.CenterOnScreen()
	windowDialog.SetContent(content)
	windowDialog.Show()
}
