package tstmod

import (
	"fmt"
	"errors"
)	 

// Hi returns a friendly greeting
func Hi(name string, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s", name), nil
	case "es":
		return fmt.Sprintf("Hola, %s", name), nil
	case "fr":
		return fmt.Sprintf("Bonjur, %s", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
