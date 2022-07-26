package service

import (
	"context"
	"dog-service/util"
	"dog-service/util/logging"
	"fmt"
	grpcDog "protobuf-v1/golang/service"
)

type dogService struct {
	grpcDog.UnimplementedDogServiceServer
}

//NewDogService - initializes the dog service
func NewDogService() grpcDog.DogServiceServer {
	return dogService{}
}

//GetBreed - takes breed name and count as input, does sanity and return image name and bytes
func (d dogService) GetBreed(ctx context.Context, request *grpcDog.GetBreedRequest) (*grpcDog.GetBreedResponse, error) {
	logging.Info("started GetBreed()")
	if len(request.Name) == 0 {
		return nil, fmt.Errorf("name cannot be blank")
	}

	if request.Count <= 0 {
		return nil, fmt.Errorf("positive count expected")
	}

	imageURLs, err := util.GetImageURLs(request.Name, request.Count)
	if err != nil {
		return nil, err
	}

	logging.Info(fmt.Sprintf("%d image urls found", len(imageURLs)))
	resp := &grpcDog.GetBreedResponse{}
	for _, imageURL := range imageURLs {
		fileName, image, err := util.DownloadImage(imageURL)
		if err != nil {
			return nil, err
		}
		resp.Images = append(resp.Images, &grpcDog.Image{Name: fileName, Image: image})
		logging.Info(fmt.Sprintf("%s image downloaded", fileName))
	}

	return resp, nil
}
