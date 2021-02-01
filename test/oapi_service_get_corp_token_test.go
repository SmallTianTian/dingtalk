package test

import (
	"testing"

	"github.com/SmallTianTian/dingtalk"
	"github.com/SmallTianTian/dingtalk/model"
)

func Test_OapiServiceGetCorpToken(t *testing.T) {
	client := dingtalk.NewDefaultDingTalkClient("https://oapi.dingtalk.com/service/get_corp_token")
	req := model.NewOapiServiceGetCorpTokenRequest()
	req.SetAuthCorpid(corpID)
	if err := client.Execute4(req, suiteKey, suiteSecrect, suiteTicket); err != nil {
		t.Error("Client error:", err)
		return
	}
	if !req.Resp.IsSuccess() {
		t.Error("failed:", req.Resp.ErrMsg)
		return
	}

	t.Log(req.Resp.Body)
}
