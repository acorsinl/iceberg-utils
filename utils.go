package utils

import (
	"code.google.com/p/go-uuid/uuid"
	"net/url"
	"time"
)

func NewUUID() string {
	return uuid.New()
}

func UnixTimestamp() int32 {
	return int32(time.Now().UTC().Unix())
}

func GetQueryParameters(urlString string) (url.Values, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	m, _ := url.ParseQuery(u.RawQuery)
	return m, nil
}
