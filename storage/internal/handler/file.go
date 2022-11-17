package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	//"io"
	"io/ioutil"
	"net/http"
)

func (h *Handler) UploadFile(c *gin.Context) {
	ctx := context.Background()
	//ctcxt := c.Request.Context()
	file, err := c.FormFile("File")

	if err != nil {
		ErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	buffer, err := file.Open()
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	email := c.PostForm("Email")

	if email == "" {
		ErrorMessage(c, http.StatusBadRequest, "can not be empty !!!")
		return
	}

	link, err := h.service.Upload(ctx, "test", file.Filename, file.Header["Content-Type"][0], email, file.Size, buffer)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	defer buffer.Close()

	c.JSON(http.StatusOK, map[string]interface{}{
		"link": link,
	})
}

func (h *Handler) DownloadFile(c *gin.Context) {
	key := c.Param("id")
	ctx := context.Background()

	ob, filename, err := h.service.Download(ctx, key)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = ob.Stat()
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	filebyte, err := ioutil.ReadFile("downloadedFiles/" + filename)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.Write(filebyte)

}

func (h *Handler) DeleteObject(c *gin.Context) {
	ctx := context.Background()
	key := c.Param("id")
	err := h.service.Delete(ctx, "test", key)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

}
