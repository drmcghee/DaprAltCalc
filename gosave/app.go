// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package main
 
import (
	"encoding/json"
    "log"
	"net/http"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type Operands struct {
    OperandOne float32 `json:"operandOne,string"`
    OperandTwo float32 `json:"operandTwo,string"`
}

func save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var operands Operands
	json.NewDecoder(r.Body).Decode(&operands)
	fmt.Println(fmt.Sprintf("%s%f", "Saving ", operands.OperandOne))

	// **** Start Binding
	// this is the bit that saves to blob storage

	//from py 
	//dapr_port = os.getenv("DAPR_HTTP_PORT", 3500)
	//dapr_url = "http://localhost:{}/v1.0/bindings/sample-topic".format(dapr_port)

	resp, err := http.Get("http://localhost:3500/v1.0/bindings/save-topic");
    if err != nil {
        print(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        print(err)
    }
    fmt.Print(string(body))

	response = requests.post(dapr_url, json=payload)
	print(response.text, flush=True)
	// **** End Binding
	
	json.NewEncoder(w).Encode(operands.OperandOne)
}
 
func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/save", save).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":6001", router))
}