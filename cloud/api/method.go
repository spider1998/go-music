package api

import (
	"cloud/app"
	"cloud/code"
	"cloud/entity"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)



func GetArtList(args map[string]string, urls string) (res entity.Artlists, err error) {
	URL, err := url.Parse(urls)
	if err != nil {
		return
	}
	query := URL.Query()
	for key, val := range args {
		query.Add(key, val)
	}
	URL.RawQuery = query.Encode()
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return
	}
	resp, err := Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		app.Logger.Warn().Str("response", string(b)).Msg("received error response.")
		var result code.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
