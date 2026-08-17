package models

type ProveedorVideo struct {
	ProveedorID     string `json:"proveedor_id"`
	NombreProveedor string `json:"nombre_proveedor"`
	URLBase         string `json:"url_base"`
	Estado          string `json:"estado"`
}
