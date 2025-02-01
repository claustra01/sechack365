package util

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
)

func GetFileType(file multipart.File) (bool, string) {
	// MIMEタイプを判定
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return false, ""
	}

	// ファイルポインタを先頭に戻す
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return false, ""
	}

	// 判定されたMIMEタイプを取得
	mimeType := http.DetectContentType(buffer)
	fmt.Println("Detected MIME Type:", mimeType)

	// 許可されたMIMEタイプのみ許可
	allowedTypes := map[string]bool{
		"image/png":  true,
		"image/jpeg": true,
	}

	if !allowedTypes[mimeType] {
		return false, ""
	}

	// MIMEタイプに対応する拡張子を取得
	extensions, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(extensions) == 0 {
		return true, ""
	}

	// 最初の拡張子を返す（例: ".jpg"）
	return true, extensions[0]
}
