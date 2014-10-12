package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ActiveState/tail"
	"github.com/fatih/color"
)

var (
	red     = color.New(color.FgRed).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	white   = color.New(color.FgWhite).SprintFunc()
	black   = color.New(color.FgBlack).SprintFunc()
)

func main() {

	defer recover()

	files := flag.String("f", "", "CSV list of files to tail")
	prefix := flag.Bool("p", false, "Prefix output with file name")
	patterns := flag.String("r", "", "List of regexs to search for")
	colours := flag.String("c", "", "List of regex colours")
	delim := flag.String("d", ",", "Regex list delimiter")

	flag.Parse()

	if *files == "" {
		fmt.Println("no file")
		return
	}

	tempList1 := strings.Split(*patterns, *delim)
	tempList2 := strings.Split(*colours, *delim)
	var patternList []string
	var colourList []func(...interface{}) string

	for _, t := range tempList1 {
		if t == "" {
			continue
		}
		patternList = append(patternList, t)
		colourList = append(colourList, red)
	}

	for i, t := range tempList2 {
		if t == "" {
			continue
		}

		if i >= len(colourList) {
			break
		}

		switch t {
		case "r":
			colourList[i] = red
		case "g":
			colourList[i] = green
		case "y":
			colourList[i] = yellow
		case "b":
			colourList[i] = blue
		case "m":
			colourList[i] = magenta
		case "c":
			colourList[i] = cyan
		case "w":
			colourList[i] = white
		case "k":
			colourList[i] = black
		}

	}

	for _, file := range strings.Split(*files, ",") {

		if file == "" {
			continue
		}

		fileList, _ := filepath.Glob(file)

		for _, expandedFile := range fileList {

			go func(file string) {

				var regs []*regexp.Regexp

				if len(patternList) >= 1 {
					for _, pat := range patternList {
						regs = append(regs, regexp.MustCompile(pat))
					}
				}

				location := &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END}

				t, err := tail.TailFile(file, tail.Config{Follow: true, ReOpen: true, Location: location})
				if err != nil {
					fmt.Println(err)
					return
				}

				for line := range t.Lines {

					output := line.Text
					if len(regs) >= 1 {
						for i, reg := range regs {

							process := func(s string) string {
								return colourList[i](s)
							}

							output = reg.ReplaceAllStringFunc(output, process)
						}
					}

					if *prefix {
						fmt.Println(file, output)
					} else {
						fmt.Println(output)
					}
				}
			}(expandedFile)

		}

	}

	select {}

}
