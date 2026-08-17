package models

import "time"

type Pago struct {
	PagoID   string    `json:"pago_id"`
	CuentaID string    `json:"cuenta_id"`
	Fecha    time.Time `json:"fecha"`
	Monto    float64   `json:"monto"`
}
