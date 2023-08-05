// issues looks for GitHub issues given a set of filters
package issues

import (
	"fmt"
	"log"
	"time"

	"gopl.io/ch4/github"
)

const monthHours = 30 * 24
const yearHours = 365 * 24

func Search(filters []string) {
	result, err := github.SearchIssues(filters)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n\n", result.TotalCount)

	var ltOneMonth, ltOneYear, gtOneYear []*github.Issue

	for _, item := range result.Items {
		elapsed := time.Since(item.CreatedAt).Hours()
		if elapsed <= monthHours {
			ltOneMonth = append(ltOneMonth, item)
		} else if elapsed <= yearHours {
			ltOneYear = append(ltOneYear, item)
		} else {
			gtOneYear = append(gtOneYear, item)
		}
	}

	fmt.Println("Issues less than a month old")
	for _, item := range ltOneMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\nIssues less than a year old")
	for _, item := range ltOneYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\nIssues more than a year old")
	for _, item := range gtOneYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}