package models

type Plan struct {
	PlanID string  `json:"plan_id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}
