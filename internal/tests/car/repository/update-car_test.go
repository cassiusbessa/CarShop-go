package tests

import (
	"context"
	"testing"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	car_repo "github.com/cassiusbessa/goCarShop/internal/db/repositories/car"
	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateCarRepository(t *testing.T) {
	assertions := assert.New(t)

	mockCarRepository := car_repo.New(mocks.TestCollection)
	tests := []struct {
		describe string
		arg1     primitive.ObjectID
		arg2     entity.ICar
		name     string
		want     *entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should update a car",
			arg1:     mocks.Seeds[0].Id,
			arg2:     mocks.UpdateCar,
			name:     "UpdateCar",
			want:     &mocks.Seeds[0],
			wantErr:  nil,
			msg:      "Updated car successfully",
		},
		{
			describe: "Should not update a car",
			arg1:     primitive.NewObjectID(),
			arg2:     mocks.UpdateCar,
			name:     "UpdateCar",
			want:     nil,
			wantErr:  utils.NewError(404, "not car founded", "repo-update-car", nil),
			msg:      "Not updated car",
		},
		}

	t.Run(tests[0].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		mockCarRepository.UpdateCar(context.Background(), tests[0].arg1, tests[0].arg2)
		got, _ := mockCarRepository.GetCar(context.Background(), tests[0].arg1)
		assertions.NotEqual(tests[0].want, got, tests[0].msg)
	})

	t.Run(tests[1].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		_, err := mockCarRepository.UpdateCar(context.Background(), tests[1].arg1, tests[1].arg2)
		assertions.Equal(tests[1].wantErr.Message, err.Message, tests[1].msg)
	})

}
