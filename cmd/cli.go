package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	MarkName  = "Golang CLI Reminder"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:%s <hh:mm> <message>", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}

	if t == nil {
		fmt.Println("No time found")
		os.Exit(2)
	}

	if now.After(t.Time) {
		fmt.Println("Time is in the past")
		os.Exit(3)
	}

	diff := t.Time.Sub(now)

	if os.Getenv(MarkName) == markValue {
		time.Sleep(diff)
		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")

		if err != nil {
			fmt.Println("Error showing notification:", err)
			os.Exit(4)
		}

	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", MarkName, markValue))
		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			os.Exit(5)
		}
		fmt.Println("Reminder will be displayed after", diff.Round(time.Second))
		os.Exit(0)
	}

}
