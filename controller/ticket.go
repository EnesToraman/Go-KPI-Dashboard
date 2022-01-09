package controller

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/repository"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

type TicketController struct{}

func (tc *TicketController) GetTicketCountGroupByDate(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var ticketCountGroupByDateDTOs []model.TicketCountGroupByDateDTO
		ticketRepo := repository.TicketRepository{}
		ticketCountGroupByDateDTOs, err := ticketRepo.GetTicketCountGroupByDate(db)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(ticketCountGroupByDateDTOs, w)
	}
}

func (tc *TicketController) GetTotalPriceGroupByDate(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var ticketTotalPriceGroupByDateDTOs []model.TicketTotalPriceGroupByDateDTO
		ticketRepo := repository.TicketRepository{}
		ticketTotalPriceGroupByDateDTOs, err := ticketRepo.GetTotalPriceGroupByDate(db)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(ticketTotalPriceGroupByDateDTOs, w)
	}
}

func (tc *TicketController) GetTicketAveragePriceGroupByDate(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var TicketAveragePriceGroupByDateDTOs []model.TicketAveragePriceGroupByDateDTO
		ticketRepo := repository.TicketRepository{}
		TicketAveragePriceGroupByDateDTOs, err := ticketRepo.GetTicketAveragePrice(db)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(TicketAveragePriceGroupByDateDTOs, w)
	}
}

func (tc *TicketController) GetTicketCountGroupByClass(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		utils.CORS(w, r)
		var ticketCountGroupByClassDTOs []model.TicketCountGroupByClassDTO
		ticketRepo := repository.TicketRepository{}
		ticketCountGroupByClassDTOs, err := ticketRepo.GetTicketCountGroupByClass(db)
		if err != nil {
			utils.RespondWithError(err, http.StatusInternalServerError, w)
		}
		utils.RespondWithSuccess(ticketCountGroupByClassDTOs, w)
	}
}
