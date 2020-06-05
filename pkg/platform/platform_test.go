package platform

import (
	"fmt"
	"testing"
)

// Message tests
func TestSplit(t *testing.T) {
	chunks := Split("This is a block of text that has been delimited by spaces only", " ", 20)

	if len(chunks) != 4 {
		t.Errorf(fmt.Sprintf("Failed TestSplit multiple chunks scenario, got %d instead of 4", len(chunks)))
	}

	mono := Split("This is a block of text that has been delimited by spaces and nothing else", " ", 100)

	if len(mono) != 1 {
		t.Errorf("Failed TestSplit single chunk scenario")
	}
}

func TestNextFormatBlock(t *testing.T) {
	italics := NextFormatBlock(" _Italics_ *Bold* ^1234^ Text", 0)

	if italics.Type != Italics {
		t.Errorf("Failed TestNextFormatBlock italics format blocks scenario")
	}

	bold := NextFormatBlock(" _Italics_ *Bold* ^1234^ Text", 10)

	if bold.Type != Bold {
		t.Errorf("Failed TestNextFormatBlock bold format blocks scenario")
	}

	sup := NextFormatBlock(" _Italics_ *Bold* ^1234^ Text", 17)

	if sup.Type != Superscript {
		t.Errorf("Failed TestNextFormatBlock bold format blocks scenario")
	}

	null := NextFormatBlock("Text only no formatting", 0)

	if null.Type != Null {
		t.Errorf("Failed TestNextFormatBlock no format blocks scenario")
	}
}

// func Format(str string, normal NormalFormatter, bold BoldFormatter, ita ItalicsFormatter, sup SuperscriptFormatter) string {
// 	var outStr string

// 	str = normal(str)

// 	pos := 0
// 	for true {
// 		block := NextFormatBlock(str, pos)
// 		if block.Type == Null {
// 			break
// 		}

// 		outStr = outStr + str[pos:block.Start]   // Add any text before the formatter
// 		fmtStr := str[block.Start+1 : block.End] // Ignore the symbols

// 		switch block.Type {
// 		case Bold:
// 			fmtStr = bold(fmtStr)
// 			break
// 		case Italics:
// 			fmtStr = ita(fmtStr)
// 			break
// 		case Superscript:
// 			fmtStr = sup(fmtStr)
// 			break
// 		}

// 		outStr = outStr + fmtStr

// 		pos = block.End + 1
// 	}

// 	// Any leftovers
// 	outStr = outStr + str[pos:]

// 	return outStr
// }
