package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiFileUploadTransactionRequest() *OapiFileUploadTransactionRequest {
	return &OapiFileUploadTransactionRequest{}
}

type OapiFileUploadTransactionRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFileUploadTransactionResponse
	AgentId         string
	ChunkNumbers    int64
	FileSize        int64
	TopHttpMethod   string
	TopResponseType string
	UploadId        string
}

func (this *OapiFileUploadTransactionRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiFileUploadTransactionRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFileUploadTransactionRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiFileUploadTransactionRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiFileUploadTransactionRequest) SetChunkNumbers(chunkNumbers2 int64) {
	this.ChunkNumbers = chunkNumbers2
}
func (this *OapiFileUploadTransactionRequest) GetChunkNumbers() int64 {
	return this.ChunkNumbers
}
func (this *OapiFileUploadTransactionRequest) SetFileSize(fileSize2 int64) {
	this.FileSize = fileSize2
}
func (this *OapiFileUploadTransactionRequest) GetFileSize() int64 {
	return this.FileSize
}
func (this *OapiFileUploadTransactionRequest) SetUploadId(uploadId2 string) {
	this.UploadId = uploadId2
}
func (this *OapiFileUploadTransactionRequest) GetUploadId() string {
	return this.UploadId
}
func (this *OapiFileUploadTransactionRequest) GetApiMethodName() string {
	return "dingtalk.oapi.file.upload.transaction"
}
func (this *OapiFileUploadTransactionRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFileUploadTransactionRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFileUploadTransactionRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFileUploadTransactionRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFileUploadTransactionRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["chunk_numbers"] = this.ChunkNumbers
	txtParams["file_size"] = this.FileSize
	txtParams["upload_id"] = this.UploadId
	return txtParams
}
func (this *OapiFileUploadTransactionRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiFileUploadTransactionRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFileUploadTransactionRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFileUploadTransactionResponse struct {
	taobao.TaobaoResponse

	MediaId  string `json:"media_id,omitempty"`
	UploadId string `json:"upload_id,omitempty"`
}
