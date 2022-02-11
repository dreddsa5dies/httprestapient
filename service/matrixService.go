package service

import (
	"context"

	"github.com/dreddsa5dies/httprestapient/config"
	"github.com/dreddsa5dies/httprestapient/ent"
)

// MatrixAttackOps
type MatrixAttackOps struct {
	ctx    context.Context
	client *ent.Client
}

// NewMatrixAttackOps
func NewMatrixAttackOps(ctx context.Context) *MatrixAttackOps {
	return &MatrixAttackOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

// MatrixAttacksGetAll - get all matrixAttack query
func (r *MatrixAttackOps) MatrixAttacksGetAll() ([]*ent.MatrixAttack, error) {

	MatrixAttacks, err := r.client.MatrixAttack.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return MatrixAttacks, nil
}

// MatrixAttackGetByID - get matrixAttack by ID (entgo ID)
func (r *MatrixAttackOps) MatrixAttackGetByID(id int) (*ent.MatrixAttack, error) {

	MatrixAttack, err := r.client.MatrixAttack.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return MatrixAttack, nil
}

// MatrixAttackCreate - create new entry matrixAttack
func (r *MatrixAttackOps) MatrixAttackCreate(newMatrixAttack ent.MatrixAttack) (*ent.MatrixAttack, error) {

	newCreatedMatrixAttack, err := r.client.MatrixAttack.Create().
		SetIdMatrix(newMatrixAttack.IdMatrix).
		SetVendorName(newMatrixAttack.VendorName).
		SetNameMatrix(newMatrixAttack.NameMatrix).
		SetVersionMatrix(newMatrixAttack.VersionMatrix).
		SetCreateDate(newMatrixAttack.CreateDate).
		SetModifyDate(newMatrixAttack.ModifyDate).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedMatrixAttack, nil
}

// MatrixAttackUpdate - update entry matrixAttack by entgo ID
func (r *MatrixAttackOps) MatrixAttackUpdate(MatrixAttack ent.MatrixAttack) (*ent.MatrixAttack, error) {

	updatedMatrixAttack, err := r.client.MatrixAttack.UpdateOneID(MatrixAttack.ID).
		SetIdMatrix(MatrixAttack.IdMatrix).
		SetVendorName(MatrixAttack.VendorName).
		SetNameMatrix(MatrixAttack.NameMatrix).
		SetVersionMatrix(MatrixAttack.VersionMatrix).
		SetCreateDate(MatrixAttack.CreateDate).
		SetModifyDate(MatrixAttack.ModifyDate).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedMatrixAttack, nil
}

// MatrixAttackDelete - delete entry matrixAttack by entgo ID
func (r *MatrixAttackOps) MatrixAttackDelete(id int) (int, error) {

	err := r.client.MatrixAttack.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
