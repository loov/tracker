package main

import (
	"fmt"

	"github.com/loov/tracker/issue"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var tracker issue.Tracker

	id, err := tracker.Open(issue.Info{
		Caption: "Hello",
		Desc:    "World",
	})
	check(err)

	fmt.Println("Issue ID: ", id)

	info, err := tracker.Load(id)
	check(err)

	fmt.Println("Info", info)

	issues, err := tracker.List()
	check(err)

	fmt.Println("Issues", issues)

	check(tracker.Close(info.ID))

	info2, err := tracker.Load(info.ID)
	check(err)

	fmt.Println("Info2", info2)
}
