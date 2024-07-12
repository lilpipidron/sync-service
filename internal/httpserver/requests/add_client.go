package requests

type AddClientRequest struct {
	ClientName  string `json:"client_name"`
	Version     int    `json:"version"`
	Image       string `json:"image"`
	CPU         string `json:"cpu"`
	Memory      string `json:"memory"`
	NeedRestart bool   `json:"need_restart"`
}
