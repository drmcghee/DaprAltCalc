// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package main
 
import (
	"bytes"
	"encoding/json"
    "log"
	"net/http"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"
)
const blobStorageOutputBindingName = "storage"

type Operands struct {
    OperandOne float32 `json:"operandOne,string"`
    OperandTwo float32 `json:"operandTwo,string"`
}

func save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var operands Operands
	json.NewDecoder(r.Body).Decode(&operands)
	fmt.Println(fmt.Sprintf("%s%f", "Saving ", operands.OperandTwo))

	// **** Start Binding
	// this is the bit that saves to blob storage

	//from py 
	//dapr_port = os.getenv("DAPR_HTTP_PORT", 3500)
	//dapr_url = "http://localhost:{}/v1.0/bindings/sample-topic".format(dapr_port)


	// emcode or marshal operands back to json
	jsonoperands, err := json.Marshal(operands)
	if err != nil {
		panic(err)
	}

	fmt.Printf("http new request with json body: %v\n",  string(jsonoperands))
	req, err :=  http.NewRequest("POST", "http://localhost:3500/v1.0/bindings/storage", bytes.NewBuffer(jsonoperands))
	if err != nil {
        fmt.Println("Error is req: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

    resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	fmt.Printf("read binding resp.Body successfully:\n%v\n", string(data))
	

	// **** End Binding
	
	json.NewEncoder(w).Encode(operands.OperandTwo)
}
 
func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/save", save).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":6001", router))
}