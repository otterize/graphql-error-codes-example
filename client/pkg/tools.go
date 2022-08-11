//go:build tools

package pkg

// make sure gqlgen&genqlient are installed as a dependency
import _ "github.com/99designs/gqlgen"
import _ "github.com/Khan/genqlient"
