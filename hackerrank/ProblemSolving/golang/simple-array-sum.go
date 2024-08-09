package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func simpleArraySum(arr []int32) int32 {
  var i int32 = 0
  for _, value := range(arr) {
    i += value
  }
  return i

}

func main() {
  var arr []int32
  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  count, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)
  
  input2, _ := reader.ReadString('\n')
  arrTemp := strings.Split(strings.TrimSpace(input2)," ")

  for i := 0; i < int(count); i++ {
    arrItemTemp, _ := strconv.ParseInt(arrTemp[i], 10, 64)
    arrItem := int32(arrItemTemp)
    arr = append(arr, arrItem)
  }

  result := simpleArraySum(arr)
  fmt.Println(result)
}