package main

import (
	"bufio"
	"fmt"
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
    ignorePaths := getIgnorePaths()

    dirs, err := os.ReadDir(".")
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to read directory")
    }

    for _, d := range dirs {
        if (d.IsDir() && !strings.HasPrefix(d.Name(), ".")) {
            dirName := d.Name()
            language, packageFile := detectLanguage(dirName)
            if language != "" {
                loc := countLinesOfCode(dirName, language, ignorePaths)
                fmt.Printf("%s, %s, %s, %d\n", dirName, packageFile, language, loc)
            } else {
                fmt.Println(dirName)
            }
        }
    }
	return nil
}

func getIgnorePaths() []string {
    ignorePathsEnv := os.Getenv("IGNORE_PATH")
    if ignorePathsEnv == "" {
        return nil
    }
    return strings.Split(ignorePathsEnv, ",")
}


func detectLanguage(dirName string) (string, string) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return "", ""
	}

    for _, f := range files {
        switch f.Name() {
        case "Gemfile":
            return "Ruby", "Gemfile"
        case "yarn.lock", "pnpm-lock.yaml":
            return "TypeScript", f.Name()
        case "go.mod":
            return "Go", "go.mod"
        case "mix.exs":
            return "Elixir", "mix.exs"
        case "Cargo.toml":
            return "Rust", "Cargo.toml"
        case "requirements.txt", "setup.py", "poetry.lock", "Pipfile.lock":
            return "Python", f.Name()
        }
    }
    return "", ""
}

// nolint:gocyclo
func countLinesOfCode(dirName, language string, ignorePaths []string) int {
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
    err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
        if strings.HasSuffix(path, extension) && !shouldIgnore(path, ignorePaths) {
            file, err := os.Open(path)
            if err != nil {
                return fmt.Errorf("failed to open file: %w", err)
            }
            defer file.Close()

            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                loc++
            }
			if err := scanner.Err(); err != nil {
	            return fmt.Errorf("failed to scan file: %w", err)
			}
        }
        return nil
    })

	if err != nil {
		return 0
	}

    return loc
}

func shouldIgnore(path string, ignorePaths []string) bool {
    for _, ignorePath := range ignorePaths {
        if strings.Contains(path, ignorePath) {
            return true
        }
    }
    return false
}