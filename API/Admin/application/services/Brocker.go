package services

type Brocker struct {
	rmq BrockerMessage
}

func NewBrocker(rmq BrockerMessage) *Brocker {
	return &Brocker{rmq: rmq}
}

func (s *Brocker) Run(id_user int) {
	s.rmq.SendConfirmInstallationMessage(id_user)
}