package model

import (
	"os"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiFileUploadSingleRequest() *OapiFileUploadSingleRequest {
	return &OapiFileUploadSingleRequest{}
}

type OapiFileUploadSingleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFileUploadSingleResponse
	AgentId         string
	File            os.File
	FileSize        int64
	TopResponseType string
}

func (this *OapiFileUploadSingleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFileUploadSingleRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiFileUploadSingleRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiFileUploadSingleRequest) SetFile(file2 os.File) {
	this.File = file2
}
func (this *OapiFileUploadSingleRequest) GetFile() os.File {
	return this.File
}
func (this *OapiFileUploadSingleRequest) SetFileSize(fileSize2 int64) {
	this.FileSize = fileSize2
}
func (this *OapiFileUploadSingleRequest) GetFileSize() int64 {
	return this.FileSize
}
func (this *OapiFileUploadSingleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.file.upload.single"
}
func (this *OapiFileUploadSingleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFileUploadSingleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFileUploadSingleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFileUploadSingleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["file_size"] = this.FileSize
	return txtParams
}
func (this *OapiFileUploadSingleRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiFileUploadSingleRequest) GetFileParams() map[string]os.File {
	params := make(map[string]os.File)
	params["file"] = this.File
	return params
}
func (this *OapiFileUploadSingleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFileUploadSingleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFileUploadSingleResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}
