package repository_test

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yenchu/go-demo/model"
	"github.com/yenchu/go-demo/repository"
)

func init() {

	config := repository.DbConfig{
		DbIP:     "127.0.0.1",
		DbPort:   3306,
		DbName:   "demo",
		Username: "root",
		Password: "root",
	}

	repository.Connect(config)
}

func TestFindAllData(t *testing.T) {

	repo := repository.NewDataRepository()

	data, err := repo.FindAllData()

	if err != nil {
		t.Fatal(err)
	}

	for _, datum := range data {
		fmt.Printf("datum: %v\n", datum)
	}
}

func TestCreateData(t *testing.T) {

	repo := repository.NewDataRepository()

	loc := model.Location{
		Lat: 40.712776,  //25.032969,
		Lng: -74.005974, //121.565414,
	}

	datum := &model.Data{
		ID:       "New York", //"Taipei",
		Location: loc,
	}

	err := repo.CreateData(datum)

	if err != nil {
		t.Fatal(err)
	}
}
