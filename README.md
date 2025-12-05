# mattmc3/getopt

This fork of github.com/rsc/getopt adds the option to disable the default lexicographical flag sorting
so that the developer can maintain better control of help output.

To preserve insertion order, call `SortFlags(false)` on your `FlagSet`.

```go
fs := getopt.NewFlagSet("example", flag.ExitOnError)
fs.SortFlags(false)  // Maintain insertion order
// ... define flags in desired order ...
fs.PrintDefaults()
```

Sorting affects `Visit`, `VisitAll`, and `PrintDefaults` output.

**Implementation:** Uses code generation (`go generate`) to wrap standard library flag methods while tracking insertion order internally. The getopt `Parse` method was modified to call `FlagSet.Set()` instead of `Value.Set()` directly, ensuring parsed flags are tracked correctly.

## Original docs

[For full package documentation, see [https://godoc.org/rsc.io/getopt](https://godoc.org/rsc.io/getopt).]

    package getopt // import "mattmc3/getopt"

Package getopt parses command lines using [_getopt_(3)](http://man7.org/linux/man-pages/man3/getopt.3.html) syntax. It is a
replacement for `flag.Parse` but still expects flags themselves to be defined
in package flag.

Flags defined with one-letter names are available as short flags (invoked
using one dash, as in `-x`) and all flags are available as long flags (invoked
using two dashes, as in `--x` or `--xylophone`).

To use, define flags as usual with [package flag](https://godoc.org/flag). Then introduce any aliases
by calling `getopt.Alias`:

    getopt.Alias("v", "verbose")

Or call `getopt.Aliases` to define a list of aliases:

    getopt.Aliases(
    	"v", "verbose",
    	"x", "xylophone",
    )

One name in each pair must already be defined in package flag (so either
"v" or "verbose", and also either "x" or "xylophone").

Then parse the command-line:

    getopt.Parse()

If it encounters an error, `Parse` calls `flag.Usage` and then exits the
program.

When writing a custom `flag.Usage` function, call `getopt.PrintDefaults` instead
of `flag.PrintDefaults` to get a usage message that includes the
names of aliases in flag descriptions.

At initialization time, package getopt installs a new `flag.Usage` that is the same
as the default `flag.Usage` except that it calls `getopt.PrintDefaults` instead
of `flag.PrintDefaults`.

This package also defines a `FlagSet` wrapping the standard `flag.FlagSet`.
