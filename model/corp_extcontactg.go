package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpExtcontactGetRequest() *CorpExtcontactGetRequest {
	return &CorpExtcontactGetRequest{}
}

type CorpExtcontactGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtcontactGetResponse
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *CorpExtcontactGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtcontactGetRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *CorpExtcontactGetRequest) GetUserId() string {
	return this.UserId
}
func (this *CorpExtcontactGetRequest) GetApiMethodName() string {
	return "dingtalk.corp.extcontact.get"
}
func (this *CorpExtcontactGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtcontactGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtcontactGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtcontactGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtcontactGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtcontactGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["user_id"] = this.UserId
	return txtParams
}
func (this *CorpExtcontactGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UserId, "userId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpExtcontactGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtcontactGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtcontactGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
