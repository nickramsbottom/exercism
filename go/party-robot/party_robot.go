package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and stands out their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour, direction string, distance float64) string {
	m := Welcome(name)
	m += fmt.Sprintf("\nYou have been assigned to table %X. ", table)
	m += fmt.Sprintf("Your table is %s, exactly %.1f meters from here.\n", direction, distance)
	m += fmt.Sprintf("You will be sitting next to %s.", neighbour)
	return m
}
