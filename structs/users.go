package structs

type CreateUsers struct {
	Id 				int 	`json:"id"`
	Nama 			string 	`json:"nama"`
	Username 		string 	`json:"username"`
	Password 		string 	`json:"password"`
	TglLahir	 	string 	`json:"tgl_lahir"`
	NoKtp 			string 	`json:"no_ktp"`
	NoHp 			string 	`json:"no_hp"`
	NoVisa 			string 	`json:"no_visa"`
	NoPasspor	 	string 	`json:"no_passpor"`
	Foto 			string 	`json:"foto"`
	IdPrivileges 	int 	`json:"id_privileges"`
	CreatedAt 		string 	`json:"created_at"`
}

type GetUser struct {
	Id 				int 	`json:"id"`
	Nama 			string 	`json:"nama"`
	Username 		string 	`json:"username"`
	TglLahir	 	string 	`json:"tgl_lahir"`
	NoKtp 			int 	`json:"no_ktp"`
	NoHp 			int 	`json:"no_hp"`
	NoVisa 			string 	`json:"no_visa"`
	NoPasspor	 	string 	`json:"no_passpor"`
	Foto 			string 	`json:"foto"`
	IdPrivileges 	int 	`json:"id_privileges"`
	CreatedAt 		string 	`json:"created_at"`
}

type GetUserLogin struct {
	Id 				int 		`json:"id"`
	Nama 			*string 	`json:"nama"`
	Username 		*string 	`json:"username"`
	TglLahir	 	*string 	`json:"tgl_lahir"`
	NoKtp 			*int 		`json:"no_ktp"`
	NoHp 			*int 		`json:"no_hp"`
	NoVisa 			*string 	`json:"no_visa"`
	NoPasspor	 	*string 	`json:"no_passpor"`
	Foto 			*string 	`json:"foto"`
	IdPrivileges 	*int 		`json:"id_privileges"`
	Role			*string		`json:"role"`
	CreatedAt 		*string 	`json:"created_at"`
	Token 			string 		`json:"token"`
}

type CekUserLogin struct {
	Id 			int 	`json:"id"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
}

type CheckId struct {
	CountId int `json:"count_id"`
}

type UpdateUsers struct {
	Id 				int 	`json:"id"`
	Nama 			string 	`json:"nama"`
	Username 		string 	`json:"username"`
	Password 		string 	`json:"password"`
	NoHp 			int 	`json:"no_hp"`
	NoVisa 			string 	`json:"no_visa"`
	NoPasspor	 	string 	`json:"no_passpor"`
	Foto 			string 	`json:"foto"`
} 
