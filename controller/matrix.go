package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dreddsa5dies/httprestapient/ent"
	"github.com/dreddsa5dies/httprestapient/utils"

	"github.com/dreddsa5dies/httprestapient/service"

	"github.com/gorilla/mux"
)

// MatrixAttacksGetAllController - wrapper Response Request JSON & HTTP Status
func MatrixAttacksGetAllController(w http.ResponseWriter, r *http.Request) {

	matrixAttacks, err := service.NewMatrixAttackOps(r.Context()).MatrixAttacksGetAll()
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, matrixAttacks)
}

// MatrixAttackGetByIDController - wrapper Response Request JSON & HTTP Status
func MatrixAttackGetByIDController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	matrixAttack, err := service.NewMatrixAttackOps(r.Context()).MatrixAttackGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, matrixAttack)
}

// MatrixAttackCreateController - wrapper Response Request JSON & HTTP Status
func MatrixAttackCreateController(w http.ResponseWriter, r *http.Request) {

	var newMatrixAttack ent.MatrixAttack
	err := json.NewDecoder(r.Body).Decode(&newMatrixAttack)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	matrixAttack, err := service.NewMatrixAttackOps(r.Context()).MatrixAttackCreate(newMatrixAttack)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, matrixAttack)
}

// MatrixAttackUpdateController - wrapper Response Request JSON & HTTP Status
func MatrixAttackUpdateController(w http.ResponseWriter, r *http.Request) {

	var newMatrixAttackData ent.MatrixAttack
	err := json.NewDecoder(r.Body).Decode(&newMatrixAttackData)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	newMatrixAttackData.ID = id

	updatedUser, err := service.NewMatrixAttackOps(r.Context()).MatrixAttackUpdate(newMatrixAttackData)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, updatedUser)
}

// MatrixAttackDeleteController - wrapper Response Request JSON & HTTP Status
func MatrixAttackDeleteController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewMatrixAttackOps(r.Context()).MatrixAttackDelete(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, deletedID)
}
