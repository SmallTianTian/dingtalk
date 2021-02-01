package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserListidRequest() *OapiUserListidRequest {
	return &OapiUserListidRequest{}
}

type OapiUserListidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserListidResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserListidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserListidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserListidRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiUserListidRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiUserListidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.listid"
}
func (this *OapiUserListidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserListidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserListidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserListidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserListidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	return txtParams
}
func (this *OapiUserListidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserListidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserListidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserListidResponse struct {
	taobao.TaobaoResponse
	Result ListUserByDeptResponse `json:"result,omitempty"`
}
type ListUserByDeptResponse struct {
	UseridList []string `json:"userid_list,omitempty"`
}
