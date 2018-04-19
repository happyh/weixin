package weixin

import (
	"testing"
	"time"
)

func TestAccessToken(t *testing.T) {
	appId := "wx5ea56cd113c8a5b4"
	appSecret := "3c17e5664dd3f9f05f960c49b0aca5f0"
	RefreshAccessToken(appId, appSecret)

	time.Sleep(3 * time.Second)
	t.Log(AccessToken())
}
