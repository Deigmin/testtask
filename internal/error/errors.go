package serviceerrors

import "errors"

var (
	ErrAvailabilityNotFount  = errors.New("availability not found")
	ErrHotelRoomNotAvailable = errors.New("hotel room is not available for selected dates")
)
