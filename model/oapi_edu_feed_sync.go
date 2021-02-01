package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduFeedSyncRequest() *OapiEduFeedSyncRequest {
	return &OapiEduFeedSyncRequest{}
}

type OapiEduFeedSyncRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduFeedSyncResponse
	AlbumId         string
	DeptId          int64
	FeeType         int64
	FeedMedias      string
	Future          string
	MediaUid        string
	OpUserId        string
	SendTime        int64
	SendUid         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduFeedSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduFeedSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduFeedSyncRequest) SetAlbumId(albumId2 string) {
	this.AlbumId = albumId2
}
func (this *OapiEduFeedSyncRequest) GetAlbumId() string {
	return this.AlbumId
}
func (this *OapiEduFeedSyncRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiEduFeedSyncRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiEduFeedSyncRequest) SetFeeType(feeType2 int64) {
	this.FeeType = feeType2
}
func (this *OapiEduFeedSyncRequest) GetFeeType() int64 {
	return this.FeeType
}
func (this *OapiEduFeedSyncRequest) SetFeedMedias(feedMedias2 string) {
	this.FeedMedias = feedMedias2
}
func (this *OapiEduFeedSyncRequest) GetFeedMedias() string {
	return this.FeedMedias
}
func (this *OapiEduFeedSyncRequest) SetFuture(future2 string) {
	this.Future = future2
}
func (this *OapiEduFeedSyncRequest) GetFuture() string {
	return this.Future
}
func (this *OapiEduFeedSyncRequest) SetMediaUid(mediaUid2 string) {
	this.MediaUid = mediaUid2
}
func (this *OapiEduFeedSyncRequest) GetMediaUid() string {
	return this.MediaUid
}
func (this *OapiEduFeedSyncRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiEduFeedSyncRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiEduFeedSyncRequest) SetSendTime(sendTime2 int64) {
	this.SendTime = sendTime2
}
func (this *OapiEduFeedSyncRequest) GetSendTime() int64 {
	return this.SendTime
}
func (this *OapiEduFeedSyncRequest) SetSendUid(sendUid2 string) {
	this.SendUid = sendUid2
}
func (this *OapiEduFeedSyncRequest) GetSendUid() string {
	return this.SendUid
}
func (this *OapiEduFeedSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.feed.sync"
}
func (this *OapiEduFeedSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduFeedSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduFeedSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduFeedSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduFeedSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["album_id"] = this.AlbumId
	txtParams["dept_id"] = this.DeptId
	txtParams["fee_type"] = this.FeeType
	txtParams["feed_medias"] = this.FeedMedias
	txtParams["future"] = this.Future
	txtParams["media_uid"] = this.MediaUid
	txtParams["op_userId"] = this.OpUserId
	txtParams["send_time"] = this.SendTime
	txtParams["send_uid"] = this.SendUid
	return txtParams
}
func (this *OapiEduFeedSyncRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FeeType, "feeType"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduFeedSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduFeedSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type IndustrySyncFeedMediaReq struct {
	MediaType    int64  `json:"media_type,omitempty"`
	MediaUrl     string `json:"media_url,omitempty"`
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
}
type OapiEduFeedSyncResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
