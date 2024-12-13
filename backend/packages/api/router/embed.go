package router

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func MetaImage(image []string) (string, string) {
	Image, Twitter_Image := "", ""
	for _, image_url := range image {
		var og string
		var ogty string
		var trog string
		var card string
		ext := filepath.Ext(image_url)
		switch ext {
		case ".mp4":
			og = "video"
			ogty = "video.other"
			trog = "player"
			card = "player"
		case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp":
			og = "image"
			trog = "image"
			ogty = "image"
			card = "summary_large_image"
		default:
			og = "image"
			trog = "image"
			ogty = "image"
			card = "summary"
		}
		Image += fmt.Sprintf("<meta property=\"og:%s\" content=\"%s\"/>\n", og, image_url)
		Image += fmt.Sprintf("<meta property=\"og:type\" content=\"%s\"/>\n", ogty)
		Twitter_Image += fmt.Sprintf("<meta property=\"twitter:%s\" content=\"%s\"/>\n", og, image_url)
		Twitter_Image += fmt.Sprintf("<meta name=\"twitter:%s\" content=\"%s\">\n", trog, card)
	}
	return Image, Twitter_Image
}

func Embed(c *gin.Context) {
	filename := strings.ReplaceAll(c.Param("filename"), "/", "")
	na := "/content/" + filename
	OpenGraph, Twitter := MetaImage([]string{na})
	data := gin.H{
		"metalink": map[string]interface{}{
			"url":         na,
			"name":        "cdn",
			"title":       "cdn",
			"description": "file-server",
			"color":       "#2bb5f0",
			"images": map[string]string{
				"OpenGraph": OpenGraph,
				"Twitter":   Twitter,
			},
		},
	}
	c.HTML(http.StatusOK, "embed.html", data)
}
