package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/util"
)

func UploadImage(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(20 << 20); err != nil { // 20MBまで許可
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to upload image"))
			returnError(w, http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("image")
		if err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to upload image"))
			returnError(w, http.StatusBadRequest)
			return
		}
		defer file.Close()

		allowed, fileType := util.GetFileType(file)
		if !allowed {
			c.Logger.Warn("Bad Request", "Error", cerror.ErrInvalidFileType)
			returnError(w, http.StatusBadRequest)
			return
		}

		fileData, err := io.ReadAll(file)
		if err != nil {
			c.Logger.Warn("Internal Server Error", "Error", cerror.Wrap(err, "failed to read file"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		filename := util.NewUuid().String() + fileType
		fmt.Println("filename:", filename)
		if err := c.Controllers.File.SaveImage(fileData, filename, "static"); err != nil {
			c.Logger.Warn("Internal Server Error", "Error", cerror.Wrap(err, "failed to save image"))
			returnError(w, http.StatusInternalServerError)
			return
		}
	}
}
