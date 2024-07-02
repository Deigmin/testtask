package repository

import (
	serviceerrors "applicationdesigntest/internal/error"
	"applicationdesigntest/internal/model"
	"time"
)

type OrderRepository interface {
	Create(order model.Order) error
	Get() []model.Order
}

type InMemoryOrderRepository struct {
	Orders []model.Order
}

func New(orders []model.Order) *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		Orders: orders,
	}
}

func (r *InMemoryOrderRepository) Create(order model.Order) error {
	r.Orders = append(r.Orders, order)
	return nil
}

func (r *InMemoryOrderRepository) Get() []model.Order {
	return r.Orders
}

type RoomAvailabilityRepository interface {
	GetAvailability(hotelID, roomID string, date time.Time) (*model.RoomAvailability, error)
	UpdateAvailability(hotelID, roomID string, date time.Time, quota int) error
}

type InMemoryRoomAvailabilityRepository struct {
	Availability []model.RoomAvailability
}

func (r *InMemoryRoomAvailabilityRepository) GetAvailability(hotelID, roomID string, date time.Time) (*model.RoomAvailability, error) {
	for i, a := range r.Availability {
		if a.HotelID == hotelID && a.RoomID == roomID && a.Date.Equal(date) {
			return &r.Availability[i], nil
		}
	}
	return nil, serviceerrors.ErrAvailabilityNotFount
}

func (r *InMemoryRoomAvailabilityRepository) UpdateAvailability(hotelID, roomID string, date time.Time, quota int) error {
	for i, a := range r.Availability {
		if a.HotelID == hotelID && a.RoomID == roomID && a.Date.Equal(date) {
			r.Availability[i].Quota = quota
			return nil
		}
	}
	return serviceerrors.ErrAvailabilityNotFount
}
