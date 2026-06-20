package main

import ("fmt")

func nilaiHuruf(total float32) string {
	var grade string;
	if (total >= 85) {
		grade = "A"
	} else if (total >= 70) {
		grade = "B"
	} else if (total >= 55) {
		grade = "C"
	} else if (total >= 40) {
		grade = "D"
	} else {
		grade = "E"
	}
	return grade;
}

func nilaiTotal(tugas int, uts int, uas int) float32 {
	const persentaseTugas float32 = 0.3
	const persentaseUTS float32 = 0.3
	const persentaseUAS float32 = 0.4

	var total float32 = (float32(tugas) * persentaseTugas) + (float32(uts) * persentaseUTS) + (float32(uas) * persentaseUAS)
	return total
}

func main() {
	var tugas, uts, uas int

	fmt.Print("Nilai Tugas = ")
  	fmt.Scan(&tugas)

	fmt.Print("Nilai UTS = ")
  	fmt.Scan(&uts)

	fmt.Print("Nilai UAS = ")
  	fmt.Scan(&uas)

	total := nilaiTotal(tugas, uts, uas)
	grade := nilaiHuruf(total)

	fmt.Printf("Nilai Akhir = %.2f\n", total)
	fmt.Printf("Grade = %v", grade)
}