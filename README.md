# Note 
Commandline note taking tool in Go



### Installation
Make sure you have your GOPATH set properly with GOPATH/bin in your PATH

```bash
$ go install github.com/ironbay/note
```
    
Include the NOTEHOME environment variable in your .bash_profile

```bash
$ export NOTEHOME=~/docs/notes
```

### Add note
Use @ symbol to tag notes. Tags at the beginning will note be included in the note

```bash
$ note @todo @ironbay rewrite note taking tool in @go 
```

The message "rewrite note taking tool in go" will be written to **todo.txt** **ironbay.txt** and **go.txt**


### List notes

```bash
$ note #lists all notes
$ note @todo @idea #list todo.txt and idea.txt
```


### Remove note //TODO
```bash
$ note rm any phrase from note
```