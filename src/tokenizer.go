package src

import ()

type tokenType string

type Token struct {
	Type    tokenType
	Literal string
}

const (
	HEADER1     tokenType = "HEADER1"
	HEADER2     tokenType = "HEADER2"
	HEADER3     tokenType = "HEADER3"
	BOLD        tokenType = "BOLD"
	ITALIC      tokenType = "ITALIC"
	BLOCKQUOTE  tokenType = "BLOCKQUOTE"
	OLIST       tokenType = "OLIST" // ordered list
	ULIST       tokenType = "ULIST" // unordered list
	CODE        tokenType = "CODE"
	HR          tokenType = "HR" // horizontal rule
	LINK        tokenType = "LINK"
	IMAGE       tokenType = "IMAGE"
	HIGHLIGHT   tokenType = "HIGHLIGHT"
	SUBSCRIPT   tokenType = "SUBSCRIPT"
	SUPERSCRIPT tokenType = "SUPERSCRIPT"
)

var lookup = map[string]tokenType{
	"#":   HEADER1,
	"##":  HEADER2,
	"###": HEADER3,
	"**":  BOLD,
	"*":   ITALIC,
	">":   BLOCKQUOTE,
	"-":   ULIST,
	"`":   CODE,
	"---": HR,
	"[":   LINK,
	"!":   IMAGE,
	"==":  HIGHLIGHT,
	"~":   SUBSCRIPT,
	"^":   SUPERSCRIPT,
}

func LookupToken(litteral string) tokenType {
	if token, exists := lookup[litteral]; exists {
		return token
	}

	return "UNKNOWN"
}

// func openFile(filepath string) {
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		fmt.Printf("Error opening file", err.Error())
// 	}
// }
