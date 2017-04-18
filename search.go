/*
Copyright 2017 Mikael Berthe

Licensed under the MIT license.  Please see the LICENSE file is this directory.
*/

package gondole

import (
	"github.com/sendgrid/rest"
)

// Search search for contents (accounts or statuses) and returns a Results
func (g *Client) Search(query string, resolve bool) (*Results, error) {
	if query == "" {
		return nil, ErrInvalidParameter
	}

	params := make(apiCallParams)
	params["q"] = query
	if resolve {
		params["resolve"] = "true"
	}

	var results Results
	if err := g.apiCall("search", rest.Get, params, &results); err != nil {
		return nil, err
	}
	return &results, nil
}
