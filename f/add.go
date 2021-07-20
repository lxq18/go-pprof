package f

import "go-pprof/c"

func Add(str string) string {
	c.Cal(str)
	data := []byte(str)
	sData := string(data)
	var sum = 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	return sData
}
