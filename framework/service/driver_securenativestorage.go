package gnonative

import "github.com/gnolang/gnonative/service"

type SecureNativeStorage interface {
	service.SecureNativeStorage
	service.Batch
}
