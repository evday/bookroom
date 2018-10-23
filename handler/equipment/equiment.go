package equipment

type CreateEquipment struct {
	Name string  `json:"name"`
	Brand string `json:"brand"`
	Model string `json:"bmodel"`
	Store int64 `json:"store"`
	CreateAt string `json:"create_at"`
}