def luasPersegi(sisi):
    return sisi * sisi

def luasPersegiPanjang(panjang, lebar):
    return panjang * lebar

def luasLingkaran(jariJari):
    return 3.14 * jariJari * jariJari

def main():
    while (True):
        print("Program Menghitung Luas Bangun")
        print("------------------------------")
        print("1. Luas Persegi")
        print("2. Luas Persegi Panjang")
        print("3. Luas Lingkaran")
        print("0. Keluar")
        print("------------------------------")
        menu = int(input("Pilih Menu: "))
        print("")

        if (menu == 1):
            sisi = float(input("Masukkan sisi: "))
            luas = luasPersegi(sisi)
            print(f"Luas Persegi: {float(luas):.2f}\n")
        elif (menu == 2):
            panjang = float(input("Masukkan panjang: "))
            lebar = float(input("Masukkan lebar: "))
            luas = luasPersegiPanjang(panjang, lebar)
            print(f"Luas Persegi Panjang: {float(luas):.2f}\n")
        elif (menu == 3):
            jariJari = float(input("Masukkan jari-jari: "))
            luas = luasLingkaran(jariJari)
            print(f"Luas Lingkaran: {float(luas):.2f}\n")
        elif (menu == 0):
            print("Berhasil Keluar Program!")
            break
        else:
            print("Menu Tidak Valid!\n")

main()