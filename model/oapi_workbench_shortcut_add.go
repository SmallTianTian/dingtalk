package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkbenchShortcutAddRequest() *OapiWorkbenchShortcutAddRequest {
	return &OapiWorkbenchShortcutAddRequest{}
}

type OapiWorkbenchShortcutAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkbenchShortcutAddResponse
	AppId           string
	BizNo           string
	Icon            string
	Name            string
	PcShortcutUri   string
	ShortcutUri     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkbenchShortcutAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkbenchShortcutAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkbenchShortcutAddRequest) SetAppId(appId2 string) {
	this.AppId = appId2
}
func (this *OapiWorkbenchShortcutAddRequest) GetAppId() string {
	return this.AppId
}
func (this *OapiWorkbenchShortcutAddRequest) SetBizNo(bizNo2 string) {
	this.BizNo = bizNo2
}
func (this *OapiWorkbenchShortcutAddRequest) GetBizNo() string {
	return this.BizNo
}
func (this *OapiWorkbenchShortcutAddRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiWorkbenchShortcutAddRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiWorkbenchShortcutAddRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiWorkbenchShortcutAddRequest) GetName() string {
	return this.Name
}
func (this *OapiWorkbenchShortcutAddRequest) SetPcShortcutUri(pcShortcutUri2 string) {
	this.PcShortcutUri = pcShortcutUri2
}
func (this *OapiWorkbenchShortcutAddRequest) GetPcShortcutUri() string {
	return this.PcShortcutUri
}
func (this *OapiWorkbenchShortcutAddRequest) SetShortcutUri(shortcutUri2 string) {
	this.ShortcutUri = shortcutUri2
}
func (this *OapiWorkbenchShortcutAddRequest) GetShortcutUri() string {
	return this.ShortcutUri
}
func (this *OapiWorkbenchShortcutAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workbench.shortcut.add"
}
func (this *OapiWorkbenchShortcutAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkbenchShortcutAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkbenchShortcutAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkbenchShortcutAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkbenchShortcutAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams["biz_no"] = this.BizNo
	txtParams["icon"] = this.Icon
	txtParams["name"] = this.Name
	txtParams["pc_shortcut_uri"] = this.PcShortcutUri
	txtParams["shortcut_uri"] = this.ShortcutUri
	return txtParams
}
func (this *OapiWorkbenchShortcutAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppId, "appId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkbenchShortcutAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkbenchShortcutAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkbenchShortcutAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64                   `json:"errcode,omitempty"`
	Errmsg  string                  `json:"errmsg,omitempty"`
	Result  IsvOrgShortcutStatusDto `json:"result,omitempty"`
}
type IsvOrgShortcutStatusDto struct {
	BizNo string `json:"biz_no,omitempty"`
}
