package urlshorter

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"github.com/DelusionTea/praktikum-go/cmd/shortener"
)

type longShortURLs struct {
	Long string
	Short string
}

const endpoint = "http://localhost:8080/"

var mapURLs = make(map[int]longShortURLs)

var globalId = 1

func Shorter(id int) string {
	return fmt.Sprintf("%s%d", endpoint, id)
}

func HandlerCreateShortURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   id := &globalId
   defer r.Body.Close()
   body, err := io.ReadAll(r.Body)
   if err != nil {
	   http.Error(w, err.Error(), 500)
	   return
   }
   long := string(body)
   w.Header().Set("Content-Type", "text/plain")
   w.WriteHeader(201)
   short := Shorting(*id)
   mapURLs[*id] = longShortURLs{
	   Long: long,
	   Short: short,
   }
   *id++
   w.Write([]byte(short))
}

func HandlerGetURLByID(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	param := params.ByName("id")
	log.Println(param)
	id, err := strconv.Atoi(param)
	log.Println(id)
	if err != nil {
		http.Error(w, "Error", 400)
		return
	}

	long := mapURLs[id].Long
	log.Println(long)
	if long == "" {
		http.Error(w, "id error", 400)
		return
	}


	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", long)
	w.WriteHeader(http.StatusTemporaryRedirect)

}

