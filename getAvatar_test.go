package main

import (
	"testing"
)

func TestGetAuthAvatar(t *testing.T) {
	var getAuthAvatar GetAuthAvatar
	client := new(Client)
	url, err := getAuthAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("should return ErrNoAvatarURL")
	}
	testUrl := "http://url-to-avatar"
	client.UserData = map[string]interface{}{"avatar_url": testUrl}
	url, err = getAuthAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("should not return err")
	}
	if url != testUrl {
		t.Error("should return its url")
	}
}
