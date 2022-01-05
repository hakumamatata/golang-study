package bird

import (
	"fmt"
	"inits/dog"
)

func init() {
	fmt.Println("bird bark!" + dog.DogName)
}
