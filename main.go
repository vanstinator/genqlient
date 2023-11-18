// genqlient is a GraphQL client generator for Go.
//
// To run genqlient:
//
//	go run github.com/vanstinator/genqlient
//
// For programmatic access, see the "generate" package, below.  For
// user documentation, see the project [GitHub].
//
// [GitHub]: https://github.com/vanstinator/genqlient
package main

import (
	"github.com/vanstinator/genqlient/generate"
)

func main() {
	generate.Main()
}
