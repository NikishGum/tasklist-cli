package utils

import "log"

func WrapErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}