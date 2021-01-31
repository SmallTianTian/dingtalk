package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduDeptGetRequest() *OapiEduDeptGetRequest {
	return &OapiEduDeptGetRequest{}
}

type OapiEduDeptGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduDeptGetResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduDeptGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduDeptGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduDeptGetRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiEduDeptGetRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiEduDeptGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.dept.get"
}
func (this *OapiEduDeptGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduDeptGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduDeptGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduDeptGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduDeptGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	return txtParams
}
func (this *OapiEduDeptGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduDeptGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduDeptGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduDeptGetResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type OpenEduDeptDetail struct {
	Chain       string `json:"chain,omitempty"`
	ContactType string `json:"contact_type,omitempty"`
	DeptId      int64  `json:"dept_id,omitempty"`
	DeptType    string `json:"dept_type,omitempty"`
	Feature     string `json:"feature,omitempty"`
	Name        string `json:"name,omitempty"`
	Nick        string `json:"nick,omitempty"`
}
