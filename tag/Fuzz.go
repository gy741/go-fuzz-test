package tag

import "bytes"

func Fuzz(data []byte) int {
	b := bytes.NewReader(data)
	_, err := ReadFLACTags(b)
	if err != nil {
		return 0
	}
	return 1
}
