package datastruct

import (
	"time"

	"api/app/common/config"
)

type HealthCheck struct {
	AppStart time.Time     `json:"app_start"`
	Config   config.Config `json:"config"`
	Db       bool          `json:"db"`
}
