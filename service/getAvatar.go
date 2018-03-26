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
