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

		urutkanSelectionSort(nomorRumah)
		hasil[i] = nomorRumah
	}

	for _, rumahTersortir := range hasil {
		for _, rumah := range rumahTersortir {
			fmt.Printf("%d ", rumah)
		}
		fmt.Println()
	}
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
