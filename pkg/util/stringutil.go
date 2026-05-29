package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

type StringProcessor struct {
	normalizer norm.Form
}

func NewStringProcessor() *StringProcessor {
	return &StringProcessor{
		normalizer: norm.NFC,
	}
}

func (p *StringProcessor) Normalize(input string) string {
	return p.normalizer.String(input)
}

func (p *StringProcessor) NormalizeNFD(input string) string {
	return norm.NFD.String(input)
}

func (p *StringProcessor) NormalizeNFKC(input string) string {
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

func NormalizeUnicode(s string) string {
	return norm.NFC.String(s)
}
