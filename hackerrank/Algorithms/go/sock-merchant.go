package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func sockMerchant(n int32, ar []int32) int32 {
  var pair int32
  counter := make(map[int32]int32)
  func(slice []int32) { for _, element := range slice { counter[element]++ } }(ar)
  for _, count := range counter { pair += count/2 }
  return pair
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  n := int32(nTemp)

  arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var ar []int32

  for i := 0; i < int(n); i++ {
    arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
    checkError(err)
    arItem := int32(arItemTemp)
    ar = append(ar, arItem)
  }

  result := sockMerchant(n, ar)

  fmt.Fprintf(writer, "%d\n", result)

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