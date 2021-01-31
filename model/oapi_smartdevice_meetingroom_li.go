package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceMeetingroomListRequest() *OapiSmartdeviceMeetingroomListRequest {
	return &OapiSmartdeviceMeetingroomListRequest{}
}

type OapiSmartdeviceMeetingroomListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceMeetingroomListResponse
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceMeetingroomListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceMeetingroomListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiSmartdeviceMeetingroomListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.meetingroom.list"
}
func (this *OapiSmartdeviceMeetingroomListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceMeetingroomListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceMeetingroomListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartdeviceMeetingroomListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Size, "size"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceMeetingroomListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceMeetingroomListResponse struct {
	taobao.TaobaoResponse
	Result OpenPageResult `json:"result,omitempty"`
}
type MeetingRoomOpenInfo struct {
	Capacity string  `json:"capacity,omitempty"`
	Labels   []int64 `json:"labels,omitempty"`
	Room     string  `json:"room,omitempty"`
	RoomId   string  `json:"room_id,omitempty"`
}
type OpenPageResult struct {
	HasMore    bool                  `json:"has_more,omitempty"`
	NextCursor string                `json:"next_cursor,omitempty"`
	Rooms      []MeetingRoomOpenInfo `json:"rooms,omitempty"`
}
