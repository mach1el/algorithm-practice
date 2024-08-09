package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func birthdayCakeCandles(candles []int32) int32 {
  max := candles[0]
  for i, _ := range(candles) { if candles[i] > max { max = candles[i] } }
  
  count := 0
  for i, _ := range(candles) { if candles[i] == max { count +=1 } }
  return int32(count)
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

  result := birthdayCakeCandles(arr)
  fmt.Println(result)
}