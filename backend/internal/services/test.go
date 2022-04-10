package services

func (s *Layer) CheckPostgresVersion() (string, error) {
	v, err := s.DBLayer.CheckPostgresVersion()
	if err != nil {
		return "", err
	}
	return v, nil
}
