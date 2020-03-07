/*
Package twofer implements functionality for sharing
things with others.
*/
package twofer

// ShareWith returns are string containing the given name
// if the name is empty 'you' is used
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
