package main

import (
  "os"
  "io"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func birthday(s []int32, d int32, m int32) int {
  count := 0
  
  for i := 0; i < len(s); i++ {
    sum := 0
    for j := i; j < int(m) + i; j++ {
      if j < len(s) { sum += int(s[j]) }
    }
    if sum == int(d) { count += 1 }
  }
  return count
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

  sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var s []int32

  for i := 0; i < int(n); i++ {
    sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
    checkError(err)
    sItem := int32(sItemTemp)
    s = append(s, sItem)
  }

  firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  dTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
  checkError(err)
  d := int32(dTemp)

  mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
  checkError(err)
  m := int32(mTemp)

  result := birthday(s, d, m)

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