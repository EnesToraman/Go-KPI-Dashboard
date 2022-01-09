package controller

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/repository"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

type PlaneController struct{}

func (ec *PlaneController) GetPlaneCountGroupByAirline(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var planeCountGroupByAirlineDTOs []model.PlaneCountGroupByAirlineDTO
		planeRepo := repository.PlaneRepository{}
		planeCountGroupByAirlineDTOs, err := planeRepo.GetPlaneCountGroupByAirline(db)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(planeCountGroupByAirlineDTOs, w)
	}
}
