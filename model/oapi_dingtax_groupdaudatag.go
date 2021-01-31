package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingtaxGroupdaudataGetRequest() *OapiDingtaxGroupdaudataGetRequest {
	return &OapiDingtaxGroupdaudataGetRequest{}
}

type OapiDingtaxGroupdaudataGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiDingtaxGroupdaudataGetResponse
	OpenConversationIdList string
	StatDate               string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiDingtaxGroupdaudataGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingtaxGroupdaudataGetRequest) SetOpenConversationIdList(openConversationIdList2 string) {
	this.OpenConversationIdList = openConversationIdList2
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetOpenConversationIdList() string {
	return this.OpenConversationIdList
}
func (this *OapiDingtaxGroupdaudataGetRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingtax.groupdaudata.get"
}
func (this *OapiDingtaxGroupdaudataGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingtaxGroupdaudataGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingtaxGroupdaudataGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_conversation_id_list"] = this.OpenConversationIdList
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiDingtaxGroupdaudataGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationIdList, "openConversationIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingtaxGroupdaudataGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingtaxGroupdaudataGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                   `json:"errcode,omitempty"`
	Errmsg  string                  `json:"errmsg,omitempty"`
	Result  []DingtaxGroupDauRecord `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type DingtaxGroupDauRecord struct {
	OpenConversationId string `json:"open_conversation_id,omitempty"`
	StatDate           string `json:"stat_date,omitempty"`
	UnionId            string `json:"union_id,omitempty"`
}
