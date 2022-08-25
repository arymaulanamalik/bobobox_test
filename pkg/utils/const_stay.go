package utils

const (
	STAY_STATUS_RESERVED  = 0
	STAY_STATUS_ARRIVAL   = 1
	STAY_STATUS_IN_HOUSE  = 2
	STAY_STATUS_DEPARTURE = 3
	STAY_STATUS_NO_SHOW   = 4
	STAY_STATUS_REFUND    = 5
)

const (
	STAY_CHECK_IN_ALLOWED        = 1
	STAY_CHECK_IN_ALLOWED_LABEL  = "Check In"
	STAY_CHECK_OUT_ALLOWED       = 2
	STAY_CHECK_OUT_ALLOWED_LABEL = "Check Out"
	STAY_IN_HOUSE                = 3
	STAY_IN_HOUSE_LABEL          = "Checked In"
	STAY_WAITING                 = 4
	STAY_WAITING_LABEL           = "Waiting for check in..."
	STAY_NO_SHOW                 = 5
	STAY_NO_SHOW_LABEL           = "No Show"
	STAY_CHECKED_OUT             = 6
	STAY_CHECKED_OUT_LABEL       = "Checked Out"
)

const (
	GUEST_TYPE_PRIMARY   = "primary"
	GUEST_TYPE_SECONDARY = "secondary"

	SHARE_STAY_AVAILABLE = "available"
	SHARE_STAY_ACCEPTED  = "accepted"
	SHARE_STAY_REMOVED   = "removed"
)
