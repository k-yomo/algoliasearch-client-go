// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractExtraHeaders(opts ...interface{}) *opt.ExtraHeadersOption {
	merged := make(map[string]string)

	for _, o := range opts {
		if v, ok := o.(*opt.ExtraHeadersOption); ok {
			for key, value := range v.Get() {
				merged[key] = value
			}
		}
	}

	if len(merged) == 0 {
		return nil
	}

	return opt.ExtraHeaders(merged)
}
