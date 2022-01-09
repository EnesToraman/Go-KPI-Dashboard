package api

import (
	"net/http"

	"github.com/EnesToraman/Go-KPI-Dashboard/config"
	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
	"github.com/julienschmidt/httprouter"
)

// GetPlanes retrieves the relevant data with join operations, then scans and returns.
func GetPlanes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.CORS(w, r)
	var planeDTO model.PlaneDTO
	var planeDTOs []model.PlaneDTO
	query := `
	SELECT airline_name, COUNT(plane_id) from plane
	GROUP BY airline_name
	`
	db := config.DB()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		utils.RespondWithError(err, http.StatusNoContent, w)
	}
	for rows.Next() {
		if err := rows.Scan(&planeDTO.AirlineName, &planeDTO.NumberOfPlanes); err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		planeDTOs = append(planeDTOs, planeDTO)
	}
	if err := rows.Err(); err != nil {
		utils.RespondWithError(err, http.StatusInternalServerError, w)
	}
	utils.RespondWithSuccess(planeDTOs, w)
}
