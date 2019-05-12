package main

import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere 
// Returnerer en slice av type []byte
// 

func main() {
	byteslice1:=AllocateVar(250)
	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteslice1[0])
	aslice := Reslice(byteslice1, 50, 100)
	fmt.Println("&aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[50]")
	fmt.Println(&byteslice1[50])
	cpy := CopySlice(aslice)
	fmt.Println("&cpy[0]")
	fmt.Println(&cpy[0])
	
}

func AllocateVar(b int) []byte {
	var slice []byte
	slice = make([]byte, b)
	return slice
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
//
func AllocateMake(b int) []byte {
	slice := make([]byte, b)
	return slice
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	slac := AllocateMake(lidx)
	slac = slc[lidx:uidx]
	return slac
}

// CopySlice ???

func CopySlice(slc []byte) []byte {
	copiedSlice := make([]byte, len(slc), cap(slc))
	copy(copiedSlice, slc)
	return copiedSlice
}