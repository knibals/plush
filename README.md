# Plush library

Plush is a universal hug dispender and incidentally **learn Go module**.

Usage:

Import the library in your code and start the hugging!

```
package main

import (
	"fmt"
	"github.com/knibals/plush"
)

func main() {
	f:= plush.Gimme("kiss") // or plush.Gimme("hug")
	fmt.Println(f) // display an Emoji
}

```
