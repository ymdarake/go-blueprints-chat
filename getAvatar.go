package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrNoAvatarURL is an error thrown if an instance of GetAvatar cannot return its url.
var ErrNoAvatarURL = errors.New("chat: unable to get avatar url")

type GetAvatar interface {
	GetAvatarURL(c *Client) (string, error)
}

type GetAuthAvatar struct{}

var UseGetAuthAvatar GetAuthAvatar

// GetAvatarURL of GetAuthAvatar gets the avatar url from auth information
func (GetAuthAvatar) GetAvatarURL(c *Client) (string, error) {
	if url, ok := c.UserData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GetGravatarAvatar struct{}

var UseGetGravatarAvatar GetGravatarAvatar

func (GetGravatarAvatar) GetAvatarURL(c *Client) (string, error) {
	if email, ok := c.UserData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			// TODO: use const, cache
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatarURL
}
