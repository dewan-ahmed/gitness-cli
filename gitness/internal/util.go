package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/urfave/cli/v2"
)

func HttpRequest(ctx *cli.Context, url string) ([]byte, error) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ctx.String("token")))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
