package adapter

import (
	"net/http"
	"net/url"
)

type quickForgetFlashSecret struct {
	client *http.Client
}

func QuickForgetFactory() quickForgetFlashSecret {
	return quickForgetFlashSecret{
		&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (qf quickForgetFlashSecret) SaveSecret(secret string) (string, error) {
	resp, err := qf.client.PostForm("https://quickforget.com/secret/submit/", url.Values{"secret": {secret}, "expire_after_views": {"1"}, "expire_after": {"24"}})
	if err != nil {
		panic(err)
	}
	secretUrl, err := resp.Location()
	if err != nil {
		panic(err)
	}
	//We have to make the first call
	_, err = http.Get(secretUrl.String())
	if err != nil {
		panic(err)
	}
	return secretUrl.String(), nil
}
