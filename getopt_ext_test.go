package getopt

import (
	"bytes"
	"flag"
	"strings"
	"testing"
)

func TestVisitAllInsertionOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(false)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	var names []string
	fs.VisitAll(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := "zebra,alpha,beta"
	got := strings.Join(names, ",")

	if got != want {
		t.Errorf("VisitAll insertion order: got %q, want %q", got, want)
	}
}

func TestVisitAllLexicographicalOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	var names []string
	fs.VisitAll(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := "alpha,beta,zebra"
	got := strings.Join(names, ",")

	if got != want {
		t.Errorf("VisitAll lexicographical order: got %q, want %q", got, want)
	}
}

func TestVisitInsertionOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(false)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	fs.Parse([]string{"--zebra", "--alpha"})

	var names []string
	fs.Visit(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := "zebra,alpha"
	got := strings.Join(names, ",")

	if got != want {
		t.Errorf("Visit insertion order: got %q, want %q", got, want)
	}
}

func TestVisitLexicographicalOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(true)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	fs.Parse([]string{"--zebra", "--alpha"})

	var names []string
	fs.Visit(func(f *flag.Flag) {
		names = append(names, f.Name)
	})

	want := "alpha,zebra"
	got := strings.Join(names, ",")

	if got != want {
		t.Errorf("Visit lexicographical order: got %q, want %q", got, want)
	}
}

func TestSortFlagsToggle(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")

	fs.SortFlags(false)
	var names1 []string
	fs.VisitAll(func(f *flag.Flag) {
		names1 = append(names1, f.Name)
	})

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
		t.Errorf("Insertion order: got %q, want %q", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("Lexicographical order: got %q, want %q", got2, want2)
	}
}

func TestPrintDefaultsInsertionOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)
	fs.SortFlags(false)

	var buf bytes.Buffer
	fs.SetOutput(&buf)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	fs.PrintDefaults()

	want := `  --zebra
    	z flag
  --alpha
    	a flag
  --beta
    	b flag
`

	if buf.String() != want {
		t.Errorf("PrintDefaults insertion order:\nhave:\n%s\nwant:\n%s", buf.String(), want)
	}
}

func TestPrintDefaultsLexicographicalOrder(t *testing.T) {
	fs := NewFlagSet("test", flag.ContinueOnError)

	var buf bytes.Buffer
	fs.SetOutput(&buf)

	fs.Bool("zebra", false, "z flag")
	fs.Bool("alpha", false, "a flag")
	fs.Bool("beta", false, "b flag")

	fs.PrintDefaults()

	want := `  --alpha
    	a flag
  --beta
    	b flag
  --zebra
    	z flag
`

	if buf.String() != want {
		t.Errorf("PrintDefaults lexicographical order:\nhave:\n%s\nwant:\n%s", buf.String(), want)
	}
}
