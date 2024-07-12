package main

import (
	"log"
	"net"
	"reservation-service/config"
	pb "reservation-service/generated/reservation_service"
	"reservation-service/logs"
	"reservation-service/service"
	"reservation-service/storage/postgres"
	"reservation-service/storage/redis"

	"google.golang.org/grpc"
)

func main() {
	logs.InitLogger()
	logs.Logger.Info("Starting the server")
	db, err := postgres.ConnectDB()
	if err != nil {
		logs.Logger.Error("Failed connect to Data Base","error",err.Error())
		panic(err)
	}
	defer db.Close()
	r := redis.ConnectR()

	config := config.Load()

	listener, err := net.Listen("tcp", config.GRPC_PORT)
	if err != nil{
		logs.Logger.Error("Failed listen","error",err.Error())
		panic(err)
	}
	s := service.NewRRestaurantService(*postgres.NewRRestaurantRepo(db, r))
	server := grpc.NewServer()
	pb.RegisterReservationServiceServer(server,s)

	logs.Logger.Info("Server is Running","PORT",config.GRPC_PORT)
	log.Println("Server is running on ", listener.Addr())
	if err = server.Serve(listener);err != nil{
		logs.Logger.Error("Failed server is running","error",err.Error())
		panic(err)
	}
}
