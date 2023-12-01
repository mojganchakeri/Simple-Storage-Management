package configs

import (
	"store_service/internal/models"
)

const (
	SwaggerEnable = true
	StoreTable    = "storage"
	TagTable      = "tag"
	StoreTagTable = "store_tag"
	ServiceName   = "storeService"
	SecretKey     = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

var Env models.Env
