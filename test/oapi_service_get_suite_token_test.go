package test

import (
	"testing"

	"github.com/SmallTianTian/dingtalk"
	"github.com/SmallTianTian/dingtalk/model"
)

func Test_OapiServiceGetSuiteToken(t *testing.T) {
	client := dingtalk.NewDefaultDingTalkClient("https://oapi.dingtalk.com/service/get_suite_token")
	req := model.NewOapiServiceGetSuiteTokenRequest()
	req.SetSuiteKey(suiteKey)
	req.SetSuiteSecret(suiteSecrect)
	req.SetSuiteTicket(suiteTicket)
	err := client.Execute(req)
	if err != nil {
		t.Error("Client error:", err)
		return
	}

	if !req.Resp.IsSuccess() {
		t.Error("fail:", req.Resp.ErrMsg)
		return
	}
	t.Log("SuiteAccessToken:", req.Resp.SuiteAccessToken)
}
