# Debugging in Go With Go 

Personal notes for presentation

## Delve

with the command line

```
dlv debug
(dlv) b example.go:5 
(dlv) print variableName
Command failed: could not find symbol value for success
(dlv) continue
(dlv) n
(dlv) print event
(dlv) n
(dlv) print event
"Go meetupt Frankfurt"
(dlv) exit
```

## Godebug

```
godebug run example.go
(godebug) l 
prints the code 
(godebug) p success
true
(godebug) q
```

