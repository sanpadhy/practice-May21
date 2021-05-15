/* 
 * 0/1 Knapshak Problem: A search meathod to find the items that address some conditions. 
 ** Example: I have a list of items. each item has weight and value. 
 ** My bag can carry x kg of weight. How I need to select items that would give the maximum value. 

  Items : Wt    Value 
          1       1
          3       4
          4       5
          5       7

	Maximum weight I can carry is: 7 

	Find the items that would give maximum value. 

*/   

package main

import (
	"fmt"
	"math"
)

type item struct {
	weight int
	value  int
}

func KnapshakSearch(I []item, max int) int {

	var matrix [][]interface{}

	j := 0
	i := 0
	
	for j < len(I) {
		for i < max {
			if i == 0 || j == 0 {
				matrix[i][j] = 0 	
			} 
			
			if I[j].weight < i {
				matrix[i][j] = matrix[i][j-1] 
			} else {
				matrix[i][j] = math.Max(matrix[i][j-1].(float64), I[j].value + I[i - I[j].weight][j-1].(float64))
			}
			i += 1 
		} 
		j += 1
	} 
	
	fmt.Println("Maximum Value: ", I[i][j])

	return I[i][j]
}

func main() {
	var items = []item{{weight: 1, value: 1},
			  {weight: 3, value: 4},
			  {weight: 4, value: 5},
			  {weight: 5, value: 7}}

	maxVal := KnapshakSearch(items, 7)
	
	fmt.Println(maxVal)
}
