package main

import (
	"fmt"
	"github.com/eurozulu/cmdhandle/cmdhandle"
	"log"
)

func main() {
	mainline.Handle("admin", doMyCommand)
	mainline.Handle("admin list", doMyCommandList)
	mainline.Handle("admin info", doMyCommandInfo)
	mainline.Handle("admin add", doMyCommandAdd)

	if err := mainline.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func doMyCommand(cmd mainline.CommandLine) error {
	fmt.Println("admin <command>")
	fmt.Println("Commands are:")
	fmt.Println("\tlist\t\t\t list all the things")
	fmt.Println("\tinfo <name>\t\t list the thing details")
	fmt.Println("\tadd <name> <url>\t adds new thing with url")
	return nil
}

func doMyCommandList(cmd mainline.CommandLine) error {
	_, ok := cmd.Flags().Get("verbose", "v")
	if ok {
		fmt.Println("doing verbose list")
	} else {
		fmt.Println("doing regular list")
	}
	return nil
}

func doMyCommandInfo(cmd mainline.CommandLine) error {
	fmt.Printf("doing info")
	return nil
}

func doMyCommandAdd(cmd mainline.CommandLine) error {
	if len(cmd.Args()) < 2 {
		return fmt.Errorf("must provide the name and url to add")
	}
	fmt.Printf("doing %v with name '%s' and '%s' url\n", cmd, cmd.Args()[0], cmd.Args()[1])
	return nil
}
