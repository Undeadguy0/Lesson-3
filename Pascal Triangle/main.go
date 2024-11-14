package main

import (
	"fmt"
	"strconv"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	var rows int
	fmt.Scanf("%d", &rows)
	fmt.Println(rows)
	var row []int
	row = append(row, 1)
	var row2 []int
	for i := 0; i < rows; i++ {
		row2 = append(row2, 1)
		for j := 1; j < len(row); j++ {
			row2 = append(row2, row[j-1]+row[j])
		}
		row2 = append(row2, 1)
		row = row2
		row2 = nil
	}
	for i := 0; i < len(row); i++ {
		fmt.Print(strconv.Itoa(row[i]) + " ")
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
