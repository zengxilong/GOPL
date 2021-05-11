package main

import (
	"GOPL/Ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

var class = make(map[string][]github.Issue)

const (
	LM string = "less than one month"
	LY string = "less than one Year"
	MY string = "more than one year"
)

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	nowYear, nowMonth, nowDay := time.Now().Date()
	for _, item := range result.Items {
		itemYear, itemMonth, itemDay := item.CreatedAt.Date()
		tempItem := *item
		switch {
		case itemYear == nowYear && itemMonth == nowMonth:
			class[LM] = append(class[LM], tempItem)
		case itemYear == nowYear && itemMonth != nowMonth:
			if nowMonth-itemMonth == 1 {
				if itemDay >= nowDay {
					class[LM] = append(class[LM], tempItem)
				} else {
					class[LY] = append(class[LY], tempItem)
				}
			} else {
				class[LY] = append(class[LY], tempItem)
			}
		case itemYear != nowYear:
			class[MY] = append(class[MY], tempItem)
		}
	}
	outPutIssues(LM)
	outPutIssues(LY)
	outPutIssues(MY)
}

func outPutIssues(choice string) {
	fmt.Printf("%d issues (%s):\n", len(class[choice]), choice)

	for _, item := range class[choice] {
		fmt.Printf("#%-5d %9.9s %v %.55s\n", item.Number, item.User.Login, item.CreatedAt, item.Title)
	}

}
