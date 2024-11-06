package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/jackc/pgx/v5"
)

var (
	windowMain fyne.Window
	AppFyne    fyne.App

	conn *pgx.Conn
)

func Start(connection *pgx.Conn) {
	conn = connection

	AppFyne = app.New()

	windowMain = AppFyne.NewWindow("")

	openWindowAuth()

	windowMain.Show()
	windowMain.SetMaster()

	AppFyne.Run()
}
