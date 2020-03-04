//Nicholas Larsen
//March 3, 2020
//creates different arrays
package main

import "fmt"

func main() {
  var product [5]string
  var price [5]float64
  //creates the arrays

  product[0]="apples"
  product[1]="bananas"
  product[2]="oranges"
  product[3]="grapes"
  product[4]="plums"
  //set the elements of the array

  price[0]=4.23
  price[1]=6.1
  price[2]=5.3
  price[3]=2.9
  price[4]=3.3
  //sets the elements of the other array

  for i:=0; i<=4; i++{
    fmt.Println(product[i],"costs",price[i])
  }
}