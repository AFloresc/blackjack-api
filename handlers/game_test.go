package handlers

import (
	"blackjack-api/errors"
	"blackjack-api/game"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestStartGameHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/start", nil)
	w := httptest.NewRecorder()

	StartGameHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperado status 200, recibido %d", resp.StatusCode)
	}

	// Validar que el cuerpo sea JSON decodificable
	var body map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Fatalf("respuesta no es JSON válido: %v", err)
	}

	// Validar campos clave
	if _, ok := body["player"]; !ok {
		t.Error("falta el campo 'player' en la respuesta")
	}
	if _, ok := body["dealer"]; !ok {
		t.Error("falta el campo 'dealer' en la respuesta")
	}
	if _, ok := body["playerScore"]; !ok {
		t.Error("falta el campo 'playerScore' en la respuesta")
	}
}

func TestStateHandlerAfterStart(t *testing.T) {
	// Inicia partida
	startReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/start", nil)
	startRes := httptest.NewRecorder()
	StartGameHandler(startRes, startReq)

	// Consulta estado
	stateReq := httptest.NewRequest(http.MethodGet, "/blackjack/api/v1/state", nil)
	stateRes := httptest.NewRecorder()
	StateHandler(stateRes, stateReq)

	resp := stateRes.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperado status 200, recibido %d", resp.StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("respuesta no es JSON válido: %v", err)
	}

	if _, ok := body["playerScore"]; !ok {
		t.Error("falta 'playerScore' en la respuesta")
	}
	if _, ok := body["dealerScore"]; !ok {
		t.Error("falta 'dealerScore' en la respuesta")
	}
}

func TestHitWithoutSession(t *testing.T) {
	session = nil // Asegura que no hay sesión activa
	// Simula petición sin iniciar partida
	req := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/hit", nil)
	res := httptest.NewRecorder()

	HitHandler(res, req)

	resp := res.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("esperado status 400, recibido %d", resp.StatusCode)
	}
}

func TestHitHandler(t *testing.T) {
	// Inicia partida
	startReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/start", nil)
	startRes := httptest.NewRecorder()
	StartGameHandler(startRes, startReq)

	// Ejecuta /blackjack/api/v1/hit
	hitReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/hit", nil)
	hitRes := httptest.NewRecorder()
	HitHandler(hitRes, hitReq)

	resp := hitRes.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperado status 200, recibido %d", resp.StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("respuesta no es JSON válido: %v", err)
	}

	player, ok := body["player"].(map[string]interface{})
	if !ok {
		t.Fatal("no se pudo acceder al campo 'player'")
	}

	hand, ok := player["hand"].([]interface{})
	if !ok {
		t.Fatal("no se pudo acceder al campo 'hand' del jugador")
	}

	if len(hand) < 3 {
		t.Errorf("esperado al menos 3 cartas tras hit, recibido %d", len(hand))
	}
}

func TestStandHandler(t *testing.T) {
	// Inicia partida
	startReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/start", nil)
	startRes := httptest.NewRecorder()
	StartGameHandler(startRes, startReq)

	// Ejecuta /blackjack/api/v1/stand
	standReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/stand", nil)
	standRes := httptest.NewRecorder()
	StandHandler(standRes, standReq)

	resp := standRes.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperado status 200, recibido %d", resp.StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("respuesta no es JSON válido: %v", err)
	}

	if gameOver, ok := body["gameOver"].(bool); !ok || !gameOver {
		t.Error("esperado gameOver=true tras stand")
	}

	if winner, ok := body["winner"].(string); !ok || winner == "" {
		t.Error("esperado campo 'winner' definido tras stand")
	}
}

func TestFullGameFlow(t *testing.T) {
	// Reiniciar sesión global
	session = nil

	// Crear router real
	router := mux.NewRouter()
	api := router.PathPrefix("/blackjack/api/v1").Subrouter()

	api.HandleFunc("/start", StartGameHandler).Methods("POST")
	api.HandleFunc("/hit", HitHandler).Methods("POST")
	api.HandleFunc("/stand", StandHandler).Methods("POST")
	api.HandleFunc("/state", StateHandler).Methods("GET")

	// Paso 1: Iniciar partida
	startReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/start", nil)
	startRes := httptest.NewRecorder()
	router.ServeHTTP(startRes, startReq)

	if startRes.Result().StatusCode != http.StatusOK {
		t.Fatalf("fallo en /start: %d", startRes.Result().StatusCode)
	}

	// Paso 2: Hit (pedir carta)
	hitReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/hit", nil)
	hitRes := httptest.NewRecorder()
	router.ServeHTTP(hitRes, hitReq)

	if hitRes.Result().StatusCode != http.StatusOK {
		t.Fatalf("fallo en /hit: %d", hitRes.Result().StatusCode)
	}

	// Paso 3: Stand (plantarse)
	standReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/stand", nil)
	standRes := httptest.NewRecorder()
	router.ServeHTTP(standRes, standReq)

	if standRes.Result().StatusCode != http.StatusOK {
		t.Fatalf("fallo en /stand: %d", standRes.Result().StatusCode)
	}

	// Paso 4: Consultar estado final
	stateReq := httptest.NewRequest(http.MethodGet, "/blackjack/api/v1/state", nil)
	stateRes := httptest.NewRecorder()
	router.ServeHTTP(stateRes, stateReq)

	if stateRes.Result().StatusCode != http.StatusOK {
		t.Fatalf("fallo en /state: %d", stateRes.Result().StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(stateRes.Body).Decode(&body); err != nil {
		t.Fatalf("respuesta no es JSON válido: %v", err)
	}

	// Validaciones finales
	if gameOver, ok := body["gameOver"].(bool); !ok || !gameOver {
		t.Error("esperado gameOver=true al final de la partida")
	}
	if winner, ok := body["winner"].(string); !ok || winner == "" {
		t.Error("esperado campo 'winner' definido al final de la partida")
	}
}

func TestRestartGameHandler(t *testing.T) {
	// Simula una partida previa
	session = game.NewGameSession()
	session.Hit()
	session.Stand()

	// Reinicia la partida
	req := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/restart", nil)
	res := httptest.NewRecorder()
	RestartGameHandler(res, req)

	resp := res.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperado status 200, recibido %d", resp.StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("respuesta no es JSON válido: %v", err)
	}

	if gameOver, ok := body["gameOver"].(bool); !ok || gameOver {
		t.Error("esperado gameOver=false tras reinicio")
	}
	if _, ok := body["player"].(map[string]interface{}); !ok {
		t.Error("esperado campo 'player' en la nueva sesión")
	}
}

func TestMultipleGameSessions(t *testing.T) {
	router := mux.NewRouter()
	api := router.PathPrefix("/blackjack/api/v1").Subrouter()

	api.HandleFunc("/start", StartGameHandler).Methods("POST")
	api.HandleFunc("/hit", HitHandler).Methods("POST")
	api.HandleFunc("/stand", StandHandler).Methods("POST")
	api.HandleFunc("/state", StateHandler).Methods("GET")

	for i := 0; i < 3; i++ {
		t.Logf("▶️ Partida #%d", i+1)

		// Iniciar partida
		startReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/start", nil)
		startRes := httptest.NewRecorder()
		router.ServeHTTP(startRes, startReq)

		if startRes.Result().StatusCode != http.StatusOK {
			t.Fatalf("fallo en /start: %d", startRes.Result().StatusCode)
		}

		// Hit
		hitReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/hit", nil)
		hitRes := httptest.NewRecorder()
		router.ServeHTTP(hitRes, hitReq)

		if hitRes.Result().StatusCode != http.StatusOK {
			t.Fatalf("fallo en /hit: %d", hitRes.Result().StatusCode)
		}

		// Consultar estado antes de stand
		stateReq := httptest.NewRequest(http.MethodGet, "/blackjack/api/v1/state", nil)
		stateRes := httptest.NewRecorder()
		router.ServeHTTP(stateRes, stateReq)

		var state map[string]interface{}
		if err := json.NewDecoder(stateRes.Body).Decode(&state); err != nil {
			t.Fatalf("respuesta no es JSON válido: %v", err)
		}

		// Solo llamar a /stand si el juego no ha terminado
		if gameOver, ok := state["gameOver"].(bool); ok && !gameOver {
			standReq := httptest.NewRequest(http.MethodPost, "/blackjack/api/v1/stand", nil)
			standRes := httptest.NewRecorder()
			router.ServeHTTP(standRes, standReq)

			if standRes.Result().StatusCode != http.StatusOK {
				t.Fatalf("fallo en /stand: %d", standRes.Result().StatusCode)
			}
		} else {
			t.Logf("⛔️ Partida #%d ya terminada antes de /stand", i+1)
		}

		// Validar estado final
		finalStateReq := httptest.NewRequest(http.MethodGet, "/blackjack/api/v1/state", nil)
		finalStateRes := httptest.NewRecorder()
		router.ServeHTTP(finalStateRes, finalStateReq)

		var final map[string]interface{}
		if err := json.NewDecoder(finalStateRes.Body).Decode(&final); err != nil {
			t.Fatalf("respuesta final no es JSON válido: %v", err)
		}

		if gameOver, ok := final["gameOver"].(bool); !ok || !gameOver {
			t.Errorf("esperado gameOver=true en partida #%d", i+1)
		}
		if winner, ok := final["winner"].(string); !ok || winner == "" {
			t.Errorf("esperado campo 'winner' definido en partida #%d", i+1)
		}
	}
}

func TestHitHandler_NoActiveGame(t *testing.T) {
	// Reset session a nil para simular que no hay partida
	session = nil

	req := httptest.NewRequest("POST", "/blackjack/api/v1/hit", nil)
	rr := httptest.NewRecorder()

	HitHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("esperado status 400, recibido %d", rr.Code)
	}

	var resp errors.ErrorResponse
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("error al parsear JSON: %v", err)
	}

	if resp.Error != "no hay partida activa" {
		t.Errorf("error esperado 'no hay partida activa', recibido '%s'", resp.Error)
	}

	if resp.Details != "Debes iniciar una partida antes de pedir una carta" {
		t.Errorf("details inesperado: %s", resp.Details)
	}

	if resp.ErrorLevel != "Guru Meditation" {
		t.Errorf("error_level inesperado: %s", resp.ErrorLevel)
	}
}

func contains(body, substr string) bool {
	return len(body) > 0 && substr != "" && (body == substr || len(body) >= len(substr) && body[:len(substr)] == substr)
}
