package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func hello() error { // hello needs to return error to be compatible with errgroup.Go
	fmt.Println("hello world")
	time.Sleep(time.Second * 10)
	fmt.Println("hello world ended")
	return nil // Return nil or an actual error
}

func hello1() error {
	fmt.Println("hello world 1")
	time.Sleep(time.Second * 1)
	return errors.New("hello world 1 goes wrong")
}

func main() {
	g, cancel := errgroup.WithContext(context.Background())
	g.SetLimit(2) // Set a limit of 2 concurrent goroutines

	// Pass the functions as values, not the result of their execution
	g.Go(hello)  // Correct way: pass the function hello
	g.Go(hello1) // Correct way: pass the function hello1

	select {
	case <-cancel.Done():
		{
			err := cancel.Err()
			if err != nil {
				fmt.Printf("Error occurred %s", err.Error())
			}
		}
	}

	_ := g.Wait()
	fmt.Println("Done")
}
