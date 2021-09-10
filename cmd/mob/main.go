package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/theguptaji/mob/pkg/rand"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "mp",
				Aliases: []string{"mp"},
				Usage:   "Make random pairs",
				Action:  makePairs,
			},
			{
				Name:    "history",
				Aliases: []string{"h"},
				Usage:   "show history of pairs",
				Action:  showHistory,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func makePairs(c *cli.Context) error {
	members := strings.Split(c.Args().First(), ",")
	s := c.Args().Get(1)
	maxMembers, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Making pairs out of...", members)
	rand.RandomPairs(members, maxMembers)
	return nil
}

func showHistory(c *cli.Context) error {
	//TODO: Implement history
	fmt.Println("Showing history:")
	return nil
}

func saveLast(c *cli.Context) error {
	//TODO: Implement saving last pairs
	fmt.Println("Saving last pairs:")
	return nil
}
