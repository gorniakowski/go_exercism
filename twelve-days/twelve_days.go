package twelve

import "fmt"

var numbers []string = []string{"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

var items []string = []string{"a Partridge in a Pear Tree", "two Turtle Doves",
	"three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying",
	"seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing",
	"ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

//Verse returns a song lirics form a given verse.
func Verse(verse int) string {
	verse--
	begin := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", numbers[verse])
	if verse == 0 {
		return fmt.Sprintf("%s %s.", begin, items[0])
	}
	for i := verse; i > 0; i-- {
		if i == verse {
			begin = fmt.Sprintf("%s %s", begin, items[i])
			continue
		}

		begin = fmt.Sprintf("%s, %s", begin, items[i])

	}
	return fmt.Sprintf("%s, and %s.", begin, items[0])
}

//Song return a whole song
func Song() string {
	var result string
	for i := 1; i <= 12; i++ {
		result = fmt.Sprintf("%s\n%s", result, Verse(i))
	}
	return result[1:]
}
