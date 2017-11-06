package helpers

import "log"

// IsError ...
func IsError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
	}
	return (err != nil)
}

// Must ...
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
