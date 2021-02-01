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

type DefaultTaobaoClient struct {
	serverUrl        string
	needCheckRequest bool
	version          string
	appKey           string
	timeout          time.Duration
}

func (dtc *DefaultTaobaoClient) TExecute(req taobao.TaobaoRequest) error {
	return dtc.TExecute2(req, "")
}

func (dtc *DefaultTaobaoClient) TExecute2(req taobao.TaobaoRequest, session string) error {
	start := CurrentTimeMillis()
	if dtc.needCheckRequest {
		if err := req.Check(); err != nil {
			// TODO
		}
	}
	return dtc.invokeApi(req, session, start)
}

func (dtc *DefaultTaobaoClient) invokeApi(req taobao.TaobaoRequest, session string, start int64) error {
	fullUrl := dtc.serverUrl
	appParamsI := req.GetTextParams()
	appParams := make(map[string]string, len(appParamsI))
	for k, v := range appParamsI {
		bs, err := json.Marshal(v)
		if err != nil {
			return err
		}
		appParams[k] = string(bs)
	}
	protocalMustParams := make(map[string]string)

	protocalMustParams[taobao.METHOD] = req.GetApiMethodName()
	if req.GetTopApiVersion() != "" {
		protocalMustParams[taobao.VERSION] = req.GetTopApiVersion()
	} else {
		protocalMustParams[taobao.VERSION] = dtc.version
	}
	if dtc.appKey != "" {
		protocalMustParams[taobao.APP_KEY] = dtc.appKey
	}
	if req.GetTimestamp() <= 0 {
		protocalMustParams[taobao.TIMESTAMP] = FormatTime(time.Now())
	} else {
		protocalMustParams[taobao.TIMESTAMP] = FormatTime(time.Unix(req.GetTimestamp()/int64(time.Millisecond), 0))
	}
	isJson := req.GetTopContentType() == "json"
	sysMustQuery := ParamToQueryString(protocalMustParams)
	appParamsQuery := ParamToQueryString(appParams)
	if sysMustQuery != "" {
		fullUrl += "?" + sysMustQuery
	}
	if isJson {
		connect := "&"
		if !strings.Contains(fullUrl, "?") {
			connect = "?"
		}
		fullUrl += connect + appParamsQuery
	}
	var data *taobao.HttpResponse
	var err error
	if uploadReq, ok := req.(taobao.TaobaoUploadRequest); ok {
		// TODO
		fmt.Println(uploadReq)
		// data = WebV2Utils.doPost(fullUrl, appParams, TaobaoUtils.cleanupMap(((TaobaoUploadRequest) request.getRequestBase()).getFileParams()), "UTF-8", this.connectTimeout, this.readTimeout, request.getHeaderMap());
	} else {
		if isJson {
			data, err = taobao.DoPostWithJson(fullUrl, req.GetTextParams(), dtc.timeout)
		} else {
			data, err = taobao.DoPost(fullUrl, dtc.timeout, appParams, req.GetHeaderMap())
		}
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
