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

	//	msg := `{"touser":"oFIl_0QmonGnhRTP_sOwDF1jm4xM","template_id":"7ofyI1mXCvBdzgIy91H1uFbK31MNA6dAAvTs62dzBT4","url":"","topcolor":"#FF0000","data":{"User": {"value":"黄先生","color":"#173177"}}}`

	msg := `
{
    "touser": "oFIl_0QmonGnhRTP_sOwDF1jm4xM",
    "template_id": "7ofyI1mXCvBdzgIy91H1uFbK31MNA6dAAvTs62dzBT4",
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
}
