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
	markName  = "GOLANG_CLI_REMINDER"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage %s <hh:mm> <text message\n>", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)

	}

	if t == nil {
		fmt.Println("Unable to parse time")
	}

	if now.After(t.Time) {
		fmt.Println("set a future time")
		os.Exit(3)
	}

	diff := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information")

		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}

		fmt.Println("Msg from Child process...")

	} else {
		fmt.Println("Starting child process...")

		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))

		cmd.Stdout = os.Stdout // Redirect child process's standard output to terminal
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Child process started. Waiting for the reminder time...")
		fmt.Println("Reminder will be displayed after", diff.Round(time.Second))
	}
	os.Exit(0)

}
