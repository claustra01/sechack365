package usecase

type IFileRepository interface {
	SaveImage(file []byte, filename string, bucket string) error
}

type FileUsecase struct {
	FileRepository IFileRepository
}

func (u *FileUsecase) SaveImage(file []byte, filename string, bucket string) error {
	return u.FileRepository.SaveImage(file, filename, bucket)
}
