package main

import "fmt"
import (
  "math/rand"
  "strings"
)

func main() {
  game()
  keepPlaying()
}

func game(){
  randNum := (rand.Int() % 10) + 1
  var userNum int

  fmt.Println("Choose a number between 1 and 10")
  fmt.Scan(&userNum)

  if userNum < randNum && userNum > 0{
    fmt.Println("Your number was too low.")
    fmt.Println("You chose" , userNum ," The correct number is ", randNum)
  } else if userNum == randNum{
    fmt.Println("Congratulations!!")
    fmt.Println("You chose", userNum, " and the correct number is ", randNum)
  }else if userNum > randNum && userNum < 11{
    fmt.Println("Your number was too high")
    fmt.Println("You chose", userNum, " The correct number is ", randNum)
  }else if userNum < 1 || userNum > 10 {
    fmt.Println("Please choose a number between 1 and 10")
  }
}

func keepPlaying() {
  var ans string
  fmt.Println("Would you like to play again?")
  fmt.Println("Y/N")
  fmt.Scan(&ans)
  if strings.ToLower(ans) == "y" || strings.ToLower(ans) == "yes" {
    main()
  }else if strings.ToLower(ans) == "n" || strings.ToLower(ans) == "no"{
    fmt.Println("Thank you for playing!")
  }
}
