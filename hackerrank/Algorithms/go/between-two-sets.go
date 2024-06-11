package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func getTotalX(a []int32, b []int32) int32 {
  var min_value, max_value int32 = 1, 1
  for i, _ := range(a) { if a[i] < min_value { min_value = a[i] } }
  for i, _ := range(b) { if b[i] > max_value { max_value = b[i] } }

  count := 0
  var arr, brr []int32
  for i := min_value; i < max_value+1; i++ {
    for _, value := range(a) { arr = append(arr, i%value) }
    for _, value := range(b) { brr = append(arr, value%i) }
    if 
  }
  fmt.Println(arr, brr)
  return min_value
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

  mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
  checkError(err)
  m := int32(mTemp)

  arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var arr []int32

  for i := 0; i < int(n); i++ {
    arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
    checkError(err)
    arrItem := int32(arrItemTemp)
    arr = append(arr, arrItem)
  }

  brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var brr []int32

  for i := 0; i < int(m); i++ {
    brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
    checkError(err)
    brrItem := int32(brrItemTemp)
    brr = append(brr, brrItem)
  }

  total := getTotalX(arr, brr)

  fmt.Fprintf(writer, "%d\n", total)

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