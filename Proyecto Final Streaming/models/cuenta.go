package models

type Cuenta struct {
	CuentaID   string `json:"cuenta_id"`
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
	Estado     string `json:"estado"`
	PlanID     string `json:"plan_id"`
}
