package main

import "fmt"

func Register() func() error {
	return func() error {
		fmt.Println("hello from two script")
		return nil
	}
}
