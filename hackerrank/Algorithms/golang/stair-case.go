package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

func StairCase(n int32) {
  text := "%" + strconv.Itoa(int(n)) + "s\n"
  for i := 1; i < int(n) + 1; i++ {
    fmt.Printf(text, strings.Repeat("#",i))
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)

  input1, _ := reader.ReadString('\n')
  count, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)

  StairCase(int32(count))
}