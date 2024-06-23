package api

import (
	"fmt"
	"log"
	"net/http"
	"streamer/internal/dto"

	"github.com/go-chi/render"
)

func UploadFileHandler(s3Data *dto.S3Data) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)

		file, header, err := r.FormFile("file")
		if nil != err {
			panic(err)
		}

		defer file.Close()

		log.Printf("Filename: %s\nFile: %v", header.Filename, file)

		err = s3Data.UploadFile(file, header.Filename)

		if nil != err {
			log.Fatal(err)
			errResponse := dto.ErrInternalServerError(err)
			render.Render(w, r, errResponse)
			return
		}

		okResponse := dto.OkResponseData(fmt.Sprintf("Uploaded file with name: %s", header.Filename))
		render.Render(w, r, okResponse)
	}
}
