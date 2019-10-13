package main

import "fmt"

func main() {
   /* local variable definition */
   var temperature int = 70
   var weather string = "hot"

   switch temperature {
      case 70: weather = "warm"
      case 40,50,60: weather = "chilly"
      case 10,20,30 : weather = "cold"
      default: weather = "warm"
   }
   switch {
      case weather == "warm" :
         fmt.Printf("It is warm outside.\n" )
      case weather == "chilly" :
         fmt.Printf("It is chilly outside\n" )
      case weather == "cold" :
         fmt.Printf("It is cold.\n" )
      default:
         fmt.Printf("Unsure.\n" );
   }
   fmt.Printf("The weather is %s\n", weather );
}
