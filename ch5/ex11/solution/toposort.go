package toposort

import (
	"fmt"
	"sort"
)

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	reportCycles(m)

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func reportCycles(m map[string][]string) {
	var cycles [][2]string

	for course, prereqs := range m {
		for _, p := range prereqs {
			if prereqs2, ok := m[p]; ok && hasPrerequisite(prereqs2, course) {
				// Sort cycle alphabetically so calls to hasCycle rely on a deterministic order
				cslice := []string{course, p}
				sort.Strings(cslice)

				var carray [2]string
				copy(carray[:], cslice)

				if !hasCycle(cycles, carray) {
					fmt.Printf("cycle detected in lessons %s and %s\n", course, p)
					cycles = append(cycles, carray)
				}
			}
		}
	}

	if cycles != nil {
		fmt.Println()
	}
}

func hasCycle(cycles [][2]string, lookup [2]string) bool {
	for _, c := range cycles {
		if c == lookup {
			return true
		}
	}

	return false
}

func hasPrerequisite(prereqs []string, lookup string) bool {
	for _, p := range prereqs {
		if p == lookup {
			return true
		}
	}

	return false
}
