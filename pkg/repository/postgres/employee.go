package postgres

import (
	"context"

	"github.com/KrishnaSindhur/go-template/pkg/contract"
)

type EmployeeRepositroy interface {
	List(ctx context.Context, location string) ([]contract.Employee, error)
	Get(ctx context.Context, empID string) (contract.Employee, error)
	Create(ctx context.Context, emp contract.CreateEmployee) (contract.Employee, error)
	Patch(ctx context.Context, empID string, updateEmp contract.UpdateEmployee) (contract.Employee, error)
	Delete(ctx context.Context, empID string) error
}
