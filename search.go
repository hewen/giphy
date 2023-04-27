package giphy

import (
	"fmt"
	"strings"
)

type SearchType string

const (
	SearchTypeGif      SearchType = "gifs"
	SearchTypeStickers SearchType = "stickers"
)

// Search returns a search response from the Giphy API
func (c *Client) Search(args []string, tp SearchType) (Search, error) {
	argsStr := strings.Join(args, " ")

	path := fmt.Sprintf("/%s/search?limit=%v&q=%s", tp, c.Limit, argsStr)
	req, err := c.NewRequest(path)
	if err != nil {
		return Search{}, err
	}

	var search Search
	if _, err = c.Do(req, &search); err != nil {
		return Search{}, err
	}

	return search, nil
}
