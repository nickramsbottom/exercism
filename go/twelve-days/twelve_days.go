package twelve

var gifts = [12]string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, ",
}

var numbers = [12]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

// Song return all Twelve Days of Christmas verses
func Song() string {
	var song string

	for i := 1; i <= 12; i++ {
		song += Verse(i)

		if i != 12 {
			song += "\n"
		}
	}

	return song
}

// Verse return a single verse of the Twelve Days of Christmas
func Verse(n int) string {
	verse := "On the " + numbers[n-1] + " day of Christmas my true love gave to me: "

	for i := n - 1; i >= 0; i-- {
		if i == 0 && n > 1 {
			verse += "and " + gifts[i]
		} else {
			verse += gifts[i]
		}
	}

	return verse
}
