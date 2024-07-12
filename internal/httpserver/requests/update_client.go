package requests

type UpdateClientRequest struct {
	ID          int64  `json:"id"`
	ClientName  string `json:"client_name,omitempty"`
	Version     int    `json:"version,omitempty"`
	Image       string `json:"image,omitempty"`
	CPU         string `json:"cpu,omitempty"`
	Memory      string `json:"memory,omitempty"`
	NeedRestart bool   `json:"need_restart,omitempty"`
}
