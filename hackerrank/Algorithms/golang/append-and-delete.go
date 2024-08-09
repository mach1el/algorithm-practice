package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
  "math"
)

func appendAndDelete(s string, t string, k int32) string {
  var i int32
  for x := 0; x < int(math.Min(float64(len(s)),float64(len(t)))); x++ {
    i+=1
    if s[x] != t[x] { break; }
  }

  num := k - int32(len(s)) - int32(len(t)) + 2*i
  var result string

  if num < 0 {
    result = "No"
  } else if s == t {
    result = "Yes"
  } else if num <= 2*i {
    if num%2 != 0 {
      result = "No"
    } else if num%k == 0 {
      result = "No"
    } else { result = "Yes" }
  } else { result = "Yes" }
  
  return result
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  s := readLine(reader)

  t := readLine(reader)

  kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  k := int32(kTemp)

  result := appendAndDelete(s, t, k)

  fmt.Fprintf(writer, "%s\n", result)

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