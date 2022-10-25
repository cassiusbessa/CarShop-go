package tests

import (
	"context"
	"testing"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	car_repo "github.com/cassiusbessa/goCarShop/internal/db/repositories/car"
	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCarsRepository(t *testing.T) {
	assertions := assert.New(t)
	mockCarRepository := car_repo.New(mocks.TestCollection)

	tests := []struct {
		describe string
		arg1     entity.ICar
		name     string
		want     *[]entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should create a car",
			arg1:     mocks.Car,
			name:     "GetAllCars",
			want:     &mocks.Seeds,
			wantErr:  nil,
			msg:      "Return all cars",
		},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			mocks.Seeder()
			defer mocks.TestCollection.Drop(context.Background())
			got, _ := mockCarRepository.GetAllCars(context.Background())
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}
