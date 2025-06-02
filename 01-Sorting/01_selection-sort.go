package main

import "fmt"

func SelectionSort(arr []int) {
  for j := range len(arr)-1 {
      minIndex:=j
      for i:=j+1; i<len(arr); i++ {
          if arr[i] < arr[minIndex] {
              minIndex = i
          }
      }
      if minIndex != j {
		arr[j], arr[minIndex] = arr[minIndex], arr[j]  // Swap the elements
      }
  }
  fmt.Println(arr)
}
