package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

func activity(name string, minutes int) (err error) {
	fmt.Println(name)

	for i := minutes; i > 0; i-- {
		var unit string
		if i == 1 {
			unit = "minute"
		} else {
			unit = "minutes"
		}
		err = beeep.Notify(name, strconv.Itoa(i)+" "+unit+" left", "")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Minute)
	}
	return
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("usage: go run main.go [activity1]:[duration1] [activity2]:[duration2] ...")
	} else {
		for _, arg := range args {
			splitArg := strings.Split(arg, ":")
			name := splitArg[0]
			minutes, err := strconv.Atoi(splitArg[1])
			if err != nil {
				panic(err)
			}
			err = activity(name, minutes)
			if err != nil {
				panic(err)
			}
		}
	}
}
