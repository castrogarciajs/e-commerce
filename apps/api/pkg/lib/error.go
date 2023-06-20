package lib

import "fmt"

func PanicError(err error) error {
	fmt.Printf("ERROR: %s", err)
	return err
}
