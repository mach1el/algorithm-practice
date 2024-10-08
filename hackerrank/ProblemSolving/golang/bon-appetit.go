package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func bonAppetit(bill []int32, k int32, b int32) {
  cost := func(slice []int32) int32 { var total int32; for _, val := range slice { total += val }; return total }(append(bill[:k], bill[k+1:]...))
  if cost/2 == b {
    fmt.Println("Bon Appetit")
  } else { fmt.Println(b - cost/2) }
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
  checkError(err)
  n := int32(nTemp)

  kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
  checkError(err)
  k := int32(kTemp)

  billTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var bill []int32

  for i := 0; i < int(n); i++ {
    billItemTemp, err := strconv.ParseInt(billTemp[i], 10, 64)
    checkError(err)
    billItem := int32(billItemTemp)
    bill = append(bill, billItem)
  }

  bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  b := int32(bTemp)

  bonAppetit(bill, k, b)
}

func readLine(reader *bufio.Reader) string {
  str, _, err := reader.ReadLine()
  if err == io.EOF { return "" }
  return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
  if err != nil { panic(err) }
}