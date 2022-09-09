# word-validator
Word validator is a library that checks if a message contains an inappropriate word or not

## This library have only two methods:
1. DefaultValidator
2. CustomValidator

NOTE: Both methods return a bool value. True (true) would be returned if there is a match

## Installation
```ssh
go get github.com/hisyntax/word-validator
```
## DefaultValidator
#### This checks if a string of words contains an inappropriate word or not using a predefined array of words to check

```go 
package main

import (
	"fmt"

	"github.com/hisyntax/word-validator/validate"
)

func main() {
	message := "hello there"
	
	res := validate.DefaultValidator(message)

    //the response would be false since there is no match
	fmt.Println(res)

}
```

## DefaultValidator
#### This checks if a string of words contains an inappropriate word(or any word at all) using a customed array of words passed int to check

```go 
package main

import (
	"fmt"

	"github.com/hisyntax/word-validator/validate"
)

func main() {
	message := "nigra"
	words := []string{
		"ass","bastard", "bitch","cock",
        "dick", "dyke","fuck","nigra",
	}
	res := validate.CustomValidator(words, message)

    //this would return true since there is a match
	fmt.Println(res)

}
```