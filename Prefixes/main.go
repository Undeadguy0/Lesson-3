package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	var strings [4]string
	strings[0] = "tree"
	strings[1] = "trena"
	strings[2] = "trever"
	strings[3] = "treuy"
	var j int
	for i := 0; i < 7; i++ {
		if !((strings[0][i] == strings[1][i]) && (strings[1][i] == strings[2][i]) && (strings[3][i] == strings[2][i])) {
			j = i
			break
		}
	}
	fmt.Print(strings[0][0:j])

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
