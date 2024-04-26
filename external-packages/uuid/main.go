/*
Cuando instalamos GO tenemos algunos packetes, pero hay otros que hay que instalarlos ya que no nos vienen con GO
*/

// PACKETE UUID -> $ go get -u github.com/google/uuid

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	p := fmt.Println
	id1 := uuid.New()
	// podemos pasarlo a string
	id1.String()
	p(id1)

	uid := uuid.New()
	p(uid.Version())

}
