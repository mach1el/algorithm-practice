package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func plusMinus(arr []int32) {
  var pos, neg, zero int
  for _, value := range(arr) {
    if value > 0 {
      pos += 1
    } else if value < 0 {
      neg += 1
    } else if value == 0 {
      zero += 1
    }
  }
  fmt.Println(float64(pos)/float64(len(arr)))
  fmt.Printf("%d\n", pos/len(arr))
  fmt.Printf("%d\n", neg/len(arr))
  fmt.Printf("%d\n", zero/len(arr))
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

  plusMinus(arr)
}
