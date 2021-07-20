package c

func Cal(str string) string {
	data := []byte(str)
	sData := string(data)
	var sum = 0
	for i := 0; i < 5000; i++ {
		sum += i
	}
	return sData
}

func CalGoroutine(str string) string {
	go func() {
		for {
			DoCalGoroutine(str)
		}
	}()
	return ""
}

func DoCalGoroutine(str string) string {
	data := []byte(str)
	sData := string(data)
	var sum = 0
	for i := 0; i < 3000; i++ {
		sum += i
	}
	return sData
}
