package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkbenchShortcutUpdateRequest() *OapiWorkbenchShortcutUpdateRequest {
	return &OapiWorkbenchShortcutUpdateRequest{}
}

type OapiWorkbenchShortcutUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkbenchShortcutUpdateResponse
	AppId           string
	BizNo           string
	Icon            string
	Name            string
	PcShortcutUri   string
	ShortcutUri     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkbenchShortcutUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetAppId(appId2 string) {
	this.AppId = appId2
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetAppId() string {
	return this.AppId
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetBizNo(bizNo2 string) {
	this.BizNo = bizNo2
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetBizNo() string {
	return this.BizNo
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetPcShortcutUri(pcShortcutUri2 string) {
	this.PcShortcutUri = pcShortcutUri2
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetPcShortcutUri() string {
	return this.PcShortcutUri
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetShortcutUri(shortcutUri2 string) {
	this.ShortcutUri = shortcutUri2
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetShortcutUri() string {
	return this.ShortcutUri
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workbench.shortcut.update"
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkbenchShortcutUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams["biz_no"] = this.BizNo
	txtParams["icon"] = this.Icon
	txtParams["name"] = this.Name
	txtParams["pc_shortcut_uri"] = this.PcShortcutUri
	txtParams["shortcut_uri"] = this.ShortcutUri
	return txtParams
}
func (this *OapiWorkbenchShortcutUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppId, "appId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkbenchShortcutUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkbenchShortcutUpdateResponse struct {
	taobao.TaobaoResponse
	Result IsvOrgShortcutDto `json:"result,omitempty"`
}
type IsvOrgShortcutDto struct {
	BizNo string `json:"biz_no,omitempty"`
}
