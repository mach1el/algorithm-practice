package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func compareTriplets(a []int32, b []int32) (int, int) {
  Alice := 0
  Bob := 0
  
  for i := range(len(a)) {
    if a[i] > b[i] {
      Alice += 1
    } else if a[i] < b[i] {
      Bob += 1
    }
  }
  return Alice, Bob
}

func main() {
  var a, b []int32

  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  aTemp := strings.Split(strings.TrimSpace(input1)," ")
  
  input2, _ := reader.ReadString('\n')
  bTemp := strings.Split(strings.TrimSpace(input2)," ")

  for i := 0; i < len(aTemp); i++ {
    aItemTemp, _ := strconv.ParseInt(aTemp[i], 10, 64)
    aItem := int32(aItemTemp)
    a = append(a, aItem)
  }

  for i := 0; i < len(bTemp); i++ {
    bItemTemp, _ := strconv.ParseInt(bTemp[i], 10, 64)
    bItem := int32(bItemTemp)
    b = append(b, bItem)
  }

  Alice, Bob := compareTriplets(a, b)
  fmt.Println(Alice, Bob)
}