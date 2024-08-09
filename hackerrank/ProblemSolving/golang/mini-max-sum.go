package main

import (
  "os"
  "fmt"
  "sort"
  "bufio"
  "strconv"
  "strings"
)

func miniMaxSum(arr []int) {
  sum := 0
  for _, value := range(arr) { sum += int(value) }
  sort.Ints(arr)
  maximum := sum - int(arr[0])
  minimum := sum - int(arr[len(arr)-1])
  fmt.Printf("%d %d\n", minimum, maximum)
}

func main() {
  var arr []int
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  arrTemp := strings.Split(strings.TrimSpace(input)," ")
  for i:= 0; i < int(len(arrTemp)); i++ {
    arrItemTemp, _ := strconv.ParseInt(arrTemp[i], 10, 64)
    arrItem := int(arrItemTemp)
    arr = append(arr, arrItem)
  }
  miniMaxSum(arr)
}