package converter

import (
	"legato_server/api"
	legatoDb "legato_server/db"
)

func SshInfoToSshDb(s *api.SshInfo) legatoDb.Ssh {
	var ss legatoDb.Ssh
	ss.Host = s.Host
	ss.Username = s.Username
	ss.Password = s.Password
	ss.SshKey = s.SshKey

	return ss
}

func SshDbToSshInfo(ss *legatoDb.Ssh) api.SshInfo {
	var s api.SshInfo
	s.Id = ss.ID
	s.Host = ss.Host
	s.Username = ss.Username
	s.Password = ss.Password
	s.SshKey = ss.SshKey
	s.ConnectionID = ss.ConnectionID

	return s
}
