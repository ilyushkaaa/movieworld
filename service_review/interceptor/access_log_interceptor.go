package interceptor

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	errorreview "kinopoisk/service_review/error"
	"log"
	"time"
)

type loggerKey int

const MyLoggerKey loggerKey = 3

func AccessLogInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("access log interceptor start")
	logger, err := GetLoggerFromContext(ctx)
	if err != nil {
		log.Printf("can not get logger from context: %s", err)
		return nil, err
	}
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	logger.Infow("Request result",
		"after incoming call ", info.FullMethod,
		"request ", req,
		"reply ", reply,
		"time of call ", time.Since(start),
		"metadata ", md,
		"error ", err,
	)
	return reply, err
}

func GetLoggerFromContext(ctx context.Context) (*zap.SugaredLogger, error) {
	myLogger, ok := ctx.Value(MyLoggerKey).(*zap.SugaredLogger)
	if !ok {
		return myLogger, errorreview.ErrorNoLogger
	}
	return myLogger, nil
}
