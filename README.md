# cmdhandle
## Command line parser modelled on http.Mux style

A command line processor to map space separated string patterns to functions.  

---

### Example  
`import cmdline "github.com/eurozulu/cmdhandle"`  
  
`cmdline.Handle("admin", doAdmin)`  
`cmdline.Handle("admin add", doAdminAdd)`  
`cmdline.Handle("admin remove", doAdminRemove)`  
`cmdline.Handle("admin remove all", doAdminRemoveAll)`  
  
When command line is `admin remove thisone -f`  
Will map to the `doAdminRemove` function with 'thisone' as the only argument and 'f' as a flag, with no value.
  
Maps four commands, all begining with 'admin'  
The most specific, longest mapping takes precedence.  
  
```
func doAdminRemove(cmd cmdline.CommandLine) error {
    // cmd.Args contains the non flag args following the "admin remove"
    if len(cmd.Args()) < 1 {
        return fmt.Errorf("must supply 'name' to remove")
    }
    name := cmd.Args(0)

    // Flags contains the flags parsed in the command line
    _, ok := cmd.Flags().Get("-force", "f")
    if !ok {
        // --force flag not present
        if !confirmRemove(name) {
            return fmt.Errorf("aborted removing %s", name)
        }
    }
   
    ... Remove "name" item ...
}

func doAdminRemoveAll(cmd cmdline.CommandLine) error {
    _, ok := cmd.Flags().Get("-force", "f")
    if !ok {
        // --force flag not present
        if !confirmRemoveAll() {
            return fmt.Errorf("aborted removing all")
        }
    }
    
    ... remove all items...
}

```


---

### Patterns
Patterns are the string patterns which match with the begining of the given command line.  
Each pattern has a 'depth', the number of space separated words in the pattern'.
The higher the depth, the more specific the match of the pattern.

  
### CommandFunc  
Each pattern maps to a `CommandFunc`, which has a single argument: `CommandLine`  
This contains the Flags and non flag argument following the matched pattern


### CommandLine
`CommandLine` is the parsed command line following the matched pattern.  
When a command is matched to the command line, the matching pattern is removed and
any remaining arguments placed in the `CommandLine`  
And remaining arguments starting with a '-' are parsed as value flags.  
e.g.  
given a command line of `mycommand one two three -f1 first -f2 --flag3 last`  
Matched to a pattern of `mycommand one`.  
A `CommandLine` would contain:  
Args : "two", "three"
Flags: "f1"="first, "f2"=""", "-flag3"="last  


