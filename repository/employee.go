package repository

import (
	"database/sql"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
)

type EmployeeRepository struct{}

func (er *EmployeeRepository) GetEmployeeCountGroupByTitle(db *sql.DB) ([]model.EmployeeCountGroupByTitleDTO, error) {
	var EmployeeCountGroupByTitleDTO model.EmployeeCountGroupByTitleDTO
	var EmployeeCountGroupByTitleDTOs []model.EmployeeCountGroupByTitleDTO
	query := `
	SELECT title, COUNT(employee_id) FROM employee
	GROUP BY title
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&EmployeeCountGroupByTitleDTO.Title, &EmployeeCountGroupByTitleDTO.NumberOfEmployees); err != nil {
			return nil, err
		}
		EmployeeCountGroupByTitleDTOs = append(EmployeeCountGroupByTitleDTOs, EmployeeCountGroupByTitleDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return EmployeeCountGroupByTitleDTOs, err
}

func (er *EmployeeRepository) GetEmployeeAverageSalaryGroupByTitle(db *sql.DB) ([]model.EmployeeAverageSalaryGroupByTitleDTO, error) {
	var EmployeeAverageSalaryGroupByTitleDTO model.EmployeeAverageSalaryGroupByTitleDTO
	var EmployeeAverageSalaryGroupByTitleDTOs []model.EmployeeAverageSalaryGroupByTitleDTO
	query := `
	SELECT title, ROUND(AVG(salary), 0) FROM employee
	GROUP BY title
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&EmployeeAverageSalaryGroupByTitleDTO.Title, &EmployeeAverageSalaryGroupByTitleDTO.AvgSalary); err != nil {
			return nil, err
		}
		EmployeeAverageSalaryGroupByTitleDTOs = append(EmployeeAverageSalaryGroupByTitleDTOs, EmployeeAverageSalaryGroupByTitleDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return EmployeeAverageSalaryGroupByTitleDTOs, nil
}
