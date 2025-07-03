package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	currentIndex := m + n - 1 // Индекс позиции для вставки элемента
	i, j := m-1, n-1          // Индексы для массивов
	for i >= 0 && j >= 0 {
		// Сравниваем и вставляем в конец массива nums1
		if nums1[i] > nums2[j] {
			nums1[currentIndex] = nums1[i]
			i--
		} else {
			nums1[currentIndex] = nums2[j]
			j--
		}
		currentIndex--
	}

	// Если остались элементы в nums2 (например, m=0), копируем их
	for j >= 0 {
		nums1[currentIndex] = nums2[j]
		j--
		currentIndex--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

	for _, value := range nums1 {
		fmt.Print(value, " ")
	}
	fmt.Print("\r\n")
}
