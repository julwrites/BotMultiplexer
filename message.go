package botmultiplexer

import (
	"fmt"
	"sort"
	"strings"
)

func Split(msg []byte, maxSize int) [][]byte {
	var splits [][]byte

	msgStr := string(msg)
	paragraphs := strings.SplitAfter(msgStr, "\n")

	var chunk string
	for _, para := range paragraphs {
		if len(chunk)+len(para) < maxSize {
			var group []string
			group = append(group, chunk)
			group = append(group, para)
			chunk = strings.Join(group, "\n")
		} else {
			splits = append(splits, []byte(chunk))
			chunk = para
		}
	}
	// Any leftovers should be accounted for
	if len(chunk) > 0 {
		splits = append(splits, []byte(chunk))
	}

	return splits
}

type BoldFormatter func(string) string
type ItalicsFormatter func(string) string
type SuperscriptFormatter func(string) string

type FormatType string

const (
	Bold        FormatType = "*"
	Italics     FormatType = "_"
	Superscript FormatType = "^"
	Null        FormatType = "/"
)

type FormatBlock struct {
	Start int
	End   int
	Type  FormatType
}

func NextFormatBlock(str string) FormatBlock {
	var candidates []int
	var formattypes = []string{
		string(Bold),
		string(Italics),
		string(Superscript),
	}

	for _, f := range formattypes {
		candidates = append(candidates, strings.Index(str, f))
	}

	sort.Ints(candidates)
	c := string(str[candidates[0]])

	var block FormatBlock

	block.Start = candidates[0]
	block.End = strings.Index(str[block.Start+1:], c)
	if block.Start != -1 && block.End != -1 {
		block.Type = FormatType(c)
	} else {
		block.Type = Null
	}

	return block
}

func Format(msg []byte, bold BoldFormatter, ita ItalicsFormatter, sup SuperscriptFormatter) []byte {
	str := string(msg)

	var outStr string

	block := NextFormatBlock(str)
	for block.Type != Null {

		var fmtStr string

		switch block.Type {
		case Bold:
			fmtStr = bold(string(str[block.Start+1 : block.End]))
			break
		case Italics:
			fmtStr = ita(string(str[block.Start+1 : block.End]))
			break
		case Superscript:
			fmtStr = sup(string(str[block.Start+1 : block.End]))
			break
		default:
			fmtStr = string(str[block.Start+1 : block.End])
		}

		outStr = fmt.Sprintf("%s%s%s%s", outStr, str[:block.Start], fmtStr, str[block.End+1:])

		block = NextFormatBlock(str[block.End+1:])
	}

	start := strings.Index(str, "^")
	for start != -1 {
		end := strings.Index(str[start+1:], "^")

		if end != -1 {
			var replace string
			for _, c := range str[start+1 : end] {
				replace = fmt.Sprintf("%s%s", replace, sup(string(c)))
			}
		}
	}

	return []byte(outStr)
}
