// This file will be generated to include all customer specific middlewares
package registry

import (
	"github.com/ruprict/vulcand/plugin"
	"github.com/ruprict/vulcand/plugin/cbreaker"
	"github.com/ruprict/vulcand/plugin/connlimit"
	"github.com/ruprict/vulcand/plugin/ratelimit"
	"github.com/ruprict/vulcand/plugin/rewrite"
	"github.com/ruprict/vulcand/plugin/trace"
)

func GetRegistry() *plugin.Registry {
	r := plugin.NewRegistry()

	specs := []*plugin.MiddlewareSpec{
		ratelimit.GetSpec(),
		connlimit.GetSpec(),
		rewrite.GetSpec(),
		cbreaker.GetSpec(),
		trace.GetSpec(),
	}

	for _, spec := range specs {
		if err := r.AddSpec(spec); err != nil {
			panic(err)
		}
	}

	return r
}
