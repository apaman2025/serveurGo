package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Fonction pour gérer l’upload de fichiers
func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur lors de l'upload du fichier"})
		return
	}

	// Créer le dossier "uploads" s'il n'existe pas
	os.MkdirAll("uploads", os.ModePerm)

	// Enregistrer le fichier reçu
	savePath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de sauvegarder le fichier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fichier reçu avec succès",
		"path":    savePath,
	})
}

func main() {
	// Initialisation du serveur Gin
	r := gin.Default()
	r.POST("/upload", uploadFile)

	fmt.Println("🚀 Serveur en ligne sur : http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erreur au démarrage du serveur :", err)
	}
}
