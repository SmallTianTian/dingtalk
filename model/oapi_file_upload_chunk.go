package model

import (
	"os"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiFileUploadChunkRequest() *OapiFileUploadChunkRequest {
	return &OapiFileUploadChunkRequest{}
}

type OapiFileUploadChunkRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFileUploadChunkResponse
	AgentId         string
	ChunkSequence   int64
	File            os.File
	TopResponseType string
	UploadId        string
}

func (this *OapiFileUploadChunkRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFileUploadChunkRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiFileUploadChunkRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiFileUploadChunkRequest) SetChunkSequence(chunkSequence2 int64) {
	this.ChunkSequence = chunkSequence2
}
func (this *OapiFileUploadChunkRequest) GetChunkSequence() int64 {
	return this.ChunkSequence
}
func (this *OapiFileUploadChunkRequest) SetFile(file2 os.File) {
	this.File = file2
}
func (this *OapiFileUploadChunkRequest) GetFile() os.File {
	return this.File
}
func (this *OapiFileUploadChunkRequest) SetUploadId(uploadId2 string) {
	this.UploadId = uploadId2
}
func (this *OapiFileUploadChunkRequest) GetUploadId() string {
	return this.UploadId
}
func (this *OapiFileUploadChunkRequest) GetApiMethodName() string {
	return "dingtalk.oapi.file.upload.chunk"
}
func (this *OapiFileUploadChunkRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFileUploadChunkRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFileUploadChunkRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFileUploadChunkRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["chunk_sequence"] = this.ChunkSequence
	txtParams["upload_id"] = this.UploadId
	return txtParams
}
func (this *OapiFileUploadChunkRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiFileUploadChunkRequest) GetFileParams() map[string]os.File {
	params := make(map[string]os.File)
	params["file"] = this.File
	return params
}
func (this *OapiFileUploadChunkRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFileUploadChunkRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFileUploadChunkResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
