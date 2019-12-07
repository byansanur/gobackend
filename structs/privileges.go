package structs

type GetPrivileges struct {
	Id 		int 	`json:"id"`
	Role 	string 	`json:"role"`
}

type GetType struct {
	Id 			int 	`json:"id"`
	TypeRekom 	string 	`json:"type_rekom"`
}
