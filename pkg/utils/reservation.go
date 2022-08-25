package utils

func GetReservationTypeLabel(reservationTypeID int) string {

	var reservationTypeLabel string

	if reservationTypeID == 1 {
		reservationTypeLabel = "daily"
	} else if reservationTypeID == 2 {
		reservationTypeLabel = "hourly"
	}

	return reservationTypeLabel

}
