package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot a named automated machine
type Robot struct {
	name string
}

var usedNames = map[string]bool{}

// Name get the name from Robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) == 10*10*10*26*26 {
		return "", errors.New("Namespace exhausted")
	}

	r.name = newName()
	for usedNames[r.name] {
		r.name = newName()
	}
	usedNames[r.name] = true

	return r.name, nil
}

// Reset generate a new name for Robot
func (r *Robot) Reset() {
	r.name = ""
}

func newName() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)

	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}
