package taobao

import "os"

type TaobaoUploadRequest interface {
	TaobaoRequest
	GetFileParams() map[string]os.File
}
