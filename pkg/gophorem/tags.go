package gophorem

import "context"

// Tags returns  a list of tags that can be used to tag articles. See https://docs.dev.to/api/#tag/tags
func (c *Client) Tags(ctx context.Context, args Arguments) (Tags, error) {
	var res Tags
	qp := args.toQueryParams().Encode()
	err := c.get(ctx, c.baseURL+"/tags?"+qp, &res)

	return res, err
}

// The structs in this file was generated via https://mholt.github.io/json-to-go/.

type Tags []struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}
