package getopt

import (
	"bytes"
	"flag"
	"strings"
	"testing"
)

func TestVisitAllInsertionOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(false) // Use insertion order

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	var names []string
	fs.VisitAll(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := []string{"zebra", "alpha", "beta"}
	got := strings.Join(names, ",")
	wantStr := strings.Join(want, ",")

	if got != wantStr {
		t.Errorf("VisitAll insertion order: got %q, want %q", got, wantStr)
	}
}

func TestVisitAllLexicographicalOrder(t *testing.T) {
	// Use lexicographical order
	fs := NewFlagSet("test", flag.ContinueOnError)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	var names []string
	fs.VisitAll(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := []string{"alpha", "beta", "zebra"}
	got := strings.Join(names, ",")
	wantStr := strings.Join(want, ",")

	if got != wantStr {
		t.Errorf("VisitAll lexicographical order: got %q, want %q", got, wantStr)
	}
}

func TestVisitInsertionOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(false) // Use insertion order

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	// Parse to actually set the flags
	fs.Parse([]string{"--zebra", "--alpha"})

	var names []string
	fs.Visit(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := []string{"zebra", "alpha"}
	got := strings.Join(names, ",")
	wantStr := strings.Join(want, ",")

	if got != wantStr {
		t.Errorf("Visit insertion order: got %q, want %q", got, wantStr)
	}
}

func TestVisitLexicographicalOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(true) // Use lexicographical order

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	// Parse to actually set the flags
	fs.Parse([]string{"--zebra", "--alpha"})

	var names []string
	fs.Visit(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := []string{"alpha", "zebra"}
	got := strings.Join(names, ",")
	wantStr := strings.Join(want, ",")

	if got != wantStr {
		t.Errorf("Visit lexicographical order: got %q, want %q", got, wantStr)
	}
}

func TestSortFlagsToggle(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")

	// Start with insertion order (default is false)
	fs.SortFlags(false)
	var names1 []string
	fs.VisitAll(func(f *flag.Flag) {
		names1 = append(names1, f.Name)
	})

	// Switch to lexicographical
	fs.SortFlags(true)
	var names2 []string
	fs.VisitAll(func(f *flag.Flag) {
		names2 = append(names2, f.Name)
	})

	want1 := "zebra,alpha"
	want2 := "alpha,zebra"
	got1 := strings.Join(names1, ",")
	got2 := strings.Join(names2, ",")

	if got1 != want1 {
		t.Errorf("Before toggle: got %q, want %q", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("After toggle: got %q, want %q", got2, want2)
	}
}

func TestPrintDefaultsInsertionOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(false) // Use insertion order

	var buf bytes.Buffer
	fs.SetOutput(&buf)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	fs.PrintDefaults()
	out := buf.String()

	// Check that zebra appears before alpha
	zebraPos := strings.Index(out, "zebra")
	alphaPos := strings.Index(out, "alpha")

	if zebraPos == -1 || alphaPos == -1 {
		t.Errorf("Expected to find both zebra and alpha in output")
	}
	if zebraPos > alphaPos {
		t.Errorf("Expected zebra to appear before alpha in insertion order, got:\n%s", out)
	}
}

func TestPrintDefaultsLexicographicalOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	// fs.SortFlags(true) // Use lexicographical order (default)

	var buf bytes.Buffer
	fs.SetOutput(&buf)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	fs.PrintDefaults()
	out := buf.String()

	// Check that alpha appears before zebra
	zebraPos := strings.Index(out, "zebra")
	alphaPos := strings.Index(out, "alpha")

	if zebraPos == -1 || alphaPos == -1 {
		t.Errorf("Expected to find both zebra and alpha in output")
	}
	if alphaPos > zebraPos {
		t.Errorf("Expected alpha to appear before zebra in lexicographical order, got:\n%s", out)
	}
}
