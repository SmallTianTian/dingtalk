package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappListByUseridRequest() *OapiMicroappListByUseridRequest {
	return &OapiMicroappListByUseridRequest{}
}

type OapiMicroappListByUseridRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappListByUseridResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiMicroappListByUseridRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiMicroappListByUseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappListByUseridRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiMicroappListByUseridRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiMicroappListByUseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.list_by_userid"
}
func (this *OapiMicroappListByUseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappListByUseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappListByUseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappListByUseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappListByUseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiMicroappListByUseridRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappListByUseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappListByUseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappListByUseridResponse struct {
	taobao.TaobaoResponse
	AppList []Applist `json:"appList,omitempty"`
}
type Applist struct {
	AgentId        int64  `json:"agentId,omitempty"`
	AppDesc        string `json:"appDesc,omitempty"`
	AppIcon        string `json:"appIcon,omitempty"`
	AppStatus      int64  `json:"appStatus,omitempty"`
	HomepageLink   string `json:"homepageLink,omitempty"`
	IsSelf         bool   `json:"isSelf,omitempty"`
	Name           string `json:"name,omitempty"`
	OmpLink        string `json:"ompLink,omitempty"`
	PcHomepageLink string `json:"pcHomepageLink,omitempty"`
}
