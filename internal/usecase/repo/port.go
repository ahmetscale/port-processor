package repo

import (
	"port-processor/internal/entity"
	"port-processor/pkg/db"
)

type PortRepo struct {
	*db.DB
}

// NewPortRepo returns a new PortRepo
func NewPortRepo(pg *db.DB) *PortRepo {
	return &PortRepo{pg}
}

// Save saves a Port
func (p *PortRepo) Save(port *entity.Port) (err error) {
	return p.DB.Where("unloc = ?", port.Unloc).Assign(*port).FirstOrCreate(port).Error
}
