package main

import (
	"fmt"
	"os"
	"shopping-sunday-service/pkg/sunday"
	"time"
)

func main() {
	var (
		err  error
		date = time.Now()
	)
	if len(os.Args) > 1 {
		passedDate := os.Args[1]
		if date, err = time.Parse(sunday.ShoppingSundayFormat, passedDate); err != nil {
			return
		}
	}

	if shopping, reasons := sunday.IsShopping(date); shopping {
		fmt.Printf("%d %s %d is a shopping Sunday!\n", date.Day(), date.Month(), date.Year())
		return
	} else {
		fmt.Printf("%d %s %d is not a shopping Sunday\n", date.Day(), date.Month(), date.Year())
		for _, reason := range reasons {
			fmt.Println(reason.Message)
		}
	}
}
