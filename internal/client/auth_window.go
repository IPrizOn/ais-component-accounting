package client

import (
	"ais/internal/database"
	"ais/internal/domain"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	userRole  string
	usersList []domain.Person
)

func openWindowAuth() {
	var err error

	usersList, err = database.GetPersonsList(conn)
	if err != nil {
		log.Println(err)
	}
	clear(usersList)

	windowMain.SetTitle("Авторизация")

	content := loadContent()

	windowMain.Resize(fyne.NewSize(300, 250))
	windowMain.CenterOnScreen()
	windowMain.SetContent(content)
}

func loadContent() *fyne.Container {
	labelLogin := widget.NewLabel("Логин")
	labelPassword := widget.NewLabel("Пароль")
	labelError := widget.NewLabel("")

	entryLogin := widget.NewEntry()
	entryPassword := widget.NewPasswordEntry()

	buttonAuth := widget.NewButton("Войти", func() {
		if entryLogin.Text != "" && entryPassword.Text != "" {
			isAuth := isAuth(entryLogin.Text, entryPassword.Text)
			if isAuth {
				openWindowMain()
			} else {
				labelError.SetText("Неправильный логин или пароль")
			}
		} else {
			labelError.SetText("Обнаружены пустые поля")
		}
	})

	content := container.NewVBox(
		labelLogin,
		entryLogin,
		labelPassword,
		entryPassword,
		container.NewHBox(layout.NewSpacer(), labelError, layout.NewSpacer()),
		container.NewHBox(layout.NewSpacer(), buttonAuth, layout.NewSpacer()),
	)

	return content
}

func isAuth(login string, password string) bool {
	if login == "admin" && password == "12345" {
		userRole = "admin"

		return true
	} else if login == "user" && password == "1111" {
		userRole = "common"

		return true
	}

	return false
}
