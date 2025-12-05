//go:generate go run ./cmd/gen_flagfuncs/main.go

package getopt

import "flag"

// SortFlags controls whether Visit and VisitAll use insertion order (false)
// or lexicographical order (true). Default is true (lexicographical order).
func (f *FlagSet) SortFlags(sort bool) {
	f.sortFlags = sort
}

// VisitAll visits the flags, calling fn for each.
// It visits all flags, even those not set.
// Order is controlled by SortFlags (lexicographical by default).
func (f *FlagSet) VisitAll(fn func(*flag.Flag)) {
	if f.sortFlags {
		f.FlagSet.VisitAll(fn)
		return
	}
	for _, flag := range f.flags {
		fn(flag)
	}
}

// Visit visits the flags, calling fn for each.
// It visits only those flags that have been set.
// Order is controlled by SortFlags (lexicographical by default).
//
// Note: This only works correctly if flags were set through the underlying
// FlagSet.Set() method. The getopt.Parse() method calls Value.Set() directly
// which bypasses FlagSet's tracking, so Visit may not work as expected after
// getopt-style parsing. Use VisitAll instead if you need to iterate over
// all defined flags.
func (f *FlagSet) Visit(fn func(*flag.Flag)) {
	if f.sortFlags {
		f.FlagSet.Visit(fn)
		return
	}

	// Build set of actual flags that were set
	actual := make(map[string]bool)
	f.FlagSet.Visit(func(flag *flag.Flag) {
		actual[flag.Name] = true
	})

	// Visit in insertion order, but only those that were set
	for _, flag := range f.flags {
		if actual[flag.Name] {
			fn(flag)
		}
	}
}
