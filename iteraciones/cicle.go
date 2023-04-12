package iteraciones

import (
	"fmt"
)

func DarVuelta() {
	for i := 10; i >= 0; i-- {

		if i == 6 {
			continue
		}

		fmt.Println(i)
	}
}
