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

func TestCreateCarRepository(t *testing.T) {
	assertions := assert.New(t)

	mockCarRepository := car_repo.New(mocks.TestCollection)
	tests := []struct {
		describe string
		arg1     entity.ICar
		name     string
		want     *entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should create a car",
			arg1:     mocks.Car,
			name:     "CreateCar",
			want:     &mocks.Car,
			wantErr:  nil,
			msg:      "Car created successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got, _ := mockCarRepository.CreateCar(context.Background(), &tt.arg1)
			defer mocks.TestCollection.Drop(context.Background())
			tt.want.Id = got.Id
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}
