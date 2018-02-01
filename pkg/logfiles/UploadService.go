package logfiles

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	UPLOADS_FOLDER string = "./uploads/"
)

/*
UploadService provides method to validate and upload files
*/
type UploadService struct {
	logger *logrus.Entry
}

/*
NewUploadService creates a new service
*/
func NewUploadService(logger *logrus.Entry) *UploadService {
	return &UploadService{
		logger: logger,
	}
}

/*
Upload takes a FileHeader from a file upload multipart-form and copies it to
the uploads folder
*/
func (s *UploadService) Upload(fileHeader *multipart.FileHeader) (int64, error) {
	var err error
	var sourceFile multipart.File
	var destinationFile *os.File
	var bytesWritten int64

	/*
	 * Read the source file
	 */
	if sourceFile, err = fileHeader.Open(); err != nil {
		return 0, errors.Wrapf(err, "Error opening source file %s", fileHeader.Filename)
	}

	defer sourceFile.Close()

	/*
	 * Upload to our uploads folder
	 */
	if destinationFile, err = os.Create(filepath.Join(UPLOADS_FOLDER, fileHeader.Filename)); err != nil {
		return 0, errors.Wrapf(err, "Error creating destination file %s at %s", fileHeader.Filename, UPLOADS_FOLDER)
	}

	defer destinationFile.Close()

	if bytesWritten, err = io.Copy(destinationFile, sourceFile); err != nil {
		return 0, errors.Wrapf(err, "Error copying file %s to destination %s", fileHeader.Filename, UPLOADS_FOLDER)
	}

	return bytesWritten, nil
}
