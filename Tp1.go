/*	NIM: 1301194031
	Nama: Muhammad Bayu Samudra Siddik
	Tugas Pendahuluan Modul1
*/ 
package main

import "fmt"

func main()	{
	var satu,dua,tiga,temp string
	
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
	
	fmt.Println("Selesai")
}