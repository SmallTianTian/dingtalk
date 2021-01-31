package dingtalk

import (
	"fmt"
	"testing"

	"github.com/SmallTianTian/dingtalk/model"
	"github.com/SmallTianTian/dingtalk/utils"
)

const (
	corpID       = "dinga77de7de5e6e80e2a39a90f97fcb1e09"
	suiteKey     = "suitejvlnilepiigew0zh"
	suiteSecrect = "plBfBRoR2RNU_kn9JhOsRpzL8IzXFj6EHklSxHTKLt_eHqnlpJLm1Y87I6vN6WzS"
	suiteTicket  = "mhrrAtD0s8mYTmuF93BbD1pCrLKsBpeo9RZ3U1d0fWBiQKNQsdqwaFj4tSJGgLCkJ0ZXL2sg4x1YnrADD5IKRN"
	// ak           = "f0a961224936328e91c4070cec7476d9"
)

func Test_sign(t *testing.T) {
	// %252F9BhzXaseU8NzUoZoHahgZthjl9t1XYUGVenm4XPaN8%3D
	start := 1611846429991
	key := utils.GetCanonicalStringForIsv(int64(start), suiteTicket)
	signature, _ := utils.ComputeSignature(suiteSecrect, key)
	fmt.Println(signature)
	// System.out.println(DingTalkSignatureUtil.urlEncode(signature, "UTF8"))
}

func Test_GetCropToken(t *testing.T) {
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/service/get_corp_token")
	req := model.NewOapiServiceGetCorpTokenRequest()
	req.SetAuthCorpid("dingc365fcabbf733c3535c2f4657eb6378f")
	err := client.Execute4(req, "suiteKey", "suiteSecrect", "suiteTicket")
	if err != nil {
		fmt.Println("Client error:", err)
		return
	}

	if req.Resp.IsSuccess() {
		fmt.Println("succ", req.Resp.Body)
	} else {
		fmt.Println("fail", req.Resp.ErrMsg)
	}
}
