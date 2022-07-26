package command

import (
	"context"
	"dog-service/util"
	"dog-service/util/config"
	"dog-service/util/logging"
	"fmt"
	"os"
	grpcDog "protobuf-v1/golang/service"
	"strconv"
	"strings"

	"go.elastic.co/apm/module/apmgrpc"
	ggrpc "google.golang.org/grpc"
)

const (
	CommandGetBreedPhoto  = commandGetBreedPhoto("get photo ")
	CommandExit           = commandExit("exit")
)

type CommandService interface {
	Execute() (string, error)
}

//NewCommand - returns command executor
func NewCommand(str string) (CommandService, error) {

	if strings.HasPrefix(str, string(CommandGetBreedPhoto)) {
		return commandGetBreedPhoto(str), nil
	}

	if strings.HasPrefix(str, string(CommandExit)) {
		return CommandExit, nil
	}

	return nil, fmt.Errorf("invalid command")

}

type commandGetBreedPhoto string

//Execute - connects to grpc service, parses inputs and save images in files
func (c commandGetBreedPhoto) Execute() (string, error) {
	//connect to grpc service
	port := config.GetString("grpc.port")
	cnString := fmt.Sprintf("%s:%s", config.GetString("grpc.name"), port)
	serviceConn, err := ggrpc.Dial(cnString, ggrpc.WithInsecure(), ggrpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
	if err != nil {
		logging.Fatalf("Error initializing grpc service client, err=%s", err.Error())
		fmt.Println(err.Error())
		return "", err
	}

	defer serviceConn.Close()

	dogServiceClient := grpcDog.NewDogServiceClient(serviceConn)

	//parse and validate input
	parts := strings.Split(string(c), " ")

	if len(parts) < 3 || len(parts[2]) == 0 {
		fmt.Println("breed name is required")
		return "", fmt.Errorf("breed name is required")
	}
	breed := parts[2]
	count := int32(1)

	if len(parts) > 4 {
		fmt.Println("too many params")
		return "", fmt.Errorf("too many params")
	}

	if len(parts) > 3 {
		count64Bit, err := strconv.ParseInt(parts[3], 10, 32)
		if err != nil || count64Bit <= 0{
			fmt.Println("invalid count")
			return "", fmt.Errorf("invalid count")
		}
		count = int32(count64Bit)
	}

	// get response from grpc
	imagesResp, err := dogServiceClient.GetBreed(context.Background(), &grpcDog.GetBreedRequest{Name: breed, Count: count})
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	// ensure directory and create file(s)
	err = util.EnsureDir("images")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	for _, image := range imagesResp.Images {
		path := fmt.Sprintf("./images/%s_%s", breed, image.Name)
		err = util.SaveImage(image.GetImage(), path)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
		fmt.Printf(fmt.Sprintf("%s saved at %s", image.Name, path) + "\n")
	}

	return "success", nil
}

type commandExit string

func (c commandExit) Execute() (string, error) {
	os.Exit(0)
	return "success", nil
}
