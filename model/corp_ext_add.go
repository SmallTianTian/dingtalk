package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpExtAddRequest() *CorpExtAddRequest {
	return &CorpExtAddRequest{}
}

type CorpExtAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtAddResponse
	Contact         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtAddRequest) SetContact(contact2 string) {
	this.Contact = contact2
}
func (this *CorpExtAddRequest) GetContact() string {
	return this.Contact
}
func (this *CorpExtAddRequest) GetApiMethodName() string {
	return "dingtalk.corp.ext.add"
}
func (this *CorpExtAddRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtAddRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact"] = this.Contact
	return txtParams
}
func (this *CorpExtAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpExtAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenExtContact struct {
	Address        string   `json:"address,omitempty"`
	CompanyName    string   `json:"company_name,omitempty"`
	FollowerUserid string   `json:"follower_userid,omitempty"`
	LabelIds       []int64  `json:"label_ids,omitempty"`
	Mobile         string   `json:"mobile,omitempty"`
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ShareDeptids   []int64  `json:"share_deptids,omitempty"`
	ShareUserids   []string `json:"share_userids,omitempty"`
	StateCode      string   `json:"state_code,omitempty"`
	Title          string   `json:"title,omitempty"`
}
type CorpExtAddResponse struct {
	taobao.TaobaoResponse
	Userid string `json:"userid,omitempty"`
}
