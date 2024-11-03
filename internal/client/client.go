package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/jackc/pgx/v5"
)

var (
	windowMain fyne.Window
	AppFyne    fyne.App

	connection *pgx.Conn
)

func Start(conn *pgx.Conn) {
	connection = conn

	AppFyne = app.New()
	windowMain = AppFyne.NewWindow("")

	openAuthWindow()

	windowMain.Show()
	windowMain.SetMaster()
	AppFyne.Run()
}
