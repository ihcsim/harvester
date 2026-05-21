package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

type TextProcessor struct {
	normalizer norm.Form
}

func NewTextProcessor() *TextProcessor {
	return &TextProcessor{
		normalizer: norm.NFC,
	}
}

func (p *TextProcessor) Normalize(input string) string {
	return p.normalizer.String(input)
}

func (p *TextProcessor) NormalizeNFD(input string) string {
	return norm.NFD.String(input)
}

func (p *TextProcessor) NormalizeNFKC(input string) string {
	return norm.NFKC.String(input)
}

func ToTitle(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func ToUpper(s string) string {
	caser := cases.Upper(language.English)
	return caser.String(s)
}

func ToLower(s string) string {
	caser := cases.Lower(language.English)
	return caser.String(s)
}
