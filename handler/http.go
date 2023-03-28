package handler

import "sesi_8/service"

type HttpServer struct {
	service service.ServiceInterface
}

func NewHttpServer(service service.ServiceInterface) HttpServer {
	return HttpServer{service: service}
}
