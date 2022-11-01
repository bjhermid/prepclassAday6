package main

import "fmt"

func nilaiABC(n int32) {
	switch {
	case n >= 90 && n <= 100:
		fmt.Println("Nilai A")
		break

	case n >= 80 && n <= 89:
		fmt.Println("Nilai B")
		break

	case n >= 70 && n <= 79:
		fmt.Println("Nilai C")
		break

	case n >= 60 && n <= 69:
		fmt.Println("Nilai D")
		break

	case n <= 59:
		fmt.Println("Nilai E")
		break
	}
}

func ganjilGenap(g int32) {
	switch {
	case g%2 == 0:
		fmt.Println("Genap")
		break
	case g%2 == 1:
		fmt.Println("Ganjil")
		break
	}
}

func kabisat(k int32) {
	switch {
	case k%100 == 0 && k%400 != 0:
		fmt.Println("Bukan Tahun Kabisat")
		break
	case k%400 == 0:
		fmt.Println("Tahun Kabisat")
		break
	case k%4 == 0:
		fmt.Println("Tahun Kabisat")
		break
	default:
		fmt.Println("Bukan Tahun Kabisat")
	}
}

func rating(u int32) {
	switch {
	case u >= 21:
		fmt.Println("Dewasa")
		break

	case u >= 13:
		fmt.Println("Remaja")
		break

	case u >= 9:
		fmt.Println("Bimbingan Ortu")
		break

	case u < 9:
		fmt.Println("Semua Usia")
		break
	}
}

func main() {

	// //02.md - Nilai
	nilaiABC(30)
	nilaiABC(75)

	// //03.md Ganjil Genap
	ganjilGenap(43)
	ganjilGenap(1032)

	//04.md-Tahun Kabisat
	kabisat(1900)
	kabisat(2000)
	kabisat(2001)
	kabisat(2016)

	//05.md- Rating Usia Film
	rating(15)
	rating(32)

}
