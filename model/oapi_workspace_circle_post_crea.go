package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceCirclePostCreateRequest() *OapiWorkspaceCirclePostCreateRequest {
	return &OapiWorkspaceCirclePostCreateRequest{}
}

type OapiWorkspaceCirclePostCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCirclePostCreateResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceCirclePostCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCirclePostCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCirclePostCreateRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiWorkspaceCirclePostCreateRequest) GetReq() string {
	return this.Req
}
func (this *OapiWorkspaceCirclePostCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.circle.post.create"
}
func (this *OapiWorkspaceCirclePostCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCirclePostCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCirclePostCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCirclePostCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCirclePostCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiWorkspaceCirclePostCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceCirclePostCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCirclePostCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FvDingPanFileContentOpenDto struct {
	MediaId string `json:"media_id,omitempty"`
	Name    string `json:"name,omitempty"`
}
type FvAttachmentsOpenDto struct {
	DingFiles []FvDingPanFileContentOpenDto `json:"ding_files,omitempty"`
}
type FvPhotoOpenDto struct {
	MediaId string `json:"media_id,omitempty"`
}
type FvPhotoContentOpenDto struct {
	Photos []FvPhotoOpenDto `json:"photos,omitempty"`
}
type FvVideoContentOpenDto struct {
	Bitrate      int64  `json:"bitrate,omitempty"`
	Duration     int64  `json:"duration,omitempty"`
	FileName     string `json:"file_name,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
	FileType     string `json:"file_type,omitempty"`
	Height       int64  `json:"height,omitempty"`
	PicMediaId   string `json:"pic_media_id,omitempty"`
	VideoMediaId string `json:"video_media_id,omitempty"`
	Width        int64  `json:"width,omitempty"`
}
type FvPostContentOpenDto struct {
	Attachments  FvAttachmentsOpenDto  `json:"attachments,omitempty"`
	ContentType  int64                 `json:"content_type,omitempty"`
	PhotoContent FvPhotoContentOpenDto `json:"photo_content,omitempty"`
	Text         string                `json:"text,omitempty"`
	VideoContent FvVideoContentOpenDto `json:"video_content,omitempty"`
}
type FvPostTagOpenDto struct {
	Name    string `json:"name,omitempty"`
	TagId   int64  `json:"tag_id,omitempty"`
	TagType int64  `json:"tag_type,omitempty"`
}
type FvPostCreateOpenDto struct {
	Content FvPostContentOpenDto `json:"content,omitempty"`
	Tags    []FvPostTagOpenDto   `json:"tags,omitempty"`
	Userid  string               `json:"userid,omitempty"`
	Uuid    string               `json:"uuid,omitempty"`
}
type OapiWorkspaceCirclePostCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
