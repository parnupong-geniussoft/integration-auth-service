package loggers

import (
	"encoding/json"
	"integration-auth-service/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MaskData struct {
	Path string   `json:"path"`
	Key  []string `json:"key"`
}

var MaskersRequest = []MaskData{{Path: "/v1/integration-api/request_token", Key: []string{"client_secret"}}}

var MaskersResponse = []MaskData{{Path: "/v1/integration-api/request_token", Key: []string{"access_token"}}}

type LoggerStruct struct {
	CreatedAt      time.Time `db:"created_at"`
	Level          string    `db:"level"`
	Type           string    `db:"type"`
	Method         string    `db:"method"`
	Path           string    `db:"path"`
	Ip             string    `db:"ip"`
	Message        string    `db:"message"`
	DurationMs     int64     `db:"duration_ms"`
	Header         []byte    `db:"header"`
	Request        string    `db:"request"`
	RequestDate    time.Time `db:"request_date"`
	XCorrelationId string    `db:"x_correlation_id"`
}

func (data *LoggerStruct) HandleResponse(c *fiber.Ctx) {
	mBody := HandlerBodyMask(c.Path(), MaskersResponse, c.Response().Body())
	data.Request = string(mBody)
	data.CreatedAt = time.Now()
	durationMs := utils.DurationMS(data.RequestDate)
	data.Type = "response"
	data.DurationMs = durationMs
}

func (data *LoggerStruct) MaskBodyRequest(c *fiber.Ctx) {
	mBody := HandlerBodyMask(c.Path(), MaskersRequest, c.Body())
	data.Request = string(mBody)
}

func (data *LoggerStruct) HandleError(c *fiber.Ctx, err error) {
	data.CreatedAt = time.Now()
	durationMs := utils.DurationMS(data.RequestDate)
	data.Type = "error"
	data.DurationMs = durationMs
	data.Message = err.Error()
}

func (data *LoggerStruct) HeaderConvert(c *fiber.Ctx) {
	headers := c.Request().Header
	headerMap := make(map[string]string)
	headers.VisitAll(func(key, value []byte) {
		if string(key) == "Authorization" {
			headerMap[string(key)] = "*****"
			return
		}
		headerMap[string(key)] = string(value)
	})
	headerConverted, _ := json.Marshal(headerMap)

	data.Header = headerConverted
}

func (data *LoggerStruct) HeaderConvertResponse(c *fiber.Ctx) {
	headers := c.Response().Header
	headerMap := make(map[string]string)
	headers.VisitAll(func(key, value []byte) {
		if string(key) == "Authorization" {
			headerMap[string(key)] = "*****"
			return
		}
		headerMap[string(key)] = string(value)
	})

	headerConverted, _ := json.Marshal(headerMap)

	data.Header = headerConverted
}
