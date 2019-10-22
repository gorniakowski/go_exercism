package strand

import "strings"

//ToRNA  given DNA strand returns its RNA complement
func ToRNA(dna string) string {

	transcription := func(i rune) rune {
		switch i {
		case 'G':
			i = 'C'
		case 'C':
			i = 'G'
		case 'T':
			i = 'A'
		case 'A':
			i = 'U'
		}
		return i
	}

	return strings.Map(transcription, dna)

}
