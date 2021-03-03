package weixin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	log "github.com/happyh/go-logging"
)

// tick := time.Tick(7 * time.Second)
const refreshTimeout = 120 * time.Minute
const tokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

type accessToken struct {
	AccessToken string `json:"access_token"` //	获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   //	凭证有效时间，单位：秒
	mutex       sync.RWMutex
}

// AccessToken 取最新的 access_token，必须使用这个接口取，内部已经加锁
var AccessToken func() string

// RefreshAccessToken 定时刷新 access_token
func RefreshAccessToken(appId, appSecret string) {
	// 内部变量，外部不可以调用
	var _token = &accessToken{}
	url := fmt.Sprintf(tokenURL, appId, appSecret)
	/*
		go func() {
			tick := time.Tick(refreshTimeout)
			for {
				new := refresh(url)

				log.Log().Noticef("old access token %+v", _token)
				log.Log().Noticef("new access token %+v", new)

				_token.mutex.Lock()
				_token.AccessToken = new.AccessToken
				_token.ExpiresIn = new.ExpiresIn
				_token.mutex.Unlock()

				<-tick // 等待下一个时钟周期到来
			}
		}()
	*/
	AccessToken = func() string {
		//		_token.mutex.RLock()
		//		defer _token.mutex.RUnlock()

		// get new token every fun
		new := refresh(url)

		//log.Log().Noticef("old access token %+v", _token)
		//log.Log().Noticef("new access token %+v", new)

		_token.mutex.Lock()
		_token.AccessToken = new.AccessToken
		_token.ExpiresIn = new.ExpiresIn
		_token.mutex.Unlock()

		//log.Log().Notice("AccessToken:", _token.AccessToken)
		return _token.AccessToken
	}
}

func refresh(url string, ns ...int) (new *accessToken) {
	n := 0
	if len(ns) > 0 {
		n = ns[0]
	}

	var err error
	defer func() {
		if err != nil {
			log.Log().Error(err)
			time.Sleep(3 * time.Minute)
			if n < 9 {
				n++
				new = refresh(url, n)
			}
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		log.Log().Errorf("weixin accesstoken refresh failed,err:", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Log().Errorf("weixin accesstoken refresh failed,err:", err)
		return
	}
	resp.Body.Close()

	new = &accessToken{}
	err = json.Unmarshal(body, new)
	if err != nil {
		log.Log().Errorf("weixin accesstoken refresh failed,err:", err)
		return
	}

	if new.AccessToken == "" {
		log.Log().Errorf("weixin accesstoken refresh failed,body:", body)
	}

	return new
}
