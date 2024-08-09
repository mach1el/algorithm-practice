package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
  var prices []int32
  
  if len(keyboards) > len(drives) {
    for _, kPrice := range keyboards { arTemp := func(slice []int32) []int32 { var ar []int32; for _, dPrice := range slice { ar = append(ar, kPrice + dPrice) }; return ar }(drives); prices = append(prices, arTemp...) }
  } else {
    for _, dPrice := range drives { arTemp := func(slice []int32) []int32 { var ar []int32; for _, kPrice := range slice { ar = append(ar, dPrice + kPrice) }; return ar }(keyboards); prices = append(prices, arTemp...) }
  }
  
  spent := func(slice []int32) []int32 { var tempSlice []int32; for _, value := range slice { if value <= b { tempSlice = append(tempSlice, value) } }; return tempSlice }(prices)

  var max int32
  if len(spent) == 0 {
    max = -1
  } else {
    for _, value := range spent { if value > max { max = value } }
  }

  return max
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 1024 * 1024)

  bnm := strings.Split(readLine(reader), " ")

  bTemp, err := strconv.ParseInt(bnm[0], 10, 64)
  checkError(err)
  b := int32(bTemp)

  nTemp, err := strconv.ParseInt(bnm[1], 10, 64)
  checkError(err)
  n := int32(nTemp)

  mTemp, err := strconv.ParseInt(bnm[2], 10, 64)
  checkError(err)
  m := int32(mTemp)

  keyboardsTemp := strings.Split(readLine(reader), " ")

  var keyboards []int32

  for keyboardsItr := 0; keyboardsItr < int(n); keyboardsItr++ {
    keyboardsItemTemp, err := strconv.ParseInt(keyboardsTemp[keyboardsItr], 10, 64)
    checkError(err)
    keyboardsItem := int32(keyboardsItemTemp)
    keyboards = append(keyboards, keyboardsItem)
  }

  drivesTemp := strings.Split(readLine(reader), " ")

  var drives []int32

  for drivesItr := 0; drivesItr < int(m); drivesItr++ {
    drivesItemTemp, err := strconv.ParseInt(drivesTemp[drivesItr], 10, 64)
    checkError(err)
    drivesItem := int32(drivesItemTemp)
    drives = append(drives, drivesItem)
  }

  moneySpent := getMoneySpent(keyboards, drives, b)

  fmt.Fprintf(writer, "%d\n", moneySpent)

  writer.Flush()
}

func readLine(reader *bufio.Reader) string {
  str, _, err := reader.ReadLine()
  if err == io.EOF { return "" }
  return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
  if err != nil { panic(err) }
}