package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduFamilyChildGetRequest() *OapiEduFamilyChildGetRequest {
	return &OapiEduFamilyChildGetRequest{}
}

type OapiEduFamilyChildGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduFamilyChildGetResponse
	ChildUserid     string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduFamilyChildGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduFamilyChildGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduFamilyChildGetRequest) SetChildUserid(childUserid2 string) {
	this.ChildUserid = childUserid2
}
func (this *OapiEduFamilyChildGetRequest) GetChildUserid() string {
	return this.ChildUserid
}
func (this *OapiEduFamilyChildGetRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduFamilyChildGetRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduFamilyChildGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.family.child.get"
}
func (this *OapiEduFamilyChildGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduFamilyChildGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduFamilyChildGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduFamilyChildGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduFamilyChildGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["child_userid"] = this.ChildUserid
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiEduFamilyChildGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChildUserid, "childUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduFamilyChildGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduFamilyChildGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduFamilyChildGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Result  ChildDto `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
type BindStudent struct {
	ClassId    string `json:"class_id,omitempty"`
	CorpId     string `json:"corp_id,omitempty"`
	PeriodCode string `json:"period_code,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
type ChildDto struct {
	Avatar       string        `json:"avatar,omitempty"`
	BindStudents []BindStudent `json:"bind_students,omitempty"`
	Nick         string        `json:"nick,omitempty"`
	OpenId       string        `json:"open_id,omitempty"`
	UnionId      string        `json:"unionId,omitempty"`
	Userid       string        `json:"userid,omitempty"`
}
