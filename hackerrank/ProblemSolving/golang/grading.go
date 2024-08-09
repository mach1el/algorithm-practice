package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func gradingStudents(grades []int32) []int32 {
  var finalGrade []int32
  for _, value := range(grades) {
    if (value < 38) { 
      finalGrade = append(finalGrade, value)
    } else if value %5 >= 3 {
      value += (5 - value%5)
      finalGrade = append(finalGrade, value)
    } else { finalGrade = append(finalGrade, value) }
  }
  return finalGrade
}

func main() {
  var arr []int32

  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  count, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)
  
  for i := 0; i < int(count); i++ {
    input2, _ := reader.ReadString('\n')
    arrRowTemp := strings.Split(strings.TrimRight(input2," \t\r\n"), " ")
    
    for _, arrRowItem := range arrRowTemp {
      arrItemTemp, _ := strconv.ParseInt(arrRowItem, 10, 64)
      arrItem := int32(arrItemTemp)
      arr = append(arr, arrItem)
    }
  }
  result := gradingStudents(arr)
  for _, value := range(result) { fmt.Println(value) }
}