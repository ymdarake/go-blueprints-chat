package service

import (
	"errors"

	"github.com/ymdarake/go-blueprints-chat/app"
)

// ErrNoAvatarURL is an error thrown if an instance of GetAvatar cannot return its url.
var ErrNoAvatarURL = errors.New("chat: unable to get avatar url")

type GetAvatar interface {
	GetAvatarURL(c *app.Client) (string, error)
}

type GetAuthAvatar struct{}

var UseGetAuthAvatar GetAuthAvatar

// GetAvatarURL of GetAuthAvatar gets the avatar url from auth information
func (GetAuthAvatar) GetAvatarURL(c *app.Client) (string, error) {
	if url, ok := c.UserData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
