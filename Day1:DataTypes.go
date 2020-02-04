package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    iInt, _ := strconv.ParseUint(scanner.Text(), 10, 64)
    scanner.Scan()
    dDouble, _ := strconv.ParseFloat(scanner.Text(), 64)
    scanner.Scan()
    sString := scanner.Text()

    fmt.Println(iInt + i)
    fmt.Printf("%.1f\n", dDouble + d)
    fmt.Println(s + sString)
}