package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkrecordUpdateRequest() *OapiWorkrecordUpdateRequest {
	return &OapiWorkrecordUpdateRequest{}
}

type OapiWorkrecordUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkrecordUpdateResponse
	RecordId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiWorkrecordUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkrecordUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkrecordUpdateRequest) SetRecordId(recordId2 string) {
	this.RecordId = recordId2
}
func (this *OapiWorkrecordUpdateRequest) GetRecordId() string {
	return this.RecordId
}
func (this *OapiWorkrecordUpdateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkrecordUpdateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkrecordUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workrecord.update"
}
func (this *OapiWorkrecordUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkrecordUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkrecordUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkrecordUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkrecordUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["record_id"] = this.RecordId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkrecordUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RecordId, "recordId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkrecordUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkrecordUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkrecordUpdateResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
