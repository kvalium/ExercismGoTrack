// Package protein translate RNA sequences into proteins
package protein

import (
	"errors"
)

// ErrStop on stop codon
var ErrStop = errors.New("stop encountered")

// ErrInvalidBase on unknown codon
var ErrInvalidBase = errors.New("invalid base")

type protein string

const (
	methionine    protein = "Methionine"
	phenylalanine         = "Phenylalanine"
	leucine               = "Leucine"
	serine                = "Serine"
	tyrosine              = "Tyrosine"
	cysteine              = "Cysteine"
	tryptophan            = "Tryptophan"
	stop                  = "STOP"
)

var codonToProteinLookup = map[string]protein{
	"AUG": methionine,
	"UUU": phenylalanine,
	"UUC": phenylalanine,
	"UUA": leucine,
	"UUG": leucine,
	"UCU": serine,
	"UCC": serine,
	"UCA": serine,
	"UCG": serine,
	"UAU": tyrosine,
	"UAC": tyrosine,
	"UGU": cysteine,
	"UGC": cysteine,
	"UGG": tryptophan,
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

var chunkSize = 3

// FromRNA translates a codon sequence (RNA) into an array of proteins
func FromRNA(rna string) ([]string, error) {
	var proteins []string

	for _, codon := range chunkCodons(rna) {
		var protein, err = FromCodon(codon)
		if err == ErrStop {
			return proteins, nil
		}
		if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

// FromCodon returns the proteins associated to the given codon
func FromCodon(codon string) (string, error) {
	var protein, err = codonToProtein(codon)
	if err != nil {
		return "", err
	}
	if protein == stop {
		return "", ErrStop
	}
	return string(protein), nil
}

func codonToProtein(c string) (protein, error) {
	if protein, ok := codonToProteinLookup[c]; ok {
		return protein, nil
	}
	return "", ErrInvalidBase
}

func chunkCodons(s string) (chunks []string) {
	for i := 0; i <= len(s); i += chunkSize {
		chunkLimit := i + chunkSize
		if chunkLimit > len(s) {
			continue
		}
		chunks = append(chunks, s[i:chunkLimit])
	}
	return
}
