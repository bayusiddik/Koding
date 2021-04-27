package main

import "fmt"

func main ()	{
	var tahun,hasil1,hasil2,hasil3 int
	var tahun1,tahun2,tahun3 bool
	
	fmt.Print("Masukan Tahun: ")
	fmt.Scanln(&tahun)
	
	hasil1 = tahun%400
	hasil2 = tahun%4
	hasil3 = tahun%100
	
	tahun1 = (hasil1 == 0)
	tahun2 = (hasil2 == 0)
	tahun3 = (hasil3 != 0)
	
	fmt.Println("KABISAT = ", ((tahun1 || tahun2) && tahun3))
	
	fmt.Println("--------------------------")
	fmt.Println("MUHAMMAD BAYU SAMUDRA SIDDIK")
	fmt.Println("1301194031")
	



}