package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkbenchShortcutListbypagingRequest() *OapiWorkbenchShortcutListbypagingRequest {
	return &OapiWorkbenchShortcutListbypagingRequest{}
}

type OapiWorkbenchShortcutListbypagingRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkbenchShortcutListbypagingResponse
	PageIndex       int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkbenchShortcutListbypagingRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkbenchShortcutListbypagingRequest) SetPageIndex(pageIndex2 int64) {
	this.PageIndex = pageIndex2
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetPageIndex() int64 {
	return this.PageIndex
}
func (this *OapiWorkbenchShortcutListbypagingRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workbench.shortcut.listbypaging"
}
func (this *OapiWorkbenchShortcutListbypagingRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkbenchShortcutListbypagingRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkbenchShortcutListbypagingRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["page_index"] = this.PageIndex
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiWorkbenchShortcutListbypagingRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkbenchShortcutListbypagingRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkbenchShortcutListbypagingResponse struct {
	taobao.TaobaoResponse
	Result IsvOrgShortcutListDTO `json:"result,omitempty"`
}
