package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
  var totalApple int
  for _, value := range(apples) {
    i := a + value
    if s <= i && i <= t { totalApple+=1 }
  }
  var totalOrange int
  for _, value := range(oranges) {
    i := b + value
    if s <= i && i <= t { totalOrange+=1 }
  }
  fmt.Println(totalApple)
  fmt.Println(totalOrange)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

  firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  sTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
  checkError(err)
  s := int32(sTemp)

  tTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
  checkError(err)
  t := int32(tTemp)

  secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  aTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
  checkError(err)
  a := int32(aTemp)

  bTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
  checkError(err)
  b := int32(bTemp)

  thirdMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  mTemp, err := strconv.ParseInt(thirdMultipleInput[0], 10, 64)
  checkError(err)
  m := int32(mTemp)

  nTemp, err := strconv.ParseInt(thirdMultipleInput[1], 10, 64)
  checkError(err)
  n := int32(nTemp)

  applesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var apples []int32

  for i := 0; i < int(m); i++ {
    applesItemTemp, err := strconv.ParseInt(applesTemp[i], 10, 64)
    checkError(err)
    applesItem := int32(applesItemTemp)
    apples = append(apples, applesItem)
  }

  orangesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

  var oranges []int32

  for i := 0; i < int(n); i++ {
    orangesItemTemp, err := strconv.ParseInt(orangesTemp[i], 10, 64)
    checkError(err)
    orangesItem := int32(orangesItemTemp)
    oranges = append(oranges, orangesItem)
  }

  countApplesAndOranges(s, t, a, b, apples, oranges)
}

func readLine(reader *bufio.Reader) string {
  str, _, err := reader.ReadLine()
  if err == io.EOF { return ""}
  return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) { 
  if err != nil { panic(err) }
}