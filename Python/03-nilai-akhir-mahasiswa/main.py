def nilaiHuruf(total):
    grade = ""
    if total >= 85:
        grade = "A"
    elif total >= 70:
        grade = "B"
    elif total >= 55:
        grade = "C"
    elif total >= 40:
        grade = "D"
    else:
        grade = "E"
    return grade

def nilaiTotal(tugas, uts, uas):
    PERSENTASE_TUGAS = 0.3
    PERSENTASE_UTS = 0.3
    PERSENTASE_UAS = 0.4

    total = (tugas * PERSENTASE_TUGAS) + (uts * PERSENTASE_UTS) + (uas * PERSENTASE_UAS)
    return total

def main():
    tugas = int(input('Nilai Tugas = '))
    uts = int(input('Nilai UTS = '))
    uas = int(input('Nilai UAS = '))

    total = nilaiTotal(tugas, uts, uas)
    grade = nilaiHuruf(total)

    print(f"Nilai Akhir = {float(total):.2f}")
    print(f"Grade = {grade}")

main()