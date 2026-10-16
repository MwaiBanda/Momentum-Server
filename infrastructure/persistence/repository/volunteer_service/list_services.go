package volunteer_service

import (
	"context"
	"fmt"

	domainservice "Momentum/domain/volunteer_service"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) ListServices(ctx context.Context) ([]domainservice.Service, error) {
	var services []domainservice.Service
	query := persistence.LoadFile("repository/volunteer_service/sql/list_services.sql")
	if err := repository.db.SelectContext(ctx, &services, query); err != nil {
		return nil, fmt.Errorf("list volunteer service catalog: %w", err)
	}

	return services, nil
}
