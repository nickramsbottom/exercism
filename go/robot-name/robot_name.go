package robotname

import (
	"errors"
	"math/rand"
	"strconv"
)

// Robot a named automated machine
type Robot struct {
	name string
}

var usedName = map[string]bool{}

// Name get the name from Robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, error := generate()

		if error != nil {
			return "", error
		}

		r.name = name
	}
	return r.name, nil
}

// Reset generate a new name for Robot
func (r *Robot) Reset() {
	r.name = ""
}

func generate() (string, error) {
	if len(usedName) == 10*10*10*26*26 {
		return "", errors.New("Namespace exhausted")
	}

	name := randomCapitalLetter() + randomCapitalLetter() + randomThreeDigitNumber()

	if usedName[name] {
		return generate()
	}

	usedName[name] = true
	return name, nil
}

func randomCapitalLetter() string {
	return string('A' - 1 + rand.Intn(27))
}

func randomThreeDigitNumber() string {
	return strconv.Itoa(rand.Intn(1000-10) + 10)
}
