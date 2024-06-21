package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func dayOfProgrammer(year int) {
  if year < 1918 {
    if year%4 != 0 { fmt.Printf("13.09.%d\n", year) } else { fmt.Printf("12.09.%d\n", year) }
  } else if year > 1918 {
    if (!(year%400==0 || (year%4==0 && year%100!=0))) { fmt.Printf("13.09.%d\n", year) } else { fmt.Printf("12.09.%d\n", year) }
  } else { fmt.Println("26.09.1918") }
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
  yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
  checkError(err)
  year := int(yearTemp)
  dayOfProgrammer(year)
}

func readLine(reader *bufio.Reader) string {
  str, _, err := reader.ReadLine()
  if err == io.EOF { return "" }
  return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
  if err != nil { panic(err) }
}