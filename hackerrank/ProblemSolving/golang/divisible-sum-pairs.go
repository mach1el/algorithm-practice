package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func divisibleSumPairs(k int32, ar []int32) int {
  count := 0
  for index, value := range ar {
    arrVal := func(slice []int32) []int32 { var tempArr []int32; for _, tempValue := range slice { tempArr = append(tempArr, value + tempValue) }; return tempArr }(ar[index+1:])
    for _, value := range arrVal { if value%k == 0 { count += 1 } }
  }
  return count
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
  checkError(err)
  n := int32(nTemp)

  kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
  checkError(err)
  k := int32(kTemp)

  arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var ar []int32

  for i := 0; i < int(n); i++ {
    arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
    checkError(err)
    arItem := int32(arItemTemp)
    ar = append(ar, arItem)
  }

  result := divisibleSumPairs(k, ar)

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