package main

import (
  "os"
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
)

func diagonalDifference(arr [][]int32) int32 {
  var a, b int32
  for i, _ := range(arr) {
    a += arr[i][i]
    b += arr[i][len(arr)-1-i]
  }
  result := math.Abs(float64(a) - float64(b))
  return int32(result)
}

func main() {
  var arr [][]int32

  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  count, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)
  
  for i := 0; i < int(count); i++ {
    input2, _ := reader.ReadString('\n')
    arrRowTemp := strings.Split(strings.TrimRight(input2," \t\r\n"), " ")
    
    var arrRow []int32
    for _, arrRowItem := range arrRowTemp {
      arrItemTemp, _ := strconv.ParseInt(arrRowItem, 10, 64)
      arrItem := int32(arrItemTemp)
      arrRow = append(arrRow, arrItem)
    }
    arr = append(arr, arrRow)
  }
  result := diagonalDifference(arr)
  fmt.Printf("%d\n", result)
}