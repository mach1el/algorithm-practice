package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func countingValleys(steps int32, path string) int32 {
  _ = steps
  var count, valley int32
  for _, c := range path {
    if c == 'U' {
      count += 1
      if count == 0 { valley += 1 }
    } else { count -= 1}
  }
  return valley
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
  checkError(err)

  defer stdout.Close()

  writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

  stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  steps := int32(stepsTemp)

  path := readLine(reader)

  result := countingValleys(steps, path)

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