package main

import (
	"applicationdesigntest/internal/app"
	"applicationdesigntest/internal/logger"
	"applicationdesigntest/internal/model"
	"applicationdesigntest/internal/repository"
	"applicationdesigntest/internal/service"
	"errors"
	"net/http"
	"os"
	"time"
)

func main() {
	orderRepository := &repository.InMemoryOrderRepository{}
	availabilityRespository := &repository.InMemoryRoomAvailabilityRepository{
		Availability: []model.RoomAvailability{
			{"reddison", "lux", date(2024, 1, 1), 1},
			{"reddison", "lux", date(2024, 1, 2), 1},
			{"reddison", "lux", date(2024, 1, 3), 1},
			{"reddison", "lux", date(2024, 1, 4), 1},
			{"reddison", "lux", date(2024, 1, 5), 0},
		},
	}

	service := service.NewBookingService(orderRepository, availabilityRespository)
	handler := app.NewBookingHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handler.CreateOrder)

	logger.Infof("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		logger.Infof("Server closed")
	} else if err != nil {
		logger.Errorf("Server failed: %s", err)
		os.Exit(1)
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
