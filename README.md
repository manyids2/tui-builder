# tui-builder

Simple builder of TUI from bash commands, some simple widgets from `tview`, `tcell` and channel concept from `bubbletea`.

Roadmap:

1. Basic app shell
2. Working example
3. Code generation / framework

## Structures

To contain geometry information and ref to object.

```go
type GridItem struct {
	P                           *tview.Primitive
	Row, Column                 int
	RowSpan, ColumnSpan         int
	minGridHeight, minGridWidth int
	focus                       bool
}
```

To maintain refs to all relevant objects and state of application.

```go
// Container grid
type GridLayout struct {
	G     *tview.Grid
	Items map[string]GridItem
	focus bool
}
```
