package actor

import (
	"crm_serviceV3/entity"
	"crm_serviceV3/repository/mocks"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateActor(t *testing.T) {
	// Membuat instance mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Membuat instance use case dengan mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Membuat actor body
	actorBody := ActorBody{
		Username: "john_doe",
		Password: "password123",
	}

	// Membuat actor yang diharapkan sebagai hasil kembalian dari repository mock
	expectedActor := entity.Actor{
		ID:        1,
		Username:  "john_doe",
		Password:  "hashed_password",
		Verified:  "false",
		Active:    "false",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// Mengatur ekspektasi pemanggilan metode CreateActor pada mock repository
	mockRepo.On("CreateActor", mock.AnythingOfType("*entity.Actor")).Return(expectedActor, nil)

	// Memanggil fungsi CreateActor pada use case
	createdActor, err := uc.CreateActor(actorBody)
	fmt.Println(createdActor)

	// Memastikan bahwa metode CreateActor pada mock repository telah dipanggil dengan argumen yang sesuai
	mockRepo.AssertCalled(t, "CreateActor", mock.AnythingOfType("*entity.Actor"))

	// Memastikan bahwa hasil kembalian sesuai dengan ekspektasi
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, createdActor)
}

func TestGetActorById(t *testing.T) {
	// Membuat instance mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Membuat instance use case dengan mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Membuat actor yang diharapkan sebagai hasil kembalian dari repository mock
	expectedActor := entity.Actor{
		ID:        1,
		Username:  "john_doe",
		Password:  "hashed_password",
		Verified:  "false",
		Active:    "false",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	mockRepo.On("GetActorById", uint(expectedActor.ID)).Return(expectedActor, nil)

	// Memanggil fungsi GetActorById pada use case
	actor, err := uc.GetActorById(uint(expectedActor.ID))
	fmt.Println(actor)

	// Memastikan bahwa hasil kembalian sesuai dengan ekspektasi
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
}
