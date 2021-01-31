package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkbenchShortcutDeleteRequest() *OapiWorkbenchShortcutDeleteRequest {
	return &OapiWorkbenchShortcutDeleteRequest{}
}

type OapiWorkbenchShortcutDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkbenchShortcutDeleteResponse
	AppId           string
	BizNo           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkbenchShortcutDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkbenchShortcutDeleteRequest) SetAppId(appId2 string) {
	this.AppId = appId2
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetAppId() string {
	return this.AppId
}
func (this *OapiWorkbenchShortcutDeleteRequest) SetBizNo(bizNo2 string) {
	this.BizNo = bizNo2
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetBizNo() string {
	return this.BizNo
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workbench.shortcut.delete"
}
func (this *OapiWorkbenchShortcutDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkbenchShortcutDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkbenchShortcutDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams["biz_no"] = this.BizNo
	return txtParams
}
func (this *OapiWorkbenchShortcutDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppId, "appId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkbenchShortcutDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkbenchShortcutDeleteResponse struct {
	taobao.TaobaoResponse
	Result IsvOrgShortcutStatusDTO `json:"result,omitempty"`
}
type IsvOrgShortcutStatusDTO struct {
	BizNo string `json:"biz_no,omitempty"`
}
