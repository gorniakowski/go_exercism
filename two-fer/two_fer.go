//Package twofer provides function returning a string containg a name of a person.
package twofer

const start, end = "One for ", ", one for me."

// ShareWith takes persons name as string and returns string
//containing the name given. If no name given funtion will retrun string with word "you"
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return start + name + end

}
