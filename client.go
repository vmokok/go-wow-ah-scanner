package ahsc

import (
	"context"
	"errors"
	"fmt"
	"github.com/vmokok/go-wow-ah-scanner/internal/gamedata"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"log"
	"net/http"
)

const (
	tokenURL = "https://oauth.battle.net/token"
)

type Client struct {
	HTTPClient *http.Client
}

type Response struct {
	Body       []byte
	Header     *Header
	StatusCode int
}

var ErrEmptyCredentials = errors.New("empty credentials")

type Option func(*Client)

// New returns new http.Client ready for querying battle.net Game Data API
func New(ctx context.Context, clientID, clientSecret string, options ...Option) (*Client, error) {
	//TODO: validate
	if len(clientID) == 0 || len(clientSecret) == 0 {
		log.Fatalf("Failed to get token: %v", ErrEmptyCredentials)
	}
	c := &Client{
		HTTPClient: http.DefaultClient,
	}
	if len(options) > 0 {
		for _, option := range options {
			option(c)
		}
	}

	ctx = context.WithValue(ctx, oauth2.HTTPClient, c.HTTPClient)

	clCfg := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}

	c.HTTPClient = clCfg.Client(ctx)

	_, err := clCfg.Token(ctx)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}

	return c, nil
}

// WithCustomHTTPClient allows to override settings of default http.Client
func WithCustomHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = client
	}
}

func (c *Client) fetch(ctx context.Context, gd gamedata.Gd, opts ...string) (gamedata.Gd, error) {
	const op = "battlenet.client.fetch"
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		gd.QueryPath(opts...),
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating request in %s: %w", op, err)
	}

	req.Header.Set(HeaderAccept, "application/json")
	req.Header.Set(HeaderBattleNetNamespace, gd.Namespace())
	//TODO: validate input http.TimeFormat
	req.Header.Set(HeaderIfModifiedSince, gd.LastUpdated())

	q := req.URL.Query()
	//TODO: move to cfg
	q.Set("locale", "en_GB")
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request in %s: %w", op, err)
	}

	defer func(Body io.ReadCloser) { _ = Body.Close() }(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body in %s: %w", op, err)
	}

	err = gamedata.Map(body, gd)
	if err != nil {
		return nil, fmt.Errorf("error mapping game data in %s: %w", op, err)
	}

	return gd, nil
}
