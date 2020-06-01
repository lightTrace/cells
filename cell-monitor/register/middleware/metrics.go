package middleware

import (
	"cells/cell-monitor/register/userservice"
	"github.com/go-kit/kit/metrics"
	"time"
)

// metricMiddleware 定义监控中间件，嵌入IUserService
// 新增监控指标项：requestCount和requestLatency
type metricMiddleware struct {
	userservice.IUserService
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// Metrics 指标采集方法
func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) userservice.ServiceMiddleware {
	return func(next userservice.IUserService) userservice.IUserService {
		return metricMiddleware{
			next,
			requestCount,
			requestLatency}
	}
}

func (mm metricMiddleware) GetUserName(userId string) (string, error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetUserName"}
		mm.requestCount.With(lvs...).Add(1)
		mm.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mm.IUserService.GetUserName(userId)
}

func (mm metricMiddleware) UpdateUserName(userId, userName string) error {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateUserName"}
		mm.requestCount.With(lvs...).Add(1)
		mm.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mm.IUserService.UpdateUserName(userId, userName)
}
