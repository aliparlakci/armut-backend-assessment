package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

func LoggerWithRequestId(c *gin.Context) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{"request_id": c.GetString("request_id"), "ip": c.Request.RemoteAddr})
}

func BodyReader(reader io.Reader) ([]byte, error) {
	rawBody := make([]byte, 0)
	chunk := make([]byte, 8)
	for {
		n, err := reader.Read(chunk)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		rawBody = append(rawBody, chunk[:n]...)
	}

	return rawBody, nil
}

func AreBodiesEqual(expected gin.H, actual io.Reader) (bool, error) {
	actualBytes, err := BodyReader(actual)
	if err != nil {
		return false, err
	}

	parsedExpected := expected

	var parsedActual gin.H
	err = json.Unmarshal(actualBytes, &parsedActual)
	if err != nil {
		return false, nil
	}

	marshalledExpected, _ := json.Marshal(parsedExpected)
	marshalledActual, _ := json.Marshal(parsedActual)

	return string(marshalledExpected) == string(marshalledActual), nil
}

func AreResultsSameLength(expectedLength int, actual io.Reader) (bool, error) {
	actualBytes, err := BodyReader(actual)
	if err != nil {
		return false, err
	}

	var parsedActual gin.H
	err = json.Unmarshal(actualBytes, &parsedActual)
	if err != nil {
		return false, nil
	}

	results, success := parsedActual["results"].([]interface{})
	if !success {
		return false, nil
	}

	return expectedLength == len(results), nil
}