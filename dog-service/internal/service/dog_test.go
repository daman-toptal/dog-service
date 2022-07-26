package service

import (
	"context"
	grpcDog "protobuf-v1/golang/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestGetBreedSuccess - success case
func TestGetBreedSuccess(t *testing.T) {
	service := NewDogService()
	req := &grpcDog.GetBreedRequest{
		Name:  "hound",
		Count: 1,
	}

	resp, err := service.GetBreed(context.Background(), req)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, 1, len(resp.GetImages()), "count should be 1")
	assert.Greater(t, len(resp.Images[0].Image), 0, "bytes should be greater than 0")
	assert.Contains(t, resp.Images[0].Name, "jpg", "image should be jpg")
}

//TestGetBreedBreedNotFound - 404 case from dog API
func TestGetBreedBreedNotFound(t *testing.T) {
	service := NewDogService()
	req := &grpcDog.GetBreedRequest{
		Name:  "abc",
		Count: 1,
	}

	_, err := service.GetBreed(context.Background(), req)
	assert.NotNil(t, err, "err should not be nil")
	assert.Contains(t, err.Error(), "breed not found")
}

//TestGetBreedInvalidCount - fail with negative count
func TestGetBreedInvalidCount(t *testing.T) {
	service := NewDogService()
	req := &grpcDog.GetBreedRequest{
		Name:  "abc",
		Count: -1,
	}

	_, err := service.GetBreed(context.Background(), req)
	assert.NotNil(t, err, "err should not be nil")
	assert.Contains(t, err.Error(), "positive count expected")
}

//TestGetBreedBlankBreed - fail with empty breed name
func TestGetBreedBlankBreed(t *testing.T) {
	service := NewDogService()
	req := &grpcDog.GetBreedRequest{
		Name:  "",
		Count: 1,
	}

	_, err := service.GetBreed(context.Background(), req)
	assert.NotNil(t, err, "err should not be nil")
	assert.Contains(t, err.Error(), "name cannot be blank")
}
