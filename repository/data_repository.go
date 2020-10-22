package repository

import (
	"errors"
	"time"

	"github.com/yenchu/go-demo/model"
)

func NewDataRepository() *DataRepository {
	return &DataRepository{}
}

type DataRepository struct {
}

func (repo *DataRepository) FindAllData() ([]*model.Data, error) {

	rows, err := db.Queryx("select id, latitude, longitude, date_added from data")

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	data := []*model.Data{}

	for rows.Next() {
		var id string
		var lat, lng float32
		var dateAdded time.Time

		err = rows.Scan(&id, &lat, &lng, &dateAdded)

		if err != nil {
			return nil, err
		}

		datum := model.Data{
			ID: id,
			Location: model.Location{
				Lat: lat,
				Lng: lng,
			},
			DateAdded: dateAdded,
		}

		data = append(data, &datum)
	}

	return data, err
}

func (repo *DataRepository) CreateData(datum *model.Data) error {

	if datum == nil {
		return errors.New("Input is nil")
	}

	tx := db.MustBegin()

	tx.MustExec("insert into data (id, latitude, longitude) values (?, ?, ?)",
		datum.ID, datum.Location.Lat, datum.Location.Lng)

	return tx.Commit()
}
