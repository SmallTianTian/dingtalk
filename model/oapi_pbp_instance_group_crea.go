package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiPbpInstanceGroupCreateRequest() *OapiPbpInstanceGroupCreateRequest {
	return &OapiPbpInstanceGroupCreateRequest{}
}

type OapiPbpInstanceGroupCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceGroupCreateResponse
	GroupParam      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceGroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceGroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceGroupCreateRequest) SetGroupParam(groupParam2 string) {
	this.GroupParam = groupParam2
}
func (this *OapiPbpInstanceGroupCreateRequest) GetGroupParam() string {
	return this.GroupParam
}
func (this *OapiPbpInstanceGroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.group.create"
}
func (this *OapiPbpInstanceGroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceGroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceGroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceGroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceGroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_param"] = this.GroupParam
	return txtParams
}
func (this *OapiPbpInstanceGroupCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiPbpInstanceGroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceGroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PunchGroupCreateParam struct {
	BizId     string `json:"biz_id,omitempty"`
	BizInstId string `json:"biz_inst_id,omitempty"`
}
type OapiPbpInstanceGroupCreateResponse struct {
	taobao.TaobaoResponse

	PunchGroupId string `json:"punch_group_id,omitempty"`
}
