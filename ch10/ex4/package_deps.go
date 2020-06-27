// Construct a tool that reports the set of all packages in the workspace that transitively depend
// on the packages specified by the arguments. Hint: you will need to run go list twice, once for
// the initial packages and once for all packages. You may want to parse its JSON output using
// the encoding/json package (§4.5).
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

type packageInfo struct {
	Name, ImportPath string
	Deps             []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <path/to/package/subtree>", os.Args[0])
		os.Exit(1)
	}

	targetInfo, err := listPackageInfo(os.Args[1:]...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	workspaceInfo, err := listPackageInfo("...")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	deps := intersection(targetInfo, workspaceInfo)
	for _, dep := range deps {
		fmt.Println(dep.ImportPath)
	}
}

func listPackageInfo(targets ...string) ([]packageInfo, error) {
	args := append([]string{"list", "-json"}, targets...)
	cmd := exec.Command("go", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var packages []packageInfo
	decoder := json.NewDecoder(bytes.NewReader(out))
	for {
		var pkg packageInfo
		if err := decoder.Decode(&pkg); err != nil {
			if err == io.EOF {
				return packages, nil
			}
			return nil, err
		}
		packages = append(packages, pkg)
	}
}

// intersection returns a slice of all packages in the workspace that transitively
// depend on the packages supplied as arguments to the program.
func intersection(targetInfo, workspaceInfo []packageInfo) []packageInfo {
	targetImports := make(map[string]bool)
	for _, pkg := range targetInfo {
		targetImports[pkg.ImportPath] = true
	}

	matches := make(chan packageInfo)
	var wg sync.WaitGroup
	for _, pkg := range workspaceInfo {
		wg.Add(1)
		go func(pkg packageInfo) {
			for _, dep := range pkg.Deps {
				if targetImports[dep] {
					matches <- pkg
					break
				}
			}
			wg.Done()
		}(pkg)
	}

	go func() {
		wg.Wait()
		close(matches)
	}()

	results := make([]packageInfo, 0)
	for pkg := range matches {
		results = append(results, pkg)
	}
	return results
}
