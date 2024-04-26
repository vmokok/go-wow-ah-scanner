package ahsc

import (
	"context"
	"github.com/vmokok/go-wow-ah-scanner/internal/gamedata"
	"strconv"
)

func gdResult[T gamedata.Gd](c *Client, ctx context.Context, a T, opts ...string) (T, error) {
	gd, err := c.fetch(ctx, a, opts...)
	if err != nil {
		return *new(T), err
	}
	return gd.(T), nil
}

func (c *Client) ConnectedRealmsIndex(ctx context.Context) (*gamedata.ConnectedRealmIndex, error) {
	return gdResult(c, ctx, &gamedata.ConnectedRealmIndex{})
}

func (c *Client) ItemSearch(ctx context.Context, opts ...string) (*gamedata.ItemSearch, error) {
	return gdResult(c, ctx, &gamedata.ItemSearch{}, opts...)
}

func (c *Client) Auctions(ctx context.Context, connectedRealmId int) (*gamedata.Auctions, error) {
	return gdResult(c, ctx, &gamedata.Auctions{}, strconv.Itoa(connectedRealmId))
}
