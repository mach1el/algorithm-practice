package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func aVeryBigSum(ar []int64) int64 {
  var x int64 = 1000000000
  for _, value := range(ar) {
    x += value
  }
  return x - 1000000000
}

func main() {
  var ar []int64

  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  count, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)
  
  input2, _ := reader.ReadString('\n')
  arNum := strings.Split(strings.TrimSpace(input2)," ")

  for i := 0; i < int(count); i++ {
    arItem, _ := strconv.ParseInt(arNum[i], 10, 64)
    ar = append(ar, arItem)
  }

  result := aVeryBigSum(ar)
  fmt.Printf("%d\n", result)
}