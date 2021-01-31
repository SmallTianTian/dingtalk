package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiMedalCorpmedalQueryRequest() *OapiMedalCorpmedalQueryRequest {
	return &OapiMedalCorpmedalQueryRequest{}
}

type OapiMedalCorpmedalQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMedalCorpmedalQueryResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiMedalCorpmedalQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMedalCorpmedalQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMedalCorpmedalQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiMedalCorpmedalQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiMedalCorpmedalQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.medal.corpmedal.query"
}
func (this *OapiMedalCorpmedalQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMedalCorpmedalQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMedalCorpmedalQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMedalCorpmedalQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMedalCorpmedalQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiMedalCorpmedalQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiMedalCorpmedalQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMedalCorpmedalQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMedalCorpmedalQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64                  `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Result  CorpMedalQueryResponse `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type CorpMedalDTO struct {
	GrantTime  time.Time `json:"grant_time,omitempty"`
	TemplateId int64     `json:"template_id,omitempty"`
	Wear       bool      `json:"wear,omitempty"`
}
type CorpMedalQueryResponse struct {
	CorpMedalList []CorpMedalDTO `json:"corp_medal_list,omitempty"`
}
