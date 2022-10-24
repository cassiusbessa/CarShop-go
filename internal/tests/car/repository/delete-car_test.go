package tests

import (
	"context"
	"fmt"
	"testing"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	car_repo "github.com/cassiusbessa/goCarShop/internal/db/repositories/car"
	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteCarRepository(t *testing.T) {
	assertions := assert.New(t)

	mockCarRepository := car_repo.New(mocks.TestCollection)
	tests := []struct {
		describe string
		arg1     primitive.ObjectID
		name     string
		want     *[]entity.ICar
		wantErr  int64
		msg      string
	}{
		{
			describe: "Should delete a car",
			arg1:     mocks.Car.Id,
			name:     "DeleteCar",
			want:     &mocks.Seeds,
			wantErr:  0,
			msg:      "Car deleted successfully",
		},
		{
			describe: "Should not delete a car",
			arg1:     primitive.NewObjectID(),
			name:     "DeleteCar",
			want:     nil,
			wantErr:  0,
			msg:      "Car not found",
		},
	}

	t.Run(tests[0].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		tests[0].arg1 = mocks.Seeds[0].Id
		mockCarRepository.DeleteCar(context.Background(), tests[0].arg1)
		got, _ := mockCarRepository.GetAllCars(context.Background())
		fmt.Println(got)
		assertions.NotContains(*got, mocks.Seeds[0], tests[0].msg)
		assertions.Equal(len(*got), len(*tests[0].want)-1, tests[0].msg)
	})

	t.Run(tests[1].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		got, _ := mockCarRepository.DeleteCar(context.Background(), tests[1].arg1)
		assertions.Equal(tests[1].wantErr, got, tests[1].msg)
	})

}
