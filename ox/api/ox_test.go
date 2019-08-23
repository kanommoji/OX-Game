package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "ox/api"
	service "ox/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_SetNewGame_Input_Player_Name_Mo_Symbol_X_Should_Be_HTTP_Status_200(t *testing.T) {
	var gameAPI GameAPI
	expected := []service.Player{service.Player{Name: "mo", Symbol: "x"}}
	body := `{"name":"mo","symbol":"x"}`
	request := httptest.NewRequest("POST", "/ox/new-game/player", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	testRoute := gin.Default()
	testRoute.POST("ox/new-game/player", gameAPI.AddPlayer)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	testRoute.POST("ox/new-game", gameAPI.SetNewGame)

	actual := gameAPI.GameAPI.Player

	if response.StatusCode != http.StatusOK {
		t.Errorf("expect \n'%d' \nbut got it \n'%d'", http.StatusCreated, response.StatusCode)
	}

	assert.Equal(t, expected, actual)
}
