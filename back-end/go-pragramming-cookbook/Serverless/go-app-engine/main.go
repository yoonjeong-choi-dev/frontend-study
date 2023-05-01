package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	log.SetOutput(os.Stdout)

	// 운영 환경(app engine)에서는 app.yaml 파일에 설정된 GCLOUD_DATASET_ID 변수 가져온다
	// 로컬 테스트 시, export GCLOUD_DATASET_ID=yj-golang 이후 빌드 파일 실행
	projectId := os.Getenv("GCLOUD_DATASET_ID")
	datastoreClient, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalln(err)
	}

	controller := Controller{datastoreClient}

	http.HandleFunc("/", controller.handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "7166"
		log.Printf("Default Port is set to %s\n", port)
	}

	log.Printf("Listeing on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
