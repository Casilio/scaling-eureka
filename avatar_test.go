package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)

	_, err := authAvatar.GetAvatarUrl(client)
	if err != ErrNoAvatarUrl {
		t.Error("AuthAvatar.GetAvatarUrl should return ErrNoAvatarUrl when there is no avatar url")
	}

	avatarURL := "http://avatar_url.com"
	client.userData = map[string]interface{}{"avatar_url": avatarURL}

	url, err := authAvatar.GetAvatarUrl(client)
	if url != avatarURL || err != nil {
		t.Error("AuthAvatar.GetAvatarUrl should return avatar url and no error")
	}
}
