package gophorem

import "context"

// Tags returns  a list of tags that can be used to tag articles.
func (c *Client) Tags(ctx context.Context, args Arguments) (Tags, error) {
	var res Tags
	qp := args.toQueryParams().Encode()
	err := c.get(ctx, c.baseURL+"/tags?"+qp, &res)

	return res, err
}

type Tags []struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}
