package main

import (
  "os"
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
  var result string
  if v2 >= v1 || int(math.Abs(float64(x2-x1)))%int(math.Abs(float64(v1-v2))) != 0 {  result = "NO" } else { result = "YES"}
  return result
}

func main() {
  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  multipleLineInput := strings.Split(strings.TrimSpace(input1)," ")
  
  x1Temp, _ := strconv.ParseInt(multipleLineInput[0], 10, 64)
  x1 := int32(x1Temp)

  v1Temp, _ := strconv.ParseInt(multipleLineInput[1], 10, 64)
  v1 := int32(v1Temp)

  x2Temp, _ := strconv.ParseInt(multipleLineInput[2], 10, 64)
  x2 := int32(x2Temp)

  v2Temp, _ := strconv.ParseInt(multipleLineInput[3], 10, 64)
  v2 := int32(v2Temp)

  result := kangaroo(x1, v1, x2, v2)
  fmt.Println(result)

}