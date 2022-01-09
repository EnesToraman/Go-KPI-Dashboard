package controller

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/repository"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

type EmployeeController struct{}

func (ec *EmployeeController) GetEmployeeCountGroupByTitle(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var employeeCountGroupByTitleDTOs []model.EmployeeCountGroupByTitleDTO
		employeeRepo := repository.EmployeeRepository{}
		employeeCountGroupByTitleDTOs, err := employeeRepo.GetEmployeeCountGroupByTitle(db)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(employeeCountGroupByTitleDTOs, w)
	}
}

func (ec *EmployeeController) GetEmployeeAverageSalaryGroupByTitle(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var employeeAverageSalaryGroupByTitleDTOs []model.EmployeeAverageSalaryGroupByTitleDTO
		employeeRepo := repository.EmployeeRepository{}
		employeeAverageSalaryGroupByTitleDTOs, err := employeeRepo.GetEmployeeAverageSalaryGroupByTitle(db)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(employeeAverageSalaryGroupByTitleDTOs, w)
	}
}
