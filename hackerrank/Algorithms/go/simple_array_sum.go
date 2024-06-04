package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func simpleArraySum(ar []int32) int32 {
  var i int32 = 0
  for _, value := range(ar) {
    i += value
  }
  return i

}

func main() {
  var ar []int32
  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  arCount, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)
  
  input2, _ := reader.ReadString('\n')
  arTemp := strings.Split(strings.TrimSpace(input2)," ")

  for i := 0; i < int(arCount); i++ {
    arItemTemp, _ := strconv.ParseInt(arTemp[i], 10, 64)
    arItem := int32(arItemTemp)
    ar = append(ar, arItem)
  }

  result := simpleArraySum(ar)
  fmt.Println(result)
}