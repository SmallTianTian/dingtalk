package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiIndustryUserGetRequest() *OapiIndustryUserGetRequest {
	return &OapiIndustryUserGetRequest{}
}

type OapiIndustryUserGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIndustryUserGetResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiIndustryUserGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIndustryUserGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIndustryUserGetRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiIndustryUserGetRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiIndustryUserGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiIndustryUserGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiIndustryUserGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.industry.user.get"
}
func (this *OapiIndustryUserGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIndustryUserGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIndustryUserGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIndustryUserGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIndustryUserGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiIndustryUserGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiIndustryUserGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIndustryUserGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiIndustryUserGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenIndustryEmp `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type OpenIndustryEmp struct {
	Feature string     `json:"feature,omitempty"`
	Name    string     `json:"name,omitempty"`
	Roles   []OpenRole `json:"roles,omitempty"`
	Unionid string     `json:"unionid,omitempty"`
}
