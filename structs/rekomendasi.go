package structs

type CreateRekomendasi struct {
	Id int `json:"id"`
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto string `json:"foto"`
	Lat string `json:"lat"`
	Lng string `json:"lng"`
	Rating string `json:"rating"`
	IdType int `json:"id_type"`
	CreatedAt string `json:"created_at"`
}
type GetRekomendasi struct {
	Id int `json:"id"`
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto string `json:"foto"`
	Rating string `json:"rating"`
	IdType int `json:"id_type"`
	TypeRekom string `json:"type_rekom"`
}

type UpdateRekomendasi struct {
	Id int `json:"id"`
	Nama string `json:"nama"`
	Alaamat string `json:"alaamat"`
	Foto string `json:"foto"`
	Rating string `json:"rating"`
}

type CheckIdRekom struct {
	CountId int `json:"count_id"`
}
