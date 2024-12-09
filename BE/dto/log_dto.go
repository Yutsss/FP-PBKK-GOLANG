package dto

type (
	LogResponse struct {
		ID           int64  `json:"id"`
		TechnicianID int64  `json:"technician_id"`
		Activity     string `json:"activity"`
	}
)
