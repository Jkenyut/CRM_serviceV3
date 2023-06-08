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

func TestGetAllActor(t *testing.T) {
	// Create an instance of the mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Define the expected results
	expectedPage := uint(1)
	expectedPerPage := uint(10)
	expectedTotal := 100
	expectedTotalPages := uint(10)
	expectedActors := []entity.Actor{
		{ID: 1, Username: "superadmin"},
		{ID: 2, Username: "admin"},
	}
	// Replace with the expected error value

	// Set up the mock repository's behavior
	mockRepo.On("GetAllActor", expectedPage, "").Return(expectedPage, expectedPerPage, expectedTotal, expectedTotalPages, expectedActors, nil)

	// Call the GetAllActor function on the use case
	page, perPage, total, totalPages, actors, err := uc.GetAllActor(expectedPage, "")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedPage, page)
	assert.Equal(t, expectedPerPage, perPage)
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, expectedTotalPages, totalPages)
	assert.Equal(t, expectedActors, actors)

	// Assert that all expected repository methods were called
	mockRepo.AssertExpectations(t)
}

func TestUpdateActorById(t *testing.T) {
	// Create an instance of the mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Define the input parameters
	actorID := uint64(2)
	updateActorBody := UpdateActorBody{
		Username: "admin",
		Password: "new_password",
		Verified: "true",
		Active:   "true",
	}

	// Define the expected results
	expectedUpdatedActor := entity.Actor{
		ID:        actorID,
		Username:  updateActorBody.Username,
		RoleID:    2,
		Password:  "hashed_password", // Replace with the expected hashed password
		Verified:  updateActorBody.Verified,
		Active:    updateActorBody.Active,
		UpdatedAt: time.Now().UTC(),
	}
	// Replace with the expected error value

	// Set up the mock repository's behavior
	mockRepo.On("UpdateActorById", uint(actorID), mock.AnythingOfType("*entity.Actor")).Return(expectedUpdatedActor, nil)

	// Call the UpdateActorById function on the use case
	updatedActor, err := uc.UpdateActorById(uint(actorID), updateActorBody)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedUpdatedActor, updatedActor)

	// Assert that the expected repository method was called with the correct arguments
	mockRepo.AssertCalled(t, "UpdateActorById", uint(actorID), mock.AnythingOfType("*entity.Actor"))
}

func TestDeleteActorById(t *testing.T) {
	// Create an instance of the mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Define the input parameter
	actorID := uint(1)

	// Set up the mock repository's behavior
	mockRepo.On("DeleteActorById", actorID).Return(nil)

	// Call the DeleteActorById function on the use case
	err := uc.DeleteActorById(actorID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected repository method was called with the correct argument
	mockRepo.AssertCalled(t, "DeleteActorById", actorID)
}
func TestActivateActorById(t *testing.T) {
	// Create an instance of the mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Define the input parameter
	actorID := uint(1)

	// Set up the mock repository's behavior
	mockRepo.On("ActivateActorById", actorID).Return(nil)

	// Call the ActivateActorById function on the use case
	err := uc.ActivateActorById(actorID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected repository method was called with the correct argument
	mockRepo.AssertCalled(t, "ActivateActorById", actorID)
}

func TestDeactivateActorById(t *testing.T) {
	// Create an instance of the mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Define the input parameter
	actorID := uint(1)

	// Set up the mock repository's behavior
	mockRepo.On("DeactivateActorById", actorID).Return(nil)

	// Call the ActivateActorById function on the use case
	err := uc.DeactivateActorById(actorID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected repository method was called with the correct argument
	mockRepo.AssertCalled(t, "DeactivateActorById", actorID)
}

func TestLoginActor(t *testing.T) {
	// Create an instance of the mock ActorRepoInterface
	mockRepo := new(mocks.ActorRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := actorUseCaseStruct{
		actorRepository: mockRepo,
	}

	// Define the input parameter
	inputActor := ActorBody{
		Username: "superadmin",
		Password: "password",
	}

	expectedActor := entity.Actor{
		Username: "superadmin",
		Password: "password",
	}

	// Set up the mock repository's behavior
	mockRepo.On("LoginActor", mock.AnythingOfType("*entity.Actor")).Return(expectedActor, nil)

	// Call the LoginActor function on the use case
	actor, err := uc.LoginActor(inputActor)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)

	// Assert that the expected repository method was called with the correct argument
	mockRepo.AssertCalled(t, "LoginActor", mock.AnythingOfType("*entity.Actor"))
}
