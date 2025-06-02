package main

import "fmt"

func insertionSort(arr []int) {
  for i:=0; i<len(arr); i++ {
      for j:=i; j>0; j-- {
          if (arr[j] < arr[j-1]) {
              temp := arr[j]
              arr[j] = arr[j-1]
              arr[j-1] = temp
          } else {
              break
          }
      }
  }
  fmt.Println(arr)
}