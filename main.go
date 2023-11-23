package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/manyids2/tui-builder/components"
	"github.com/manyids2/tui-builder/layouts"
	"github.com/rivo/tview"
)

// Item with geometry, focus and ref to primitive
type GridItem struct {
	P                           *tview.Box
	Row, Column                 int
	RowSpan, ColumnSpan         int
	MinGridHeight, MinGridWidth int
	Focus                       bool
}

// Container grid
type GridLayout struct {
	G      *tview.Grid
	N      *components.Navbar
	Layout *tview.Grid
}

func GridsFromYaml(path string) (map[string]*tview.Grid, []string) {
	// Read layout config
	var names []string
	config := layouts.NewYamlLayoutOuter(path)
	for _, v := range config.YamlLayout {
		names = append(names, v.Name)
	}
	if len(names) == 0 {
		log.Fatalln("No layouts found: ", path)
	}

	// Initialize grid from yaml config
	grids := make(map[string]*tview.Grid)
	for _, view := range config.YamlLayout {
		grid := tview.NewGrid()
		grid.SetRows(view.RowSpans...).SetColumns(view.ColumnSpans...)
		for k, v := range view.Items {
			box := tview.NewBox().SetBorder(true).SetTitle(k)
			grid.AddItem(box, v[0], v[1], v[2], v[3], view.MinGridWidth, view.MinGridHeight, false)
		}
		grids[view.Name] = grid
	}

	return grids, names
}

func main() {
	// Get path from args
	path := "layouts/two-column.yaml"
	grids, names := GridsFromYaml(path)

	// Define grid for our app
	layout := GridLayout{
		G: tview.NewGrid(),
		N: components.NewNavbar([]string{}),
	}
	layout.G.SetRows(1, -1).SetColumns(-1)
	layout.G.AddItem(layout.N, 0, 0, 1, 1, 0, 0, false)
	layout.N.Labels = names

	// Display first view
	name := layout.N.Labels[layout.N.Current]
	layout.G.AddItem(grids[name], 1, 0, 1, 1, 0, 0, false)

	// Initialize application
	app := tview.NewApplication()
	app.SetRoot(layout.G, true).SetFocus(layout.G)

	// Capture user input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			app.Stop()
			return nil
		case tcell.KeyEnter:
			// Go to next view
			layout.N.Current = (layout.N.Current + 1) % len(layout.N.Labels)
			name := layout.N.Labels[layout.N.Current]

			// Complete reset, hopefully not expensive op
			// If forms, will lose state?, so probably better to use tview pages
			layout.G.Clear()
			layout.G.AddItem(layout.N, 0, 0, 1, 1, 0, 0, false)
			layout.G.AddItem(grids[name], 1, 0, 1, 1, 0, 0, false)
			app.SetRoot(layout.G, true).SetFocus(layout.G)

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
