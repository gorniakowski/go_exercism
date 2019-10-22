package protein

import (
	"errors"
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

//ErrStop occurs when STOP codon occures
var ErrStop = errors.New("Stop codon occured")

//ErrInvalidBase occurs when wrong codon is encountered
var ErrInvalidBase = errors.New("Invalid base")

//FromCodon retruns result of codon transaltion
func FromCodon(codon string) (string, error) {

	var protein = codons[codon]

	if protein == "STOP" {
		return "", ErrStop
	}
	if protein == "" {
		return "", ErrInvalidBase
	}

	return protein, nil
}

//FromRNA translates RNA :)
func FromRNA(rna string) ([]string, error) {
	var result []string
	for i := 0; i < len(rna)-1; i = i + 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			return result, nil
		}
		if err == ErrInvalidBase {
			return result, err
		}
		result = append(result, protein)

	}

	return result, nil
}
