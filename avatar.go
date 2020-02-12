package main

import "errors"

var ErrNoAvatarUrl = errors.New("chat: could not find an avatar url")

type Avatar interface {
	GetAvatarUrl(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarUrl(c *client) (string, error) {
	rawURL, ok := c.userData["avatar_url"]
	if !ok {
		return "", ErrNoAvatarUrl
	}

	url, ok := rawURL.(string)
	if !ok {
		return "", ErrNoAvatarUrl
	}

	return url, nil
}
