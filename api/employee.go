package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

// GetEmployeeTitle retrieves the relevant data with join operations, then scans and returns.
func GetEmployeeTitle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var employeeDTO model.EmployeeDTO
	var employeeDTOs []model.EmployeeDTO
	query := `
	SELECT title, COUNT(employee_id) FROM employee
	GROUP BY title
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		utils.RespondWithError(err, http.StatusNoContent, w)
	}
	for rows.Next() {
		if err := rows.Scan(&employeeDTO.Title, &employeeDTO.NumberOfEmployees); err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		employeeDTOs = append(employeeDTOs, employeeDTO)
	}
	if err := rows.Err(); err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	utils.RespondWithSuccess(employeeDTOs, w)
}

// GetEmployeeSalary retrieves the relevant data with join operations, then scans and returns.
func GetEmployeeSalary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var employeeSalary model.EmployeeSalary
	var employeeSalaries []model.EmployeeSalary
	query := `
	SELECT title, ROUND(AVG(salary), 0) FROM employee
	GROUP BY title
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	for rows.Next() {
		if err := rows.Scan(&employeeSalary.Title, &employeeSalary.AvgSalary); err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		employeeSalaries = append(employeeSalaries, employeeSalary)
	}
	if err := rows.Err(); err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	utils.RespondWithSuccess(employeeSalaries, w)
}
