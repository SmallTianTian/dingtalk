package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappListbypageRequest() *OapiMicroappListbypageRequest {
	return &OapiMicroappListbypageRequest{}
}

type OapiMicroappListbypageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappListbypageResponse
	AgentId         int64
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappListbypageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappListbypageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappListbypageRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappListbypageRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappListbypageRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiMicroappListbypageRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiMicroappListbypageRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiMicroappListbypageRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiMicroappListbypageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.listbypage"
}
func (this *OapiMicroappListbypageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappListbypageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappListbypageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappListbypageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappListbypageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiMicroappListbypageRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappListbypageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappListbypageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappListbypageResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
}
