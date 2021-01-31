package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmRosterMetaGetRequest() *OapiSmartworkHrmRosterMetaGetRequest {
	return &OapiSmartworkHrmRosterMetaGetRequest{}
}

type OapiSmartworkHrmRosterMetaGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmRosterMetaGetResponse
	Agentid         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmRosterMetaGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.roster.meta.get"
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	return txtParams
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmRosterMetaGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmRosterMetaGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64           `json:"errcode,omitempty"`
	Errmsg  string          `json:"errmsg,omitempty"`
	Result  []GroupMetaInfo `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
