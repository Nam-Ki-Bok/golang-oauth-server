package utils

func IsValidClient(id, pw string) bool {
	err := mariaDB.Where("client_id = ?", id).
		Where("client_secret = ?", pw).
		Where("server_ip = ?", "1.1.1.1"). // 1.1.1.1 -> c.ClientIP()
		Find(responseClient).Error

	if err != nil {
		return false
	} else {
		return true
	}
}
