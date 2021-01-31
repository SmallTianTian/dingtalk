package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	. "github.com/SmallTianTian/dingtalk/utils"
)

type DefaultDingTalkClient struct {
	URL              string
	needCheckRequest bool
	timeout          time.Duration
}

func NewDefaultDingTalkClient(url string) *DefaultDingTalkClient {
	return &DefaultDingTalkClient{
		URL:     url,
		timeout: 3 * time.Second,
	}
}

func (client *DefaultDingTalkClient) Execute3(req taobao.TaobaoRequest, accessKey, accessSecret string) error {
	return client.Execute5(req, accessKey, accessSecret, "", "")
}

func (client *DefaultDingTalkClient) Execute4(req taobao.TaobaoRequest, accessKey, accessSecret, suiteTicket string) error {
	return client.Execute5(req, accessKey, accessSecret, suiteTicket, "")
}

func (client *DefaultDingTalkClient) Execute5(req taobao.TaobaoRequest, accessKey, accessSecret, suiteTicket, corpId string) error {
	if IsEmpty(req.GetTopApiCallType()) || req.GetTopApiCallType() == "top" {
		// super.execute(request, null) : (T) ;
	}
	return client.executeOApi(req, "", accessKey, accessSecret, suiteTicket, corpId)
}

func (client *DefaultDingTalkClient) ExecuteSession(req taobao.TaobaoRequest, session string) error {
	if IsEmpty(req.GetTopApiCallType()) || req.GetTopApiCallType() == "top" {
		// super.execute(request, session)
	}
	return client.executeOApiSession(req, session)
}

func (client *DefaultDingTalkClient) Execute(req taobao.TaobaoRequest, session, accessKey, accessSecret, suiteTicket, corpId string) error {
	return nil
}

func (client *DefaultDingTalkClient) executeOApiSession(req taobao.TaobaoRequest, session string) error {
	return client.executeOApi(req, session, "", "", "", "")
}

func (client *DefaultDingTalkClient) executeOApi(req taobao.TaobaoRequest, session, accessKey, accessSecret, suiteTicket, corpId string) error {
	var fullUrl string
	if client.needCheckRequest {
		if err := req.Check(); err != nil {
			// TODO
		}
	}
	appParamsI := req.GetTextParams()
	appParams := make(map[string]string, len(appParamsI))
	for k, v := range appParamsI {
		bs, err := json.Marshal(v)
		if err != nil {
			return err
		}
		appParams[k] = string(bs)
	}

	if accessKey != "" {
		timestamp := CurrentTimeMillis()
		signature, err := ComputeSignature(accessSecret, GetCanonicalStringForIsv(timestamp, suiteTicket))
		if err != nil {
			return err
		}
		ps := make(map[string]string)
		ps["accessKey"] = accessKey
		ps["signature"] = signature
		ps[taobao.TIMESTAMP] = Int64S(timestamp)

		if !IsEmpty(suiteTicket) {
			ps["suiteTicket"] = suiteTicket
		}

		if !IsEmpty(corpId) {
			ps["corpId"] = corpId
		}

		queryStr := ParamToQueryString(ps)

		connect := "&"
		if strings.IndexByte(client.URL, '?') == -1 {
			connect = "?"
		}
		fullUrl = client.URL + connect + queryStr
	}

	var data *taobao.HttpResponse
	var err error
	if "GET" == req.GetApiMethodName() {
		data, err = taobao.DoGet(fullUrl, appParams, client.timeout)
	} else if uq, ok := req.(taobao.TaobaoUploadRequest); ok {
		fmt.Println(uq)
		// TODO
		// data = WebV2Utils.doPost(fullUrl, appParams, TaobaoUtils.cleanupMap(((TaobaoUploadRequest) request).getFileParams()), "UTF-8", this.connectTimeout, this.readTimeout, request.getHeaderMap());
	} else {
		jsonParams := make(map[string]interface{})
		for k, v := range appParams {
			var vi interface{}
			switch true {
			case strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]"):
				fallthrough
			case strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}"):
				if err := json.Unmarshal([]byte(v), &vi); err != nil {
					// TODO
					return nil
				}
			default:
				vi = v
			}
			jsonParams[k] = v
		}
		data, err = taobao.DoPostWithJson(fullUrl, jsonParams, client.timeout)
	}
	if err != nil {
		return err
	}

	if set, ok := req.(taobao.SetResponse); ok {
		return set.SetResponse(req, data)
	} else {
		return errors.New("Not can set response.")
	}
}
