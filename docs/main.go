/*
package commment here
*/
package blah

import (
	"fmt"
)

// One calls Other()
func One() {
	Other()
}

// Other prints "ASdf"
func Other() {
	fmt.Println("ASdf")
}

func main() {
	One()
}
