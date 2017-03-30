# file
A file manager for golang

# Installation
> go get github.com/hwfy/file

# Example

```go
package main

import (
	"fmt"

	"github/hwfy/file"
)

func main() {
	dir := "library"
	name := "food_1.go"
	path := dir + "/" + name
	files := map[string]string{
		"book_1.go": "hello book_1\n",
		"book_2.go": "hello book_2\n",
		"food_1.go": "hello food_1\n",
	}
	for name, content := range files {
		_, err := file.Save(dir, name, []byte(content))
		if err != nil {
			panic(err)
		}
	}
	names, err := file.Names(dir)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s : %s\n", dir, names)

	err = file.AppendContent(path, " test append\n")
	if err != nil {
		panic(err)
	}
	content, err := file.GetContent(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:\n %s", name, content)

	err = file.RemoveLine(path, " test append")
	if err != nil {
		panic(err)
	}
	content, err = file.GetContent(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:\n %s", name, content)

	// OutPut:
	library : map[book_2:Book2 book_1:Book1 food_1:Food1]
	food_1.go:
 	 hello food_1
 	 test append
	food_1.go:
 	 hello food_1
}
```
