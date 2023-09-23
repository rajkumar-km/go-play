// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/rajkumar-km/go-play/go-excercises/ch04/10-github-search/github"
)

// go run . repo:golang/go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	
	// Sort the github issues based on created time
	sort.Slice(result.Items, func (i int, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})

	// Print age wise
	fmt.Printf("Found %d issues:\n", result.TotalCount)
	i := 0

	// Less than a month old
	fmt.Printf("\nLess than a month old\n")
	for ; i < len(result.Items); i++ {
		item := result.Items[i]
		duration := time.Now().Sub(item.CreatedAt)
		if duration > 31*24*time.Hour {
			break
		}
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)		
	}

	// Less than a year old
	fmt.Printf("\nLess than a year old\n")
	for ; i < len(result.Items); i++ {
		item := result.Items[i]
		duration := time.Now().Sub(item.CreatedAt)
		if duration > 12*31*24*time.Hour {
			break
		}
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)		
	}

	// More than a year old
	fmt.Printf("\nMore than a year old\n")	
	for ; i < len(result.Items); i++ {
		item := result.Items[i]
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)		
	}
}