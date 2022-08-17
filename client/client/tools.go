//go:build tools

package client

// make sure genqlient is installed as a dependency so go generate works
import _ "github.com/Khan/genqlient"
