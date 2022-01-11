package phones

import (
	"io"
)

type PhoneReader string

func (p PhoneReader) Read(bs []byte) (int, error) {

	count := 0

	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			bs[count] = p[i]
			count++
		}
	}

	return count, io.EOF
}
