package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var root = "~/docs/notes/"

func main() {
	root = os.Getenv("NOTEHOME")
	args := append(os.Args[:0], os.Args[1:]...)
	add(args)

}

func list(tags []string) {
	if len(tags) == 0 {
		files, _ := ioutil.ReadDir(root)
		for _, info := range files {
			tags = append(tags, info.Name()[0:len(info.Name())-4])
		}

	}
	for _, tag := range tags {
		content, _ := ioutil.ReadFile(path.Join(root, tag+".txt"))
		if content == nil {
			continue
		}
		fmt.Println()
		fmt.Println(string(content))
	}
}

func add(args []string) {
	tags := make([]string, 0)
	var buffer bytes.Buffer
	started := false
	for _, arg := range args {
		if arg[0] == '@' {
			arg = arg[1:]
			tags = append(tags, arg)
			if !started {
				continue
			}
		}
		buffer.WriteString(" " + arg)
		started = true
	}
	message := strings.Trim(buffer.String(), " ")
	if message == "" {
		list(tags)
		return
	}
	fmt.Println(tags, message)
	message = "* " + message

	for _, file := range tags {
		full := path.Join(root, file+".txt")
		payload := ""
		if _, err := os.Stat(full); err != nil {
			payload += file + "\n"
		}
		payload += message + "\n"
		f, _ := os.OpenFile(full, os.O_APPEND|os.O_CREATE, 0666)
		f.WriteString(payload)
		f.Close()
	}

}
