package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkbenchShortcutGetguideuriRequest() *OapiWorkbenchShortcutGetguideuriRequest {
	return &OapiWorkbenchShortcutGetguideuriRequest{}
}

type OapiWorkbenchShortcutGetguideuriRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkbenchShortcutGetguideuriResponse
	AppId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkbenchShortcutGetguideuriRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) SetAppId(appId2 string) {
	this.AppId = appId2
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) GetAppId() string {
	return this.AppId
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workbench.shortcut.getguideuri"
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	return txtParams
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppId, "appId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkbenchShortcutGetguideuriRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkbenchShortcutGetguideuriResponse struct {
	taobao.TaobaoResponse
	Errcode  int64  `json:"errcode,omitempty"`
	Errmsg   string `json:"errmsg,omitempty"`
	GuideUri string `json:"guide_uri,omitempty"`
}
