package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetTime() (time.Time, error) {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().Add(response.ClockOffset), nil
}

func main() {
	time, err := GetTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	zone, _ := time.Zone()
	hour, min, sec := time.Clock()

	fmt.Printf("%s time: %d h %d min %d sec\n", zone, hour, min, sec)
}
