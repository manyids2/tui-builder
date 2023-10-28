package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Initialize application
	app := tview.NewApplication()

	labelA := tview.NewBox().SetBorder(true).SetTitle("labelA")
	labelB := tview.NewBox().SetBorder(true).SetTitle("labelB")

	gridA := tview.NewGrid()
	gridA.SetRows(-1).SetColumns(-1, -1)

	gridA.AddItem(labelA, 0, 0, 1, 1, 0, 0, false)
	gridA.AddItem(labelB, 0, 1, 1, 1, 0, 0, false)

	gridB := tview.NewGrid()
	gridB.SetRows(-1).SetColumns(-1)
	gridB.AddItem(labelB, 0, 0, 1, 1, 0, 0, false)

	focused := "G"
	app.SetRoot(gridA, true).SetFocus(gridA)

	// Capture user input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			switch focused {
			case "A":
				focused = "B"
				app.SetRoot(gridA, true)
				app.SetFocus(labelB)
			case "B":
				focused = "G"
				app.SetRoot(gridB, true)
				app.SetFocus(labelB)
			case "G":
				focused = "A"
				app.SetRoot(gridA, true)
				app.SetFocus(labelA)
			}
			return nil
		case tcell.KeyEsc, tcell.KeyEnter:
			app.Stop()
			return nil
		}
		return event
	})

	// Run the application
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
