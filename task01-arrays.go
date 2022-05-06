package homework

func average(input [15]float32) (result float32) {
	var acum float32 = 0
	counter := 0

	for _, e := range input {
		if e != 0 {
			counter++
			acum += e
		}
	}

	if counter == 0 {
		return 0
	}

	result = acum / float32(counter)

	return
}
