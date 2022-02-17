package main

func main() {
	// TYPE keyword

	type Celcius float32 // Celcius is now alias for float32 / alias [Эйлиэс]
	var tempreture Celcius = 13.5
	tempreture += 0

	type IDnumber int // IDnumber = int
	var id IDnumber = 123
	id += 0
}
