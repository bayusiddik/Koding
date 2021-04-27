package main

import(
		"fmt"
)
func main ()	{
	var put ArrType
	var ki string

	fmt.Println(rmax(put))
	fmt.Println(imin(put))
	fmt.Println(found(put,ki))
	fmt.Println(pos(put,ki))

	fmt.Println("_________________________________________")
	fmt.Println("Nama: MUHAMMAD BAYU SAMUDRA SIDDIK")
	fmt.Println("NIM: 1301194031")
}
const N = 2019

type RecType struct	{
	f1 string
	f2 int
	f3 float64
}
type ArrType [N] RecType

func rmax(data ArrType) int	{
	//fungsi yang mengembalikan nilai real(f3) terbesar dalam array
	
	var max float64
	var i int

	max = data[0].f3
	
	for i < len(data)	{
		if max < data[i].f3 {
			max = data[i].f3
		}
		i++
	}
	return int(max)
}
func imin(data ArrType) int {
	//fungsi yang mencari index dimana elemen dengan nilai integer terkecil tersimpan

	var min, indeks int
	var i int

	i = 0
	min = data[0].f2
	for i < len(data)	{
		if min < data[i].f2 {
			min = data[i].f2
			indeks = i
		}
		i++
	}
	return indeks
}
func found(data ArrType, key string) bool	{
	var i int

	i = 0
	for i < len(data) && data[i].f1 != key {
		i++
	}
	if data[i].f1 == key {
		return true
	} else {
		return false
	}
}
func pos(data ArrType, key string) int	{
	//jika data diketahui terurut membesar
	//implementasi binaty search untuk mengembalikan index dimana data key ditemukan.

	var kiri, kanan, mid int

	kiri = 0
	kanan = len(data)
	mid = (kiri + kanan) / 2
	
	for kiri < kanan && data[mid].f1 != key {
		if key > data[mid].f1 {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
		mid = (kiri + kanan) / 2
	}

	if data[mid].f1 == key {
		return mid
	} else {
		return -1
	}
}