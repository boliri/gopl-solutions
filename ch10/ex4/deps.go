// The program reports the set of packages in a Go workspace that transitively depend on a set
// of packages passed as arguments
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// importPath symbolizes the import path of a Go package
type importPath string

// pkgMetadata represents the metadata of a Go package
//
// includes the package's import path and its transitive dependencies
type pkgMetadata struct {
	ImportPath importPath   `json:"ImportPath"`
	Deps       []importPath `json:"Deps"`
}

// SplitPkgMetadata is a split function for a bufio.Scanner that returns each package's
// metadata from the output of a go list -json <pkgs> command as a token
func SplitPkgMetadata(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// JSONs returned by the go list command are supposed to be wellformed, so it's safe
	// to reduce the problem to just finding parity in curly braces
	closestClosingCurlyAt := bytes.IndexByte(data, '}')
	for !curlyMatch(data[:closestClosingCurlyAt+1]) {
		nextClosingCurlyAt := bytes.IndexByte(data[closestClosingCurlyAt+1:], '}')
		closestClosingCurlyAt += nextClosingCurlyAt + 1
	}

	nextOpeningCurlyAt := bytes.IndexByte(data[closestClosingCurlyAt+1:], '{')
	if nextOpeningCurlyAt == -1 {
		return len(data), data[:closestClosingCurlyAt+1], nil
	}

	return closestClosingCurlyAt + nextOpeningCurlyAt + 1, data[:closestClosingCurlyAt+1], nil
}

// curlyMatch reports whether there's '{' and '}' parity in data or not
func curlyMatch(data []byte) bool {
	return bytes.Count(data, []byte{'{'}) == bytes.Count(data, []byte{'}'})
}

// stringify converts importPaths in a slice to strings
func stringify(list []importPath) []string {
	var paths []string
	for _, p := range list {
		paths = append(paths, string(p))
	}
	return paths
}

func main() {
	targetPkgs := os.Args[1:]

	b, err := exec.Command("go", "list", "-json", "...").Output()
	if err != nil {
		fmt.Println("deps:", err)
		os.Exit(1)
	}

	buf := bytes.NewReader(b)
	s := bufio.NewScanner(buf)
	s.Buffer(make([]byte, 0, len(b)), len(b))
	s.Split(bufio.SplitFunc(SplitPkgMetadata))

	pkgToDeps := make(map[importPath]map[importPath]bool)
	for s.Scan() {
		pmd := pkgMetadata{}
		err := json.Unmarshal(s.Bytes(), &pmd)
		if err != nil {
			fmt.Println("deps: unmarshall:", err)
			return
		}

		if len(pmd.Deps) == 0 {
			// package does not have any transitive deps; skip
			continue
		}

		pkgToDeps[pmd.ImportPath] = make(map[importPath]bool)
		for _, d := range pmd.Deps {
			pkgToDeps[pmd.ImportPath][d] = true
		}
	}

	pkgsFound := []importPath{}
OUTER:
	for pkg, deps := range pkgToDeps {
		for _, tp := range targetPkgs {
			if _, ok := deps[importPath(tp)]; !ok {
				continue OUTER
			}
		}
		pkgsFound = append(pkgsFound, pkg)
	}

	if len(pkgsFound) == 0 {
		fmt.Printf("no packages transitively depending on %s found\n", strings.Join(targetPkgs, ", "))
		os.Exit(0)
	}

	pkgsFoundStr := stringify(pkgsFound)
	sort.Strings(pkgsFoundStr)
	fmt.Println(strings.Join(pkgsFoundStr, "\n"))
}
