package api

type SshInfo struct {
	Id           uint   `json:"id"`
	Username     string `json:"username"`
	Host         string `json:"host"`
	Password     string `json:"password"`
	SshKey       string `json:"sshkey"`
	ConnectionID uint   `json:"connectionid"`
}
