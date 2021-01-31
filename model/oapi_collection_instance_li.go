package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiCollectionInstanceListRequest() *OapiCollectionInstanceListRequest {
	return &OapiCollectionInstanceListRequest{}
}

type OapiCollectionInstanceListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionInstanceListResponse
	ActionDate      string
	BizType         int64
	FormCode        string
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionInstanceListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionInstanceListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionInstanceListRequest) SetActionDate(actionDate2 string) {
	this.ActionDate = actionDate2
}
func (this *OapiCollectionInstanceListRequest) GetActionDate() string {
	return this.ActionDate
}
func (this *OapiCollectionInstanceListRequest) SetBizType(bizType2 int64) {
	this.BizType = bizType2
}
func (this *OapiCollectionInstanceListRequest) GetBizType() int64 {
	return this.BizType
}
func (this *OapiCollectionInstanceListRequest) SetFormCode(formCode2 string) {
	this.FormCode = formCode2
}
func (this *OapiCollectionInstanceListRequest) GetFormCode() string {
	return this.FormCode
}
func (this *OapiCollectionInstanceListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiCollectionInstanceListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiCollectionInstanceListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiCollectionInstanceListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiCollectionInstanceListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.instance.list"
}
func (this *OapiCollectionInstanceListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionInstanceListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionInstanceListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionInstanceListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionInstanceListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["action_date"] = this.ActionDate
	txtParams["biz_type"] = this.BizType
	txtParams["form_code"] = this.FormCode
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiCollectionInstanceListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FormCode, "formCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCollectionInstanceListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionInstanceListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCollectionInstanceListResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
type FormInstanceResponse struct {
	CreateTime        time.Time  `json:"create_time,omitempty"`
	FormInstanceId    string     `json:"form_instance_id,omitempty"`
	Forms             []FormData `json:"forms,omitempty"`
	ModifyTime        time.Time  `json:"modify_time,omitempty"`
	StudentClassId    int64      `json:"student_class_id,omitempty"`
	StudentClassName  string     `json:"student_class_name,omitempty"`
	StudentName       string     `json:"student_name,omitempty"`
	StudentUserId     string     `json:"student_user_id,omitempty"`
	SubmitterUserName string     `json:"submitter_user_name,omitempty"`
	SubmitterUserid   string     `json:"submitter_userid,omitempty"`
}
