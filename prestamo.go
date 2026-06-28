package models

type Prestamo struct {
	ID        int `json:"id"`
	UsuarioID int `json:"usuario_id"`
	LibroID   int `json:"libro_id"`
}
