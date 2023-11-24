# Grids

`tview` grids:

```go
// Basic 2 column grid
//
// +---+---+
// |   |   |
// +---+---+
//
gridA := tview.NewGrid()

// SetColumns defines how the columns of the grid are distributed. Each value
// defines the size of one column, starting with the leftmost column. Values
// greater than 0 represent absolute column widths (gaps not included). Values
// less than or equal to 0 represent proportional column widths or fractions of
// the remaining free space, where 0 is treated the same as -1. That is, a
// column with a value of -3 will have three times the width of a column with a
// value of -1 (or 0). The minimum width set with SetMinSize() is always
// observed.
//
// Primitives may extend beyond the columns defined explicitly with this
// function. A value of 0 is assumed for any undefined column. In fact, if you
// never call this function, all columns occupied by primitives will have the
// same width. On the other hand, unoccupied columns defined with this function
// will always take their place.
//
// Assuming a total width of the grid of 100 cells and a minimum width of 0, the
// following call will result in columns with widths of 30, 10, 15, 15, and 30
// cells:
//
//	grid.SetColumns(30, 10, -1, -1, -2)
//
// If a primitive were then placed in the 6th and 7th column, the resulting
// widths would be: 30, 10, 10, 10, 20, 10, and 10 cells.
//
// If you then called SetMinSize() as follows:
//
//	grid.SetMinSize(15, 20)
//
// The resulting widths would be: 30, 15, 15, 15, 20, 15, and 15 cells, a total
// of 125 cells, 25 cells wider than the available grid width.

// one row, 2 equal size columns
gridA.SetRows(-1).SetColumns(-1, -1)

// AddItem(p Primitive, row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool)
// AddItem adds a primitive and its position to the grid. The top-left corner
// of the primitive will be located in the top-left corner of the grid cell at
// the given row and column and will span "rowSpan" rows and "colSpan" columns.
// For example, for a primitive to occupy rows 2, 3, and 4 and columns 5 and 6:
//
//	grid.AddItem(p, 2, 5, 3, 2, 0, 0, true)
//
// If rowSpan or colSpan is 0, the primitive will not be drawn.
//
// You can add the same primitive multiple times with different grid positions.
// The minGridWidth and minGridHeight values will then determine which of those
// positions will be used. This is similar to CSS media queries. These minimum
// values refer to the overall size of the grid. If multiple items for the same
// primitive apply, the one that has at least one highest minimum value will be
// used, or the primitive added last if those values are the same. Example:
//
//	grid.AddItem(p, 0, 0, 0, 0, 0, 0, true). // Hide in small grids.
//	  AddItem(p, 0, 0, 1, 2, 100, 0, true).  // One-column layout for medium grids.
//	  AddItem(p, 1, 1, 3, 2, 300, 0, true)   // Multi-column layout for large grids.
//
// To use the same grid layout for all sizes, simply set minGridWidth and
// minGridHeight to 0.
//
// If the item's focus is set to true, it will receive focus when the grid
// receives focus. If there are multiple items with a true focus flag, the last
// visible one that was added will receive focus.


// First column
gridA.AddItem(labelA, 0, 0, 1, 1, 0, 0, false)

// Second column
gridA.AddItem(labelB, 0, 1, 1, 1, 0, 0, false)
```

