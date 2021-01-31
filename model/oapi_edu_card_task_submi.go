package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCardTaskSubmitRequest() *OapiEduCardTaskSubmitRequest {
	return &OapiEduCardTaskSubmitRequest{}
}

type OapiEduCardTaskSubmitRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCardTaskSubmitResponse
	CardType        string
	Content         string
	MeteringNumber  string
	TopHttpMethod   string
	TopResponseType string
	UserCardTaskId  int64
	Userid          string
}

func (this *OapiEduCardTaskSubmitRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCardTaskSubmitRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCardTaskSubmitRequest) SetCardType(cardType2 string) {
	this.CardType = cardType2
}
func (this *OapiEduCardTaskSubmitRequest) GetCardType() string {
	return this.CardType
}
func (this *OapiEduCardTaskSubmitRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiEduCardTaskSubmitRequest) GetContent() string {
	return this.Content
}
func (this *OapiEduCardTaskSubmitRequest) SetMeteringNumber(meteringNumber2 string) {
	this.MeteringNumber = meteringNumber2
}
func (this *OapiEduCardTaskSubmitRequest) GetMeteringNumber() string {
	return this.MeteringNumber
}
func (this *OapiEduCardTaskSubmitRequest) SetUserCardTaskId(userCardTaskId2 int64) {
	this.UserCardTaskId = userCardTaskId2
}
func (this *OapiEduCardTaskSubmitRequest) GetUserCardTaskId() int64 {
	return this.UserCardTaskId
}
func (this *OapiEduCardTaskSubmitRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduCardTaskSubmitRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduCardTaskSubmitRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.card.task.submit"
}
func (this *OapiEduCardTaskSubmitRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCardTaskSubmitRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCardTaskSubmitRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCardTaskSubmitRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCardTaskSubmitRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["card_type"] = this.CardType
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["metering_number"] = this.MeteringNumber
	txtParams["user_card_task_id"] = this.UserCardTaskId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduCardTaskSubmitRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCardTaskSubmitRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCardTaskSubmitRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCardTaskSubmitResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
