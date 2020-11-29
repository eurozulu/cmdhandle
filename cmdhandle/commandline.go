package cmdhandle

import (
	"strings"
)

// CommandLine is a parsed representation of a 'raw' command line argument list.
// The arguments are diveded into 'args' and 'flags'. Flags re args starting with '-',
// args are the non flags preceeding the first flag.
type CommandLine interface {
	// Flags contains all the given flags (arguments begining with"-"
	Flags() Flags
	// Args gets all the non flag arguments form the command line
	Args() []string
	// Arg is a helper to get an unnamed argument from a given index.
	// if the given index is <0 or > len Args, empty string is returned.
	Arg(index int) string
}

// Flags are a set of named values.  Keys are the 'flag' name. The value is the following, non flag arguments, joined with a space.
type Flags map[string]string

// Get the named flag value.
// name is the key following the '-'.  short is an alternative name for the same flag.
// e.g. v, ok := Get("verbose", "v")
// Returns the flag value if it exists. bool indicates if flag is present.
func (f Flags) Get(name, short string) (string, bool) {
	n, ok := f[name]
	if !ok && short != "" {
		n, ok = f[short]
	}
	return n, ok
}

type commandLine struct {
	args  []string
	flags Flags
}

// Args gets the non flag arguments from the command line
func (c commandLine) Args() []string {
	return c.args
}

func (c commandLine) Arg(index int) string {
	if index < 0 || index >= len(c.args) {
		return ""
	}
	return c.args[index]
}

// Flags gets the flags from the command line
func (c commandLine) Flags() Flags {
	return c.flags
}

// ParseCommandLine parses the given arguments into 'flags' and non flags.
// A flag is an argument starting with a dash '-'. All arguments following a flag, which do not start with a dash,
// are joined with spaces to become the flag value.
// Arguments not starting with a dash and preceding the first flag are the CommandLine arguments.
// Note: Flags starting with double dash retain one of the dashes in the flag key.
func ParseCommandLine(arguments ...string) *commandLine {
	flags := make(Flags)
	var args []string

	var k string
	for _, arg := range arguments {
		if strings.HasPrefix(arg, "-") {
			k = arg[1:]
			flags[k] = ""
			continue
		}
		if k == "" {
			args = append(args, arg)
		} else {
			flags[k] = strings.Join([]string{flags[k], arg}, " ")
		}
	}
	return &commandLine{
		args:  args,
		flags: flags,
	}
}
