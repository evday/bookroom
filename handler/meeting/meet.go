package meeting

type Createmeet struct {
	Date string 
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	Username string `json:"username"`
	Member string 	`json:"member"`
	Theme string 	`json:"theme"`
	RoomId int 	`json:"roomId"`
}

