package botmultiplexer

import "strings"

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
