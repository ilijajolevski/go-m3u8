package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// ByteRange represents sub range of a resource
type ByteRange struct {
	Length *uint64
	Start  *uint64
}

// NewByteRange parses a text line in playlist file and returns a *ByteRange
func NewByteRange(text string) (*ByteRange, error) {
	if text == "" {
		return nil, nil
	}

	values := strings.Split(text, "@")

	lengthValue, err := strconv.ParseUint(values[0], 10, 64)

	if err != nil {
		return nil, err
	}

	br := ByteRange{Length: &lengthValue}

	if len(values) >= 2 {
		startValue, err := strconv.ParseUint(values[1], 10, 64)
		if err != nil {
			return &br, err
		}
		br.Start = &startValue
	}

	return &br, nil
}

func (br *ByteRange) String() string {
	if br.Start == nil {
		return fmt.Sprintf("%d", *br.Length)
	}

	return fmt.Sprintf("%d@%d", *br.Length, *br.Start)
}
