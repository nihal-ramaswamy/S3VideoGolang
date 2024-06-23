package main

import (
	"net/http"
	"os"
	"streamer/internal/api"
	"streamer/internal/constants"
	"streamer/internal/dto"
	"streamer/internal/utils"
)

func main() {
	args := os.Args
	argsMap := utils.ReadArgs(args[1:])

	accessKey := utils.GetDotEnvVariable(constants.ACCESS_KEY)
	secretAccessKey := utils.GetDotEnvVariable(constants.SECRET_ACCESS_KEY)

	s3Data := dto.NewS3DataDefault(argsMap[constants.BUCKET], argsMap[constants.REGION], accessKey, secretAccessKey)

	// HTTP Handlers
	router := api.NewRouter()
	router.Get("/upload", api.UploadFileHandler(s3Data))
	http.ListenAndServe(":1337", router)
}
