package room

type CreateRoom struct {
	ID int64 `json:"id"`
	Name string `json:"roomname"`
	Location string `json:"location"`
	Capacity string `json:"capacity"`
	State bool `json:"state"`
	Equipment []interface{}
}
