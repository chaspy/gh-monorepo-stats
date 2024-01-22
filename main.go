package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	err := run()
	if err != nil {
		log.Fatal(err) //nolint:forbidigo
	}
}

func run() error {
    dirs, err := ioutil.ReadDir(".")
    if err != nil {
        fmt.Println(err)
        return nil
    }

    for _, d := range dirs {
        if d.IsDir() {
            dirName := d.Name()
            language, packageFile := detectLanguage(dirName)
            if language != "" {
                loc := countLinesOfCode(dirName, language)
                fmt.Printf("%s, %s, %s, %d\n", dirName, packageFile, language, loc)
            } else {
                fmt.Println(dirName)
            }
        }
    }

	return nil
}

func detectLanguage(dirName string) (string, string) {
    files, _ := ioutil.ReadDir(dirName)
    for _, f := range files {
        switch f.Name() {
        case "Gemfile":
            return "Ruby", "Gemfile"
        case "yarn.lock", "pnpm-lock.yaml":
            if f.Name() == "pnpm-lock.yaml" {
                return "TypeScript", "pnpm-lock.yaml"
            }
            return "TypeScript", "yarn.lock"
        case "go.mod":
            return "Go", "go.mod"
        case "mix.exs":
            return "Elixer", "mix.exs"
        case "Cargo.toml":
            return "Rust", "Cargo.toml"
        case "requirements.txt", "setup.py", "poetry.lock", "Pipfile.lock":
            return "Python", f.Name()
        }
    }
    return "", ""
}

func countLinesOfCode(dirName, language string) int {
    var extension string
    switch language {
    case "Ruby":
        extension = ".rb"
    case "TypeScript":
        extension = ".ts"
    case "Go":
        extension = ".go"
    case "Elixer":
        extension = ".ex"
    case "Rust":
        extension = ".rs"
    case "Python":
        extension = ".py"
    }

    var loc int
    filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
        if strings.HasSuffix(path, extension) {
            file, err := os.Open(path)
            if err != nil {
                return err
            }
            defer file.Close()

            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                loc++
            }
            return scanner.Err()
        }
        return nil
    })

    return loc
}
