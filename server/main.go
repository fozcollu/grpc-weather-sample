package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-weather-sample/api"
	"math/rand"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("Server starting...")
	panic(srv.Serve(lis))
}

type myWeatherService struct {
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, request *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "tr_izm", CityName: "Izmir"},
			&api.CityEntry{CityCode: "tr_ist", CityName: "Istanbul"},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(request *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&api.WeatherResponse{
			Temperature: rand.Float32()*10 + 10,
		})

		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}

	return nil
}
