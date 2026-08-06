package main

import "testing"

func TestValidChar(t *testing.T) {
    cases := []struct {
        input rune
        want  bool
    }{
        {'.', true},
        {'1', true},
        {'9', true},
        {'0', false},
        {'a', false},
        {'/', false},
    }

    for _, c := range cases {
        got := validChar(c.input)
        if got != c.want {
            t.Errorf("validChar(%q) = %v, want %v", c.input, got, c.want)
        }
    }
}

func TestSolveValidPuzzle(t *testing.T) {
    grid := [9][9]int{
        {0, 9, 6, 2, 4, 0, 0, 0, 1},
        {1, 0, 0, 0, 6, 0, 0, 0, 4},
        {5, 0, 4, 8, 1, 0, 3, 9, 0},
        {0, 0, 7, 9, 5, 0, 0, 4, 3},
        {0, 3, 0, 0, 8, 0, 0, 0, 0},
        {4, 0, 5, 0, 2, 3, 1, 8, 0},
        {0, 1, 0, 6, 3, 0, 0, 5, 9},
        {0, 5, 9, 0, 7, 0, 8, 3, 0},
        {0, 0, 3, 5, 9, 0, 0, 7, 0},
    }

    if !solve(&grid) {
        t.Fatal("expected puzzle to be solvable")
    }

    want := [9][9]int{
        {3, 9, 6, 2, 4, 5, 7, 8, 1},
        {1, 7, 8, 3, 6, 9, 5, 2, 4},
        {5, 2, 4, 8, 1, 7, 3, 9, 6},
        {2, 8, 7, 9, 5, 1, 6, 4, 3},
        {9, 3, 1, 4, 8, 6, 2, 7, 5},
        {4, 6, 5, 7, 2, 3, 9, 1, 8},
        {7, 1, 2, 6, 3, 8, 4, 5, 9},
        {6, 5, 9, 1, 7, 4, 8, 3, 2},
        {8, 4, 3, 5, 9, 2, 1, 6, 7},
    }

    if grid != want {
        t.Errorf("solve() produced incorrect solution\n got: %v\nwant: %v", grid, want)
    }
}

func TestSolveInvalidPuzzle(t *testing.T) {
    grid := [9][9]int{
        {0, 9, 6, 2, 4, 0, 0, 0, 1},
        {1, 0, 0, 0, 6, 0, 1, 0, 4},
        {5, 0, 4, 8, 1, 0, 3, 9, 0},
        {0, 0, 7, 9, 5, 0, 0, 4, 3},
        {0, 3, 0, 0, 8, 0, 0, 0, 0},
        {4, 0, 5, 0, 2, 3, 1, 8, 0},
        {0, 1, 0, 6, 3, 0, 0, 5, 9},
        {0, 5, 9, 0, 7, 0, 8, 3, 0},
        {0, 0, 3, 5, 9, 0, 0, 7, 0},
    }

    if solve(&grid) {
        t.Fatal("expected puzzle to be unsolvable")
    }
}
