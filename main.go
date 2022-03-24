package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"lncount/utils"
	"lncount/worker"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	// DEFAULT: values below are default values
	fpath          = "./"
	Suffix         = "go"
	IgnoredFolders = []string{}
	DebugMode      = false
)

func main() {
	// setup input parse
	args := os.Args[1:]
	for i := range args {
		if len(args[i]) > 1 && args[i][0] == '-' {
			if strings.ToLower(args[i][1:]) == "help" {
				fmt.Println("LNCounter Help: ")
				fmt.Println("     -path=<Path>")
				fmt.Println("     -suffix=<File Ext>")
				fmt.Println("     -ignore=<folders joined with `,`>")
				fmt.Println("           Example: ")
				fmt.Println("                -ignore=vendor,test,worker")
				fmt.Println("     -debug=<true/false>")
				fmt.Println("\r\nPlease note, every arguement is not case sensitve")
				return
			}
		}

		kv := strings.Split(args[i][1:], "=")
		switch strings.ToLower(kv[0]) {
		case "path":
			fpath = kv[1]
		case "suffix":
			Suffix = kv[1]
		case "debug":
			opt := strings.ToLower(kv[1])
			if opt == "true" {
				DebugMode = true
			} else if opt == "false" {
				DebugMode = false
			}
		case "ignore":
			folders := strings.Split(kv[1], ",")
			for _, v := range folders {
				IgnoredFolders = append(IgnoredFolders, v)
			}

		}
	}
	go BackgroundWorker()
	if !DebugMode {
		for {
			log.Println("[+]==================[+]")
			log.Println("  Selfrep | Linecounter")
			log.Println("  Path: ", fpath)
			log.Println("  Suffix: ", Suffix)
			log.Println("  Lines: ", worker.GlobalLines)
			log.Println("  Ignored Folders: ", strings.Join(IgnoredFolders, ","))
			log.Println("  Debug: ", DebugMode)
			log.Println("[+]==================[+]")
			utils.Sleep(30)
			utils.Clear()
		}
	} else {
		for {
			log.Println("Lines: ", worker.GlobalLines)
			utils.Sleep(1000)
		}
	}

}

func BackgroundWorker() {
	Queue := worker.StartQueue()

	filepath.Walk(fpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err.Error())
		}
		if info.IsDir() {
			if found := InArray(info.Name(), IgnoredFolders); found {
				return filepath.SkipDir
			}
		}
		// check suffix
		if strings.HasSuffix(info.Name(), fmt.Sprintf(".%s", Suffix)) && !info.IsDir() {
			utils.Sleep(150)
			// Open file and read how many lines
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err.Error())
			}
			if DebugMode {
				log.Println("Opening FD: ", f.Fd())

			}
			defer f.Close()
			// read file
			scanner := bufio.NewScanner(f)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				// This may need configuring depending on what language you are using, the `//` `/*` `#` are all comment prefixes but may be different in other languages
				if len(scanner.Text()) > 0 && !strings.HasPrefix(scanner.Text(), "//") && !strings.HasPrefix(scanner.Text(), "/*") && !strings.HasPrefix(scanner.Text(), "#") {
					Queue.Iterate()
				}
			}
		}

		return nil
	})
}

func InArray(str string, arr []string) bool {
	for _, v := range arr {
		if strings.EqualFold(str, v) {
			return true
		}
	}
	return false
}
