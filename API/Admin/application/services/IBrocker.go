package services

type BrockerMessage interface {
	SendConfirmInstallationMessage(id_user int)
}