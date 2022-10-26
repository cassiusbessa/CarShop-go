package tests

import (
	"context"
	"net/http"
	"testing"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	car_repo "github.com/cassiusbessa/goCarShop/internal/db/repositories/car"
	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetCarRepository(t *testing.T) {
	assertions := assert.New(t)
	mockCarRepository := car_repo.New(mocks.TestCollection)

	tests := []struct {
		describe string
		arg1     primitive.ObjectID
		name     string
		want     *entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should get a car",
			arg1:     mocks.Seeds[0].Id,
			name:     "GetCar",
			want:     &mocks.Seeds[0],
			wantErr:  nil,
			msg:      "Return a car",
		},
		{
			describe: "Should get a car",
			arg1:     primitive.NewObjectID(),
			name:     "GetCar",
			want:     nil,
			wantErr:  utils.NewError(http.StatusInternalServerError, "not car founded", "repo-get-car", nil),
			msg:      "Return a car",
		},
	}

	t.Run(tests[0].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		got, _ := mockCarRepository.GetCar(context.Background(), tests[0].arg1)
		assertions.Equal(tests[0].want, got, tests[0].msg)
	})
	t.Run(tests[1].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		_, err := mockCarRepository.GetCar(context.Background(), tests[1].arg1)
		assertions.Equal(tests[1].wantErr.Message, err.Message, tests[1].msg)
	})
}
