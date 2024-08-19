package helper

import "log"

func PanicIfError(err error)  {
	if err != nil {
		log.Fatalf("Error: %v", err) // Using Fatalf to log the error and exit
	}
}