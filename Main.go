package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatemehmirarab/gameapp/entity"
	mySQL "github.com/fatemehmirarab/gameapp/repository/mysql"
	"github.com/fatemehmirarab/gameapp/service/userservice"
)

func main() {

	log.Println("start")
	mux := http.NewServeMux()
	mux.HandleFunc("/user/register", userRegisterHandler)
	mux.HandleFunc("/healthchek", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"message":"server is ok"}`)
	})
	http.ListenAndServe("localhost:8080", mux)

}
func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("🔥 Recovered panic:", r)
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(`{"error":"internal server error"}`))
		}
	}()

	fmt.Println("▶️ Handler started")

	if req.Method != http.MethodPost {
		fmt.Println("❌ Invalid method:", req.Method)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte(`{"error":"invalid method"}`))
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("❌ Error reading body:", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	fmt.Println("📦 Body received:", string(data))

	var userReq userservice.RegisterRequest
	if err := json.Unmarshal(data, &userReq); err != nil {
		fmt.Println("❌ JSON unmarshal error:", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	fmt.Printf("✅ Parsed user: %+v\n", userReq)

	mysqlRepo := mySQL.New()
	if mysqlRepo == nil {
		fmt.Println("❌ mysqlRepo is nil")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error":"database error"}`))
		return
	}

	userSvc := userservice.New(mysqlRepo)
	if userSvc.Repo == nil {
		fmt.Println("❌ userSvc is nil")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error":"service init failed"}`))
		return
	}

	_, errRegister := userSvc.Register(userReq)
	if errRegister != nil {
		fmt.Println("❌ Register failed:", errRegister)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, errRegister.Error())))
		return
	}

	writer.Write([]byte(`{"message":"user created"}`))
}

func test() {
	mysqlRepo := mySQL.New()

	user := entity.User{
		Id:          0,
		Name:        "Mohammad",
		PhoneNumber: "09383837745",
	}

	if _, err := mysqlRepo.IsPhoneNumberUnique(user.PhoneNumber); err != nil {
		fmt.Println("isPhoneNumberUnique error %w", err)
	}

	createdUser, err := mysqlRepo.Register(user)

	if err != nil {
		fmt.Println("can not create %w", err)
	} else {
		fmt.Println("User Created %w", createdUser)
	}

}
