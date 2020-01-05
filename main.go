package sample

import (
	"context"
	"net/http"

	"github.com/awa/go-iap/appstore"
)

func VerifyReceipt(w http.ResponseWriter, r *http.Request) {
	client := appstore.New()
	req := appstore.IAPRequest{
		ReceiptData: "receipt data",
		Password:    "password",
	}
	resp := &appstore.IAPResponse{}
	ctx := context.Background()
	err := client.Verify(ctx, req, resp)

	if err == nil {
		w.Write([]byte("Finish!!"))
		println(resp.Status)
		println(resp.Environment)
	}
}

func SampleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
