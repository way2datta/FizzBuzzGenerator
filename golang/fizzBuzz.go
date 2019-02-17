package main

import (
  "fmt"
  "strconv"
)


func main() {
  for i:=0; i<=100; i++ {
    print(i)
  }
}

func print(number int)  {
    fmt.Println(FizzBuzz(number))
}

func FizzBuzz(number int)  string {

  if numberIsDivisibleByThree(number) && numberIsDivisibleByFive(number) {
    return ("FizzBuzz")
    } else if numberIsDivisibleByFive(number) {
      return ("Buzz")
    } else if numberIsDivisibleByThree(number) {
      return ("Fizz")
    }

    return (strconv.Itoa(number))

}

func numberIsDivisibleByThree(number int) bool {
  return number%3 ==0
}

func numberIsDivisibleByFive(number int) bool {
  return number%5 ==0
}
