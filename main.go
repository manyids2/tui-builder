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
	A *tview.Application // Application
	N *components.Navbar // Navbar
	G *tview.Grid        // Root grid
	L *tview.Grid        // View grid

	ShowNavbar bool // toggle navbar

	Path  string                 // path to yaml
	Grids map[string]*tview.Grid // all views from yaml
	Names []string               // names of views for convinience
	Name  string                 // current
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

// Parent grid for our app with navbar
func NewGridLayout(path string) *GridLayout {
	grids, names := GridsFromYaml(path)
	layout := GridLayout{
		A:          tview.NewApplication(),
		G:          tview.NewGrid(),
		N:          components.NewNavbar(names),
		ShowNavbar: true,
		Path:       path,
		Grids:      grids,
		Names:      names,
		Name:       names[0],
	}
	layout.G.SetRows(1, -1).SetColumns(-1)
	layout.SetKeymaps()
	layout.Name = names[layout.N.Current]
	layout.Render()
	return &layout
}

func (l *GridLayout) SetKeymaps() {
	// Capture user input
	l.A.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			l.A.Stop()
			return nil
		case tcell.KeyCtrlR:
			// Refresh grids and names
			grids, names := GridsFromYaml(l.Path)
			l.Grids = grids
			l.Names = names
			// TODO: Check that name is in names, hasnt changed index, etc.
			l.Render()
			return nil

		case tcell.KeyCtrlSpace:
			// Toggle navbar
			l.ShowNavbar = !l.ShowNavbar
			l.Render()
			return nil
		case tcell.KeyEnter, tcell.KeyTab:
			// Go to next view
			l.N.Current = (l.N.Current + 1) % len(l.N.Labels)
			l.Name = l.N.Labels[l.N.Current]
			l.Render()
			return nil
		}
		return event
	})
}

func (l *GridLayout) Render() {
	// Complete reset, hopefully not expensive op
	// If forms, will lose state?, so probably better to use tview pages
	// Remove all elements
	l.G.Clear()

	// Add content and navbar if needed
	if l.ShowNavbar {
		l.G.AddItem(l.N, 0, 0, 1, 1, 0, 0, false)
		l.G.AddItem(l.Grids[l.Name], 1, 0, 1, 1, 0, 0, false)
	} else {
		l.G.AddItem(l.Grids[l.Name], 0, 0, 2, 1, 0, 0, false)
	}

	// Reflect changes to app
	l.A.SetRoot(l.G, true).SetFocus(l.G)
}

func main() {
	// Get path from args
	path := "layouts/two-column.yaml"
	layout := NewGridLayout(path)

	// Run the application
	err := layout.A.Run()
	if err != nil {
		log.Fatal(err)
	}
}
