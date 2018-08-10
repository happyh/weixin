package main

import (
	"fmt"
	"github.com/happyh/weixin"
	"time"
)

func main() {

	appId := "wx5ea56cd113c8a5b4"
	appSecret := "3c17e5664dd3f9f05f960c49b0aca5f0"
	weixin.RefreshAccessToken(appId, appSecret)

	time.Sleep(3 * time.Second)
	openIds, total, count, nextOpenId, err := weixin.GetUserList()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v %#v %#v %#v", openIds, total, count, nextOpenId)
	}

	msg := `
{
    "touser": "oFIl_0QmonGnhRTP_sOwDF1jm4xM",
    "template_id": "tSjnK9dDVAJpHY5G-tqYhpIjkTCrJvvRyCNvDOc2vmM",
    "url": "",
    "topcolor": "#FF0000",
    "data": {
        "User": {
            "value": "阿广",
            "color": "#173177"
        },
        "Time": {
            "value": "20180419 11:47:30",
            "color": "#173177"
        },
        "Direct": {
            "value": "买入",
            "color": "#173177"
        },
        "Stock": {
            "value": "sh600016",
            "color": "#173177"
        },
        "Count": {
            "value": "1000",
            "color": "#173177"
        },
        "Price": {
            "value": "7.85",
            "color": "#173177"
        }
    }
}
	`
	err = weixin.SendTempleMsg(nextOpenId, msg)
	if err != nil {
		fmt.Println(err)
	}

	ct := &weixin.CustText{
		Content: "我们都是好孩子",
	}

	err = weixin.SendCustomMsg(nextOpenId, ct)
	if err != nil {
		fmt.Println(err)
	}

	TestCreateMenu()
}

func TestCreateMenu() {
	buttons := []weixin.Button{
		weixin.Button{
			Name: "扫码",
			SubButton: []weixin.Button{
				weixin.Button{
					Name: "扫码带提示",
					Type: weixin.MenuTypeScancodeWaitmsg,
					Key:  "rselfmenu_0_0",
				},
				weixin.Button{
					Name: "扫码推事件",
					Type: weixin.MenuTypeScancodePush,
					Key:  "rselfmenu_0_1",
				},
			},
		},
		weixin.Button{
			Name: "发图",
			SubButton: []weixin.Button{
				weixin.Button{
					Name: "系统拍照发图",
					Type: weixin.MenuTypePicSysphoto,
					Key:  "rselfmenu_1_0",
				},
				weixin.Button{
					Name: "拍照或者相册发图",
					Type: weixin.MenuTypePicPhotoOrAlbum,
					Key:  "rselfmenu_1_1",
				},
				weixin.Button{
					Name: "微信相册发图",
					Type: weixin.MenuTypePicWeixin,
					Key:  "rselfmenu_1_2",
				},
			},
		},
		weixin.Button{
			Name: "测试",
			SubButton: []weixin.Button{
				weixin.Button{
					Name: "腾讯",
					Type: weixin.MenuTypeView,
					URL:  "http://qq.com",
				},
				weixin.Button{
					Name: "发送位置",
					Type: weixin.MenuTypeLocationSelect,
					Key:  "rselfmenu_2_0",
				},
			},
		},
	}

	err := weixin.CreateMenu(buttons)
	if err != nil {
		fmt.Println("CreateMenu failed,err:", err)
	} else {
		fmt.Println("CreateMenu OK")
	}
}
