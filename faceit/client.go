package faceit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type auth struct {
	key string
}

type midFunc func(*http.Request) *http.Request

func fromContext(ctx context.Context) (*auth, error) {
	key, ok := ctx.Value("key").(string)
	if !ok || key == "" {
		return nil, errors.New("auth key not found")
	}
	return &auth{key: key}, nil
}

type Client struct {
	a      *auth
	url    string
	client *http.Client
}

func NewClient(base string, ctx context.Context) *Client {
	var err error
	c := new(Client)
	c.url = base
	c.a, err = fromContext(ctx)
	if err != nil {
		panic(err)
	}
	c.client = http.DefaultClient
	return c
}

func (c *Client) withBearer(r *http.Request) *http.Request {
	r.Header.Set("Authorization", c.a.key)
	return r
}

func (c *Client) get(url string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(c.withBearer(r))
}

func (c *Client) GetCS() (any, error) {
	u, err := url.JoinPath(c.url, "games", "cs2")
	if err != nil {
		return nil, err
	}
	r, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (c *Client) FindPlayer(nickname string) (*Player, error) {
	u, err := url.JoinPath(c.url, "players")
	if err != nil {
		return nil, err
	}
	u += "?nickname=" + nickname
	r, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var p Player
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *Client) GetPlayer(id string) (any, error) {
	u, err := url.JoinPath(c.url, "players", id)
	if err != nil {
		return nil, err
	}
	r, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (c *Client) GetStats(id string) (any, error) {
	u, err := url.JoinPath(c.url, "players", id, "stats", "cs2")
	if err != nil {
		return nil, err
	}
	fmt.Println(u)
	r, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (c *Client) GetMatches(id string) (MatchTab, error) {
	u, err := url.JoinPath(c.url, "players", id, "history")
	if err != nil {
		return MatchTab{}, err
	}
	fmt.Println(u)
	r, err := c.get(u)
	if err != nil {
		return MatchTab{}, err
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return MatchTab{}, err
	}
	var m MatchTab
	err = json.Unmarshal(data, &m)
	if err != nil {
		return MatchTab{}, err
	}
	return m, nil
}

func (c *Client) GetMatch(id string) (Match, error) {
	u, err := url.JoinPath(c.url, "matches", id)
	if err != nil {
		return Match{}, err
	}
	fmt.Println(u)
	r, err := c.get(u)
	if err != nil {
		return Match{}, err
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return Match{}, err
	}
	var m Match
	err = json.Unmarshal(data, &m)
	if err != nil {
		return Match{}, err
	}
	return m, nil
}

// face it not allowing downloading demos
//func (c *Client) GetDemo(demo string) (any, error) {
//	r, err := c.get(demo)
//	if err != nil {
//		return nil, err
//	}
//	defer r.Body.Close()
//	data, err := io.ReadAll(r.Body)
//	if err != nil {
//		return nil, err
//	}
//	return string(data[:1024]), nil
//}
