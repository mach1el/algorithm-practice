package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func timeConversion(s string) {
  arrTime := strings.Split(s, ":")
  
  for i, value := range arrTime { arrTime[i] = strings.TrimSpace(value) }
  
  str1, min, str3 := arrTime[0], arrTime[1], arrTime[2]
  
  var house, sec string
  if strings.HasSuffix(str3, "AM") {
    if str1 == "12" { house = "00" } else { house = str1 }
    sec = strings.ReplaceAll(str3, "AM", "")
  } else  {
    if str1 != "12" {
      intH, _ := strconv.Atoi(str1)
      house = strconv.Itoa(intH + 12)
    } else { house = str1 }
    sec = strings.ReplaceAll(str3, "PM", "")
  }
  fmt.Printf("%s:%s:%s\n", house, min, sec)
}

func main() {
  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')

  timeConversion(input1)
}