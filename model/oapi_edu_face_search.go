package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduFaceSearchRequest() *OapiEduFaceSearchRequest {
	return &OapiEduFaceSearchRequest{}
}

type OapiEduFaceSearchRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduFaceSearchResponse
	ClassId         int64
	Height          int64
	Synchronous     bool
	TopHttpMethod   string
	TopResponseType string
	Url             string
	Userid          string
	Width           int64
}

func (this *OapiEduFaceSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduFaceSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduFaceSearchRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduFaceSearchRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduFaceSearchRequest) SetHeight(height2 int64) {
	this.Height = height2
}
func (this *OapiEduFaceSearchRequest) GetHeight() int64 {
	return this.Height
}
func (this *OapiEduFaceSearchRequest) SetSynchronous(synchronous2 bool) {
	this.Synchronous = synchronous2
}
func (this *OapiEduFaceSearchRequest) GetSynchronous() bool {
	return this.Synchronous
}
func (this *OapiEduFaceSearchRequest) SetUrl(url2 string) {
	this.Url = url2
}
func (this *OapiEduFaceSearchRequest) GetUrl() string {
	return this.Url
}
func (this *OapiEduFaceSearchRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduFaceSearchRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduFaceSearchRequest) SetWidth(width2 int64) {
	this.Width = width2
}
func (this *OapiEduFaceSearchRequest) GetWidth() int64 {
	return this.Width
}
func (this *OapiEduFaceSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.face.search"
}
func (this *OapiEduFaceSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduFaceSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduFaceSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduFaceSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduFaceSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["height"] = this.Height
	txtParams["synchronous"] = this.Synchronous
	txtParams["url"] = this.Url
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["width"] = this.Width
	return txtParams
}
func (this *OapiEduFaceSearchRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduFaceSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduFaceSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduFaceSearchResponse struct {
	taobao.TaobaoResponse
	Result  TopSubmitFaceSearchResponse `json:"result,omitempty"`
	Success bool                        `json:"success,omitempty"`
}
type Faces struct {
	FaceType string `json:"face_type,omitempty"`
	Score    string `json:"score,omitempty"`
	TagId    string `json:"tag_id,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
type TopSubmitFaceSearchResponse struct {
	Faces  []Faces `json:"faces,omitempty"`
	TaskId string  `json:"task_id,omitempty"`
}
