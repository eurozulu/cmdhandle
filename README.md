# cmdhandle
## Command line parser modelled on http.Mux style

A command line processor to map space separated string patterns to functions.  

---

### Patterns
Patterns are the string patterns which match with the begining of the given command line.  
Each pattern has a 'depth', the number of space separated words in the pattern'.
The higher the depth, the more specific the match of the pattern.

  
### CommandFunc  
Each pattern maps to a `CommandFunc`, which has a single argument: `CommandLine`  
This contains the Flags and non flag argument following the matched pattern


### Example  
`cmdline.Handle("admin", doAdmin)`  
`cmdline.Handle("admin add", doAdminAdd)`  
`cmdline.Handle("admin remove", doAdminRemove)`  
`cmdline.Handle("admin remove all", doAdminRemoveAll)`  
  
Maps four commands, all begining with 'admin'  
The most specific, longest mapping takes precedence.  
  
```
func doAdminAdd(cmd cmdline.CommandLine) error {
    if len(cmd.Args()) < 1 {
        return fmt.Errorf("must supply 'name'")
    }
    // cmd.Args contains the non flag args following the "admin add"
    
    mf, ok := cmd.Flags().Get("myflag", "m")
    if ok {
        // the -myflag has been specified.
    }
}
```
  
When command line is `admin remove thisone -f thatfile`  
Will map to the `doAdminRemove` function with 'thisone' as an argument and 'f' as a flag.

