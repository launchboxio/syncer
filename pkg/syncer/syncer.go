package syncer

type Syncer struct {
	Config *Config
}

func New(config *Config) *Syncer {
	return &Syncer{
		Config: config,
	}
}

func (s *Syncer) Run() error {
	return nil
}
