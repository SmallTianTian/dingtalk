package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessWorkrecordTaskQueryRequest() *OapiProcessWorkrecordTaskQueryRequest {
	return &OapiProcessWorkrecordTaskQueryRequest{}
}

type OapiProcessWorkrecordTaskQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordTaskQueryResponse
	Count           int64
	Offset          int64
	Status          int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiProcessWorkrecordTaskQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordTaskQueryRequest) SetCount(count2 int64) {
	this.Count = count2
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetCount() int64 {
	return this.Count
}
func (this *OapiProcessWorkrecordTaskQueryRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiProcessWorkrecordTaskQueryRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiProcessWorkrecordTaskQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.task.query"
}
func (this *OapiProcessWorkrecordTaskQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordTaskQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordTaskQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["count"] = this.Count
	txtParams["offset"] = this.Offset
	txtParams["status"] = this.Status
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProcessWorkrecordTaskQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Count, "count"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordTaskQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessWorkrecordTaskQueryResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
type FormItemVo struct {
	Content string `json:"content,omitempty"`
	Title   string `json:"title,omitempty"`
}
type WorkRecordVo struct {
	Forms      []FormItemVo `json:"forms,omitempty"`
	InstanceId string       `json:"instance_id,omitempty"`
	TaskId     string       `json:"task_id,omitempty"`
	Title      string       `json:"title,omitempty"`
	Url        string       `json:"url,omitempty"`
}
