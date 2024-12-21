package main

import (
	"fmt"
)

func calculateMedian(data []int) int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	n := len(data)
	if n%2 == 1 {
		return data[n/2]
	}
	return (data[(n/2)-1] + data[n/2]) / 2
}

func main() {
	var input string
	fmt.Println("Masukkan data (akhiri dengan -5313):")
	fmt.Scanln(&input)
	values := []rune(input)

	var data []int
	var num int
	for _, value := range values {
		if value == ' ' {
			continue
		} else if value == '0' {
			if len(data) == 0 {
				fmt.Println("Data kosong, tidak dapat menghitung median.")
			} else {
				median := calculateMedian(data)
				fmt.Println(median)
			}
			data = nil
		} else if value == '-' {
			break
		} else {
			num = int(value - '0')
			data = append(data, num)
		}
	}
}
