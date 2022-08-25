package utils

const (
	ROOM_STATUS_OCCOUPIED                     = 11
	ROOM_STATUS_VACANT                        = 12
	ROOM_STATUS_OCCOUPIED_CLEAN               = 13
	ROOM_STATUS_OCCOUPIED_DIRTY               = 14
	ROOM_STATUS_VACANT_CLEAN_INSPECTED        = 15
	ROOM_STATUS_VACANT_CLEAN                  = 16
	ROOM_STATUS_VACANT_DIRTY                  = 17
	ROOM_STATUS_COMPLIMENT                    = 18
	ROOM_STATUS_HOUSE_USE                     = 19
	ROOM_STATUS_DO_NOT_DISTURB                = 20
	ROOM_STATUS_SLEEP_OUT                     = 21
	ROOM_STATUS_SKIPPER                       = 22
	ROOM_STATUS_OUT_OF_SERVICE                = 23
	ROOM_STATUS_OUT_OF_ORDER                  = 24
	ROOM_STATUS_DUE_OUT_OR_EXPECTED_DEPARTURE = 25
	ROOM_STATUS_EXPECTED_ARRIVAL              = 26
	ROOM_STATUS_CHECK_OUT                     = 27
	ROOM_STATUS_LATE_CHECK_OUT                = 28
	ROOM_STATUS_OCCUPIED_NO_LUGGAGE           = 29
	ROOM_STATUS_DOUBLE_LOCK                   = 3
)

func MappingRoomStatus(statusCode int) string {

	var (
		roomStatus     string
		roomStatusList = map[int]string{
			11: "Occupied",
			12: "Vacant",
			13: "Occupied Clean",
			14: "Occupied Dirty",
			15: "Vacant Clean Inspected",
			16: "Vacant Clean",
			17: "Vacant Dirty",
			18: "Compliment",
			19: "House Use",
			20: "Do Not Disturb",
			21: "Sleep Out",
			22: "Skipper",
			23: "Out Of Service",
			24: "Out Of Order",
			25: "Due Out or Expected Depature",
			26: "Expected Arrival",
			27: "Check Out",
			28: "Late Check Out",
			29: "Occupied No Luggage",
			30: "Double Lock",
		}
	)

	roomStatus = roomStatusList[statusCode]

	return roomStatus

}
