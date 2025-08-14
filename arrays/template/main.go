package main

import (
	"fmt"
	"time"
)

func main() {
	// Define the date you want to convert to epoch time
	dateString := "2023-12-31"

	// Parse the date string into a time.Time object
	now := time.Now()

	// Convert the date to epoch time
	epochTime := now.Unix()
	convertUnix(uint64(epochTime))
	fmt.Println("Epoch time for", dateString, "is:", epochTime)
}

func convertUnix(epochTime uint64) {
	// Convert the epoch time back to a time.Time object
	t := time.Unix(int64(epochTime), 0)

	// Format the time as a string
	formattedTime := t.Format("2006-01-02 15:04:05")

	fmt.Println("Converted time from epoch:", formattedTime)
}
