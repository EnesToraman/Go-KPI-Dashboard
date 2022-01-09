package repository

import (
	"database/sql"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
)

type PlaneRepository struct{}

func (pr *PlaneRepository) GetPlaneCountGroupByAirline(db *sql.DB) ([]model.PlaneCountGroupByAirlineDTO, error) {
	var PlaneCountGroupByAirlineDTO model.PlaneCountGroupByAirlineDTO
	var PlaneCountGroupByAirlineDTOs []model.PlaneCountGroupByAirlineDTO
	query := `
	SELECT a.name, COUNT(p.plane_id) FROM plane AS p
	INNER JOIN airline AS a ON a.airline_id = p.airline_id
	GROUP BY a.name
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&PlaneCountGroupByAirlineDTO.AirlineName, &PlaneCountGroupByAirlineDTO.NumberOfPlanes); err != nil {
			return nil, err
		}
		PlaneCountGroupByAirlineDTOs = append(PlaneCountGroupByAirlineDTOs, PlaneCountGroupByAirlineDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return PlaneCountGroupByAirlineDTOs, nil
}
