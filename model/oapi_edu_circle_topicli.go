package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCircleTopiclistRequest() *OapiEduCircleTopiclistRequest {
	return &OapiEduCircleTopiclistRequest{}
}

type OapiEduCircleTopiclistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCircleTopiclistResponse
	BizType         int64
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduCircleTopiclistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCircleTopiclistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCircleTopiclistRequest) SetBizType(bizType2 int64) {
	this.BizType = bizType2
}
func (this *OapiEduCircleTopiclistRequest) GetBizType() int64 {
	return this.BizType
}
func (this *OapiEduCircleTopiclistRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduCircleTopiclistRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduCircleTopiclistRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduCircleTopiclistRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduCircleTopiclistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.circle.topiclist"
}
func (this *OapiEduCircleTopiclistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCircleTopiclistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCircleTopiclistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCircleTopiclistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCircleTopiclistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	txtParams["class_id"] = this.ClassId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduCircleTopiclistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCircleTopiclistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCircleTopiclistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCircleTopiclistResponse struct {
	taobao.TaobaoResponse
	Result  []OpenCircleTopicResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
