package cmdhandle

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type CommandFunc func(cmd CommandLine) error

// DefaultHandler is the default CommandHandler which uses the 'os.Args' list.
var DefaultHandler CommandHandler

// Handle maps the given pattern to the given handler in the DefaultHandler
func Handle(pattern string, handler CommandFunc) {
	DefaultHandler.Handle(pattern, handler)
}

// Serve the 'os.Args' using the DefaultHandler, processes the os command line.
func Serve() error {
	return DefaultHandler.Serve(os.Args[1:]...)
}

// command is a predefined mapping of a command pattern, to a handler function
type command struct {
	pattern []string
	handler CommandFunc
}

func newCommand(pattern string, handler CommandFunc) *command {
	args := strings.Split(pattern, " ")
	return &command{pattern: args, handler: handler}
}

// Depth indicates how 'deep' in the command line the pattern is.
// a single word command is one, or three word command is three.
// e.g. "this is a command" has a depth of four.
func (c command) Depth() int {
	return len(c.pattern)
}

func (c command) Match(args ...string) bool {
	if len(args) < c.Depth() {
		return false
	}
	for i, arg := range c.pattern {
		if arg != args[i] {
			return false
		}
	}
	return true
}

type CommandHandler struct {
	patterns []*command
}

// Handle maps the given pattern to the given CommandFunc
func (ch *CommandHandler) Handle(pattern string, handler CommandFunc) {
	ch.patterns = append(ch.patterns, newCommand(pattern, handler))
}

// Serve matches the most specific pattern to the given arguments and calls the mapped function.
func (ch CommandHandler) Serve(args ...string) error {
	cmdLine := ParseCommandLine(args...)
	var ptns []*command
	for _, p := range ch.patterns {
		if p.Match(cmdLine.Args()...) {
			ptns = append(ptns, p)
		}
	}
	if len(ptns) == 0 {
		return fmt.Errorf("unknown command")
	}

	// Sort so first is deepest
	sort.Slice(ptns, func(i, j int) bool {
		return ptns[i].Depth() > ptns[j].Depth()
	})

	cmdLine.args = cmdLine.args[ptns[0].Depth():] // Trim off pattern matched from args.
	return ptns[0].handler(cmdLine)
}
