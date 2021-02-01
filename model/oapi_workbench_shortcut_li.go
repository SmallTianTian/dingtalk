package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkbenchShortcutListRequest() *OapiWorkbenchShortcutListRequest {
	return &OapiWorkbenchShortcutListRequest{}
}

type OapiWorkbenchShortcutListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkbenchShortcutListResponse
	AppId           string
	PageIndex       int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkbenchShortcutListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkbenchShortcutListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkbenchShortcutListRequest) SetAppId(appId2 string) {
	this.AppId = appId2
}
func (this *OapiWorkbenchShortcutListRequest) GetAppId() string {
	return this.AppId
}
func (this *OapiWorkbenchShortcutListRequest) SetPageIndex(pageIndex2 int64) {
	this.PageIndex = pageIndex2
}
func (this *OapiWorkbenchShortcutListRequest) GetPageIndex() int64 {
	return this.PageIndex
}
func (this *OapiWorkbenchShortcutListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiWorkbenchShortcutListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiWorkbenchShortcutListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workbench.shortcut.list"
}
func (this *OapiWorkbenchShortcutListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkbenchShortcutListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkbenchShortcutListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkbenchShortcutListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkbenchShortcutListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams["page_index"] = this.PageIndex
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiWorkbenchShortcutListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppId, "appId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkbenchShortcutListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkbenchShortcutListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkbenchShortcutListResponse struct {
	taobao.TaobaoResponse
	Result IsvOrgShortcutListDTO `json:"result,omitempty"`
}
type IsvOrgShortcutDTO struct {
	BizNo         string `json:"biz_no,omitempty"`
	Icon          string `json:"icon,omitempty"`
	Name          string `json:"name,omitempty"`
	PcShortcutUri string `json:"pc_shortcut_uri,omitempty"`
	ShortcutUri   string `json:"shortcut_uri,omitempty"`
}
type IsvOrgShortcutListDTO struct {
	ShortcutList []IsvOrgShortcutDTO `json:"shortcut_list,omitempty"`
}
