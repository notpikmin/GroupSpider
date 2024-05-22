package main

import "fmt"

func CheckForErr(err error) bool {
	if err != nil {
		fmt.Println(err)
	}
	return err != nil
}
