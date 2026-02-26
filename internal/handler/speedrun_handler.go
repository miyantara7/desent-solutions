package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SpeedrunHandler struct {
	client *http.Client
}

func NewSpeedrunHandler() *SpeedrunHandler {
	return &SpeedrunHandler{
		client: &http.Client{Timeout: 15 * time.Second},
	}
}

func getBaseURL(c *gin.Context) string {
	scheme := "http"

	if proto := c.Request.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	} else if c.Request.TLS != nil {
		scheme = "https"
	}

	return scheme + "://" + c.Request.Host
}

func decodeBody(resp *http.Response) map[string]interface{} {
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var out map[string]interface{}
	_ = json.Unmarshal(body, &out)

	return map[string]interface{}{
		"status_code": resp.StatusCode,
		"body":        out,
	}
}

func (h *SpeedrunHandler) Run(c *gin.Context) {
	start := time.Now()
	result := make(map[string]interface{})

	baseURL := getBaseURL(c)

	tokenResp, err := h.client.Post(baseURL+"/auth/token", "application/json", nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tokenData := decodeBody(tokenResp)
	result["auth_token"] = tokenData

	token := ""
	if body, ok := tokenData["body"].(map[string]interface{}); ok {
		if t, ok := body["token"].(string); ok {
			token = t
		}
	}

	payload := map[string]string{
		"title":  "Speed Run Book",
		"author": "Runner",
	}
	bodyBytes, _ := json.Marshal(payload)

	createResp, err := h.client.Post(
		baseURL+"/books",
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	createData := decodeBody(createResp)
	result["create_book"] = createData

	bookID := ""
	if body, ok := createData["body"].(map[string]interface{}); ok {
		if id, ok := body["id"].(string); ok {
			bookID = id
		}
	}

	listResp, _ := h.client.Get(baseURL + "/books")
	result["list_books"] = decodeBody(listResp)

	getResp, _ := h.client.Get(baseURL + "/books/" + bookID)
	result["get_book"] = decodeBody(getResp)

	updatePayload := map[string]string{"title": "Speed Run Updated"}
	updateBytes, _ := json.Marshal(updatePayload)

	reqUpdate, _ := http.NewRequest(
		http.MethodPut,
		baseURL+"/books/"+bookID,
		bytes.NewBuffer(updateBytes),
	)
	reqUpdate.Header.Set("Content-Type", "application/json")

	updateResp, _ := h.client.Do(reqUpdate)
	result["update_book"] = decodeBody(updateResp)

	reqProtected, _ := http.NewRequest(
		http.MethodGet,
		baseURL+"/protected/books",
		nil,
	)
	reqProtected.Header.Set("Authorization", "Bearer "+token)

	protectedResp, _ := h.client.Do(reqProtected)
	result["protected_books"] = decodeBody(protectedResp)

	reqDelete, _ := http.NewRequest(
		http.MethodDelete,
		baseURL+"/books/"+bookID,
		nil,
	)
	deleteResp, _ := h.client.Do(reqDelete)
	result["delete_book"] = decodeBody(deleteResp)

	result["duration_ms"] = time.Since(start).Milliseconds()

	c.JSON(http.StatusOK, gin.H{
		"message":  "Speed run completed",
		"base_url": baseURL,
		"result":   result,
	})
}
