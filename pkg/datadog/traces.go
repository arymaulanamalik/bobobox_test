package datadog

import (
	"fmt"
	"net"

	ddtrace "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	ddprofiler "gopkg.in/DataDog/dd-trace-go.v1/profiler"

	"github.com/arymaulanamalik/bobobox_test/config"
	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
)

func init() {

	logger.Info("Init Tracer Datadog.", nil)
	logger.Info("Init Profiler run", nil)
	logger.Info("Version Code :"+config.DD_VERSION, nil)

	addr := net.JoinHostPort(config.DD_AGENT_HOST, config.DD_AGENT_PORT)
	serviceName := fmt.Sprintf("%s-%s-%s", config.DOMAIN, config.SERVICE_NAME, config.TYPE)

	if config.ENABLE_DD_PROFILER {
		ddprofiler.Start(
			ddprofiler.WithVersion(config.DD_VERSION),
			ddprofiler.WithAgentAddr(addr),
			ddprofiler.WithService(serviceName),
			ddprofiler.WithEnv(config.NAMESPACE),
			ddprofiler.WithTags("namespace:"+config.NAMESPACE+",domain:"+config.DOMAIN+"serviceName:"+config.SERVICE_NAME+",type:"+config.TYPE),
		)
	}

	if config.ENABLE_DD_TRACE {
		ddtrace.Start(
			ddtrace.WithAgentAddr(addr),
			ddtrace.WithAnalytics(true),
			ddtrace.WithDebugMode(config.DD_AGENT_DEBUG),
			ddtrace.WithEnv(config.NAMESPACE),
			ddtrace.WithGlobalTag("namespace", config.NAMESPACE),
			ddtrace.WithGlobalTag("domain", config.DOMAIN),
			ddtrace.WithGlobalTag("serviceName", config.SERVICE_NAME),
			ddtrace.WithGlobalTag("type", config.TYPE),
			ddtrace.WithService(serviceName),
			ddtrace.WithServiceVersion(config.DD_VERSION),
		)
	}

}
