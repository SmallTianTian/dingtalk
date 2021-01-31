package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessInstanceCommentAddRequest() *OapiProcessInstanceCommentAddRequest {
	return &OapiProcessInstanceCommentAddRequest{}
}

type OapiProcessInstanceCommentAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessInstanceCommentAddResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessInstanceCommentAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessInstanceCommentAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessInstanceCommentAddRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessInstanceCommentAddRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessInstanceCommentAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.instance.comment.add"
}
func (this *OapiProcessInstanceCommentAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessInstanceCommentAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessInstanceCommentAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessInstanceCommentAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessInstanceCommentAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessInstanceCommentAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessInstanceCommentAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessInstanceCommentAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Attachment struct {
	FileId   string `json:"file_id,omitempty"`
	FileName string `json:"file_name,omitempty"`
	FileSize string `json:"file_size,omitempty"`
	FileType string `json:"file_type,omitempty"`
	SpaceId  string `json:"space_id,omitempty"`
}
type AddCommentRequest struct {
	CommentUserid     string `json:"comment_userid,omitempty"`
	File              File   `json:"file,omitempty"`
	ProcessInstanceId string `json:"process_instance_id,omitempty"`
	Text              string `json:"text,omitempty"`
}
type OapiProcessInstanceCommentAddResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
