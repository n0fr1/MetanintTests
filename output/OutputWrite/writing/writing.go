package writing

type PhoneWriter struct{}

func (p PhoneWriter) Write(bs []byte) (int, error) {

	buffer := make([]byte, len(bs))

	if len(bs) == 0 {
		return 0, nil
	}

	count := 0

	for i := 0; i < len(bs); i++ {

		if bs[i] >= '0' && bs[i] <= '9' {
			buffer[count] = bs[i]
			count++

		}

	}

	copy(bs, buffer) //копируем содержимое слайса с номером телефона без знаков в первоначальный слайс

	return count, nil

}
