package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func pageCount(n int32, p int32) int32 {
  var num int32
  if n == 2 { 
    num = 0 
  } else {
    if (n - p == 1) {
      if n%2 == 0 { num = 1 } else { num = 0 }
    } else {
      start_page := p / 2
      end_page := (n - p) / 2
      if start_page < end_page { num = start_page } else { num = end_page }
    }
  }
  return num
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

  pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  p := int32(pTemp)

  result := pageCount(n, p)

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