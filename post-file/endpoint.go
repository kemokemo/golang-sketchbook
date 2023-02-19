package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadPost(c *gin.Context) {
	id := "file"
	file, err := c.FormFile(id)
	if err != nil {
		log.Printf("failed to load '%s', %v", id, err)
		c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Cannot use file with name %s sent by Form", id))
		return
	}

	sessionID := "sessionID" // TODO: use sessions middleware
	dstDir := filepath.Join("uploaded-files", sessionID)
	dstFile := filepath.Join("uploaded-files", sessionID, file.Filename)
	saveFile(c, sessionID, file, dstDir, dstFile)

	// TODO: save dstFile to session info

	// TODO: run another buisiness logic with dstFile info

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!\n", file.Filename))
}

func saveFile(c *gin.Context, sessionID string, file *multipart.FileHeader, dstDir, dstFile string) error {
	err := os.RemoveAll(dstDir)
	if err != nil {
		log.Printf("failed to cleanup dst directory '%s', %v", dstDir, err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to prepare a directory for saving '%s'", file.Filename))
		return fmt.Errorf("failed to save file")
	}

	err = os.MkdirAll(dstDir, 0700)
	if err != nil {
		log.Printf("failed to create dst directory '%s', %v", dstDir, err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create a directory for saving '%s'", file.Filename))
		return fmt.Errorf("failed to save file")
	}

	err = c.SaveUploadedFile(file, dstFile)
	if err != nil {
		log.Printf("failed to save file '%s' to '%s', %v", file.Filename, dstFile, err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to save '%s'", file.Filename))
		return fmt.Errorf("failed to save file")
	}

	return nil
}
