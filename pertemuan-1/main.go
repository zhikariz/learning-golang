package main

import (
	"fmt"
	"strconv"
)

func main() {
	var namaLengkap string = "Helmi Adi Prasetyo"
	var namaLengkap2 = "Helmi Adi Prasetyo"
	NamaDepan := "Helmi"
	NamaBelakang := "Adi Prasetyo"

	fmt.Println(namaLengkap)
	fmt.Println(namaLengkap2)
	fmt.Println("Panjang dari nama lengkapku adalah", len(namaLengkap))
	fmt.Println(NamaDepan + " " + NamaBelakang)
	fmt.Printf("Nama Lengkapku adalah %s %s %s\n", NamaDepan, NamaBelakang)

	umur := 20
	fmt.Println(umur)

	const PHI float32 = 3.14

	fmt.Printf("%T\n", PHI)
	fmt.Printf("%T\n", umur)

	umurString := strconv.Itoa(umur)
	fmt.Printf("%T\n", umurString)

	harga := "Rp. 100000"
	hargaString, err := strconv.Atoi(harga)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("tipe data dari harga sekarang %T\n", hargaString)
}
