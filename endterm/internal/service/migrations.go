package service

import "twitter/internal/store"

type MigrationService interface {
	Migrate()
}

type migrationService struct {
	store store.Store
}

func NewMigrationService(s store.Store) MigrationService {
	return &migrationService{
		store: s,
	}
}

func (s *migrationService) Migrate() {
	s.store.Migrate()
}
