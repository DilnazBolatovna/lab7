package main

import (
 "fmt"
)

func oddNumbers(nums []int) []int {
 result := []int{}
 for _, num := range nums {
  if num%2 == 1 {
   result = append(result, num)
  }
 }
 return result
}
func main() {
 nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
 numbers := oddNumbers(nums)
 fmt.Println(numbers)
}