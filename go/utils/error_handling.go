package utils

import "log"

func Check(e error) {
	if e != nil {
		log.Panicln(e)
	}
}
