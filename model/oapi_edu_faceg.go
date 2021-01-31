package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduFaceGetRequest() *OapiEduFaceGetRequest {
	return &OapiEduFaceGetRequest{}
}

type OapiEduFaceGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduFaceGetResponse
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduFaceGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduFaceGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduFaceGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduFaceGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduFaceGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduFaceGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduFaceGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.face.get"
}
func (this *OapiEduFaceGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduFaceGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduFaceGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduFaceGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduFaceGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduFaceGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduFaceGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduFaceGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduFaceGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                  `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Result  TopQueryFaceIdResponse `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type TopQueryFaceIdResponse struct {
	HasRecordFace bool `json:"has_record_face,omitempty"`
}
