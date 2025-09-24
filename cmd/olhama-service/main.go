package olhamaservice

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/Enilsonn/Olhama_Service/internal/api"
	"github.com/Enilsonn/Olhama_Service/internal/olhama"
	"github.com/Enilsonn/Olhama_Service/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	llhamaURL := "nao sei qual é"
	olhamaClient := olhama.NewClient(llhamaURL)
	olhamaservice := service.NewIAService(olhamaClient)
	apiHandler := api.NewIAService(olhamaservice)

	r := chi.NewRouter()
	r.Post("/genarete", apiHandler.GenerateHandler)

	server := &http.Server{
		Addr:           ":8080", // lembrar de passar a porta como parametro para o arquivo de variaveis de ambiente
		Handler:        r,
		ReadTimeout:    1 * time.Second, // se p payload for muito grande, breaka
		WriteTimeout:   1 * time.Second, // acho que o próprio llhama lida com isso, mas timitando o tamanhalho da resposta tbm
		TLSNextProto:   map[string]func(*http.Server, *tls.Conn, http.Handler){},
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
