package handlers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// UploadFile handles the file upload via a REST API endpoint.
func UploadCSVFile(c echo.Context) error {
	// 1. Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]string{"error": fmt.Sprintf("failed to get file from form: %v", err)})
	}
	fmt.Printf("Received file: %+v\n", file)

	// 2. Open the file
	src, err := file.Open()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to open file"})
	}
	defer src.Close()

	// 3. Create the destination directory
	uploadDir := filepath.Join("uploads", "csv_files")
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return c.JSON(500, map[string]string{"error": "failed to create directory"})
		}
	}

	// 4. Generate a unique filename
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		return c.JSON(400, map[string]string{"error": "invalid file type, no extension found"})
	}
	newFileName := uuid.New().String() + ext
	destinationPath := filepath.Join(uploadDir, newFileName)

	// 5. Create the destination file on the server
	dst, err := os.Create(destinationPath)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to create destination file"})
	}
	defer dst.Close()

	// 6. Copy the file contents
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(500, map[string]string{"error": "failed to copy file contents"})
	}

	// 7. Return the path of the saved file
	filePath := fmt.Sprintf("/%s/%s", filepath.Join("uploads", "csv_files"), newFileName)
	return c.JSON(200, map[string]string{"filePath": filePath})
}