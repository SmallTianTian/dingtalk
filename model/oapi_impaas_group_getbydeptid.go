package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImpaasGroupGetbydeptidRequest() *OapiImpaasGroupGetbydeptidRequest {
	return &OapiImpaasGroupGetbydeptidRequest{}
}

type OapiImpaasGroupGetbydeptidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupGetbydeptidResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupGetbydeptidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupGetbydeptidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupGetbydeptidRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiImpaasGroupGetbydeptidRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiImpaasGroupGetbydeptidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.group.getbydeptid"
}
func (this *OapiImpaasGroupGetbydeptidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupGetbydeptidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupGetbydeptidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupGetbydeptidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupGetbydeptidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	return txtParams
}
func (this *OapiImpaasGroupGetbydeptidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImpaasGroupGetbydeptidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupGetbydeptidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImpaasGroupGetbydeptidResponse struct {
	taobao.TaobaoResponse
	Result  BaseGroupInfo `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
type BaseGroupInfo struct {
	ConversationId     string `json:"conversation_id,omitempty"`
	Icon               string `json:"icon,omitempty"`
	OpenConversationId string `json:"open_conversation_id,omitempty"`
	Owner              int64  `json:"owner,omitempty"`
	Tag                int64  `json:"tag,omitempty"`
	Title              string `json:"title,omitempty"`
}
