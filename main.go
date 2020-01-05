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
		// TODO: Statusごとに処理
		/*
			switch status {
			case 0:
				return nil

			case 21000:
				message = "The App Store could not read the JSON object you provided."

			case 21002:
				message = "The data in the receipt-data property was malformed or missing."

			case 21003:
				message = "The receipt could not be authenticated."

			case 21004:
				message = "The shared secret you provided does not match the shared secret on file for your account."

			case 21005:
				message = "The receipt server is not currently available."

			case 21007:
				message = "This receipt is from the test environment, but it was sent to the production environment for verification. Send it to the test environment instead."

			case 21008:
				message = "This receipt is from the production environment, but it was sent to the test environment for verification. Send it to the production environment instead."

			case 21010:
				message = "This receipt could not be authorized. Treat this the same as if a purchase was never made."
			}
		*/
		println(resp.Status)
		println(resp.Environment)
	}
}

func SampleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
