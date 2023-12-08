package components

import (
	"log"

	"github.com/gdamore/tcell/v2"
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
type App struct {
	A *tview.Application // Application
	N *Navbar            // Navbar
	S *Sidebar           // Sidebar
	G *tview.Grid        // Root grid
	L *tview.Grid        // View grid

	ShowNavbar  bool // toggle navbar
	ShowSidebar bool // toggle sidebar

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
func NewApp(path string) *App {
	grids, names := GridsFromYaml(path)
	layout := App{
		A:          tview.NewApplication(),
		G:          tview.NewGrid(),
		N:          NewNavbar(names),
		S:          NewSidebar(names),
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

func (l *App) SetKeymaps() {
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

		case tcell.KeyCtrlN:
			// Toggle navbar
			l.ShowNavbar = !l.ShowNavbar
			l.Render()
			return nil

		case tcell.KeyCtrlS:
			// Toggle navbar
			l.ShowSidebar = !l.ShowSidebar
			l.Render()
			return nil

		case tcell.KeyEnter, tcell.KeyTab:
			// Go to next view
			l.N.Current = (l.N.Current + 1) % len(l.N.Labels)
			l.S.Current = (l.S.Current + 1) % len(l.S.Labels)
			l.Name = l.N.Labels[l.N.Current]
			l.Render()
			return nil
		}
		return event
	})
}

func (l *App) Render() {
	// Complete reset, hopefully not expensive op
	// If forms, will lose state?, so probably better to use tview pages
	// Remove all elements
	l.G.Clear()

	// Add content and navbar if needed
	if l.ShowSidebar {
		if l.ShowNavbar {
			// sidebar, navbar, content
			l.G.SetRows(1, -1).SetColumns(-1, -4)
			l.G.AddItem(l.S, 0, 0, 2, 1, 0, 0, false)
			l.G.AddItem(l.N, 0, 1, 1, 1, 0, 0, false)
			l.G.AddItem(l.Grids[l.Name], 1, 1, 1, 1, 0, 0, false)
		} else {
			// sidebar, content
			l.G.SetRows(-1).SetColumns(-1, -4)
			l.G.AddItem(l.S, 0, 0, 1, 1, 0, 0, false)
			l.G.AddItem(l.Grids[l.Name], 0, 1, 1, 1, 0, 0, false)
		}
	} else {
		if l.ShowNavbar {
			// navbar, content
			l.G.SetRows(1, -1).SetColumns(-1)
			l.G.AddItem(l.N, 0, 0, 1, 1, 0, 0, false)
			l.G.AddItem(l.Grids[l.Name], 1, 0, 1, 1, 0, 0, false)
		} else {
			// only content
			l.G.SetRows(-1).SetColumns(-1)
			l.G.AddItem(l.Grids[l.Name], 0, 0, 1, 1, 0, 0, false)
		}
	}

	// Reflect changes to app
	l.A.SetRoot(l.G, true).SetFocus(l.G)
}

func RunApp(path string) {
	// Get path from args
	// path := "layouts/two-column.yaml"
	layout := NewApp(path)

	// Run the application
	err := layout.A.Run()
	if err != nil {
		log.Fatal(err)
	}
}
