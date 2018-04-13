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

func TestGetGravatarAvatar(t *testing.T) {
	var getGravatarAvatar GetGravatarAvatar
	client := new(Client)

	//NOTE: this pair of email and url is written in the official doc
	client.UserData = map[string]interface{}{"email": "MyEmailAddress@example.com"}
	url, err := getGravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GetGravatarAvatarURL.GetAvatarURL should not return err")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("incorrect value %s returned from GetGravatarAvatarURL.GetAvatarURL", url)
	}
}
