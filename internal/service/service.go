package service

import (
	serviceerrors "applicationdesigntest/internal/error"
	"applicationdesigntest/internal/model"
	"applicationdesigntest/internal/repository"
	"time"
)

type BookingService struct {
	orderRepository        repository.OrderRepository
	availabilityRepository repository.RoomAvailabilityRepository
}

func NewBookingService(orderRepository repository.OrderRepository, availabilityRepository repository.RoomAvailabilityRepository) *BookingService {
	return &BookingService{
		orderRepository:        orderRepository,
		availabilityRepository: availabilityRepository,
	}
}

func (s *BookingService) CreateOrder(order model.Order) error {
	daysToBook := daysBetween(order.From, order.To)
	unavailableDays := make(map[time.Time]struct{})

	for _, day := range daysToBook {
		availability, err := s.availabilityRepository.GetAvailability(order.HotelID, order.RoomID, day)
		if err != nil || availability.Quota < 1 {
			unavailableDays[day] = struct{}{}
		}
	}

	if len(unavailableDays) != 0 {
		return serviceerrors.ErrHotelRoomNotAvailable
	}

	for _, day := range daysToBook {
		availability, _ := s.availabilityRepository.GetAvailability(order.HotelID, order.RoomID, day)
		s.availabilityRepository.UpdateAvailability(order.HotelID, order.RoomID, day, availability.Quota-1)
	}

	return s.orderRepository.Create(order)
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
