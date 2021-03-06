// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/webmin7761/go-school/homework/final/internal/biz"
	"github.com/webmin7761/go-school/homework/final/internal/cache/redis"
	"github.com/webmin7761/go-school/homework/final/internal/conf"
	"github.com/webmin7761/go-school/homework/final/internal/data"
	"github.com/webmin7761/go-school/homework/final/internal/mq/kafka"
	"github.com/webmin7761/go-school/homework/final/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confData *conf.Data, cache *conf.Cache, messageQueue *conf.MessageQueue, logger log.Logger) (func() error, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	fareRepo := data.NewFareRepo(dataData, logger)
	fareUsecase := biz.NewFareUsecase(fareRepo, logger)
	radixRC3, err := redis.NewRedisClient(cache)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	kafkaClient := kafka.NewKafkaClient(messageQueue)
	jobService := service.NewJobService(fareUsecase, radixRC3, kafkaClient, logger)
	v := newApp(jobService)
	return v, func() {
		cleanup()
	}, nil
}
