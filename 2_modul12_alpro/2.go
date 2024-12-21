package main

import (
	"fmt"
)

func main() {
	var jumlahDaerah int
	fmt.Scanln(&jumlahDaerah)

	hasil := make([][]int, jumlahDaerah)
	for i := 0; i < jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Scan(&jumlahRumah)
		if jumlahRumah <= 0 || jumlahRumah >= 1000000 {
			fmt.Println("Jumlah rumah harus dalam rentang 1 sampai 999999.")
			return
		}

		nomorRumah := make([]int, jumlahRumah)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		ganjil, genap := pisahkanGanjilGenap(nomorRumah)
		urutkanSelectionSort(ganjil)
		urutkanSelectionSortTerbalik(genap)
		hasil[i] = append(ganjil, genap...)
	}

	for _, rumahTersortir := range hasil {
		for _, rumah := range rumahTersortir {
			fmt.Printf("%d ", rumah)
		}
		fmt.Println()
	}
}

func pisahkanGanjilGenap(arr []int) (ganjil, genap []int) {
	for _, angka := range arr {
		if angka%2 == 0 {
			genap = append(genap, angka)
		} else {
			ganjil = append(ganjil, angka)
		}
	}
	return
}

func urutkanSelectionSort(arr []int) {
	panjang := len(arr)
	for i := 0; i < panjang-1; i++ {
		indeksMin := i
		for j := i + 1; j < panjang; j++ {
			if arr[j] < arr[indeksMin] {
				indeksMin = j
			}
		}
		arr[i], arr[indeksMin] = arr[indeksMin], arr[i]
	}
}

func urutkanSelectionSortTerbalik(arr []int) {
	panjang := len(arr)
	for i := 0; i < panjang-1; i++ {
		indeksMax := i
		for j := i + 1; j < panjang; j++ {
			if arr[j] > arr[indeksMax] {
				indeksMax = j
			}
		}
		arr[i], arr[indeksMax] = arr[indeksMax], arr[i]
	}
}
