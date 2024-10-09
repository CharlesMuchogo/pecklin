package main

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/data/local/database"
	"main.go/data/remote"
	"main.go/pkg/controllers/loader"
	"main.go/pkg/utils"
	"main.go/presentation"
	"os"
)

func main() {
	database.InitializeDatabase()

	m := loader.InitialModel()
	p := tea.NewProgram(m)
	go func() {

		err := remote.FetchPractices()
		if err != nil {
			p.Send(loader.ErrMsg(err))
			return
		}

		p.Send(loader.DataLoadedMsg{})
		utils.ClearScreen()
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	navigation.InitialRoute(func() {
		presentation.MainMenu()
	})
}
