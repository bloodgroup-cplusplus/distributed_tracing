package playlistsapi

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

const serviceName = "playlists-api"

var environment = os.GetEnv("ENVIRONMENT")
var redis_host = os.GetEnv("REDIS HOST")
var redis_port = os.GetEnv("REDIS_PORT")
var ctx = context.Background()
var rbd *redis.Client

func main() {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		// "const " sampler is a binary sampling strategy: 0-nvever sample, 1-always sample
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		// log the emitted spans to stdout
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "jaeger:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init jaeger: %v\n", err))
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.))
}
