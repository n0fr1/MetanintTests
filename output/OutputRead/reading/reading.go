package reading

import "io"

type PhoneReader string

func (ph PhoneReader) Read(p []byte) (int, error) { //реализуем метод Read интерфейса Reader

	count := 0

	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}

	}

	return count, io.EOF
}
