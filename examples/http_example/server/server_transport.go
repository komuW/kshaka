package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/komuw/kshaka"
)

// HTTPtransport provides a http based transport that can be
// used to communicate with kshaka/CASPaxos on remote machines.
type HTTPtransport struct {
	NodeAddrress string
	NodePort     string
}

type HTTPtransportPrepareRequest struct {
	B   kshaka.Ballot
	Key []byte
}

func (ht *HTTPtransport) TransportPrepare(b kshaka.Ballot, key []byte) (kshaka.AcceptorState, error) {
	fmt.Println("TransportPrepare called....")
	acceptedState := kshaka.AcceptorState{}

	prepReq := HTTPtransportPrepareRequest{B: b, Key: key}
	url := "http://" + ht.NodeAddrress + ":" + ht.NodePort + "/prepare"
	prepReqJSON, err := json.Marshal(prepReq)
	if err != nil {
		return acceptedState, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(prepReqJSON))
	if err != nil {
		return acceptedState, err
	}
	req.Header.Set("Content-Type", "application/json")
	// todo: ideally, client should be resused across multiple requests
	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	if err != nil {
		return acceptedState, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return acceptedState, errors.New(fmt.Sprintf("url:%v returned http status:%v instead of status:%v", url, resp.StatusCode, http.StatusOK))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return acceptedState, err
	}

	err = json.Unmarshal(body, &acceptedState)
	if err != nil {
		return acceptedState, err
	}

	fmt.Println("TransportPrepare response body::", body, string(body), acceptedState)
	return acceptedState, nil
}

type HTTPtransportAcceptRequest struct {
	B     kshaka.Ballot
	Key   []byte
	State []byte
}

func (ht *HTTPtransport) TransportAccept(b kshaka.Ballot, key []byte, state []byte) (kshaka.AcceptorState, error) {
	fmt.Println("TransportAccept called....")
	acceptedState := kshaka.AcceptorState{}
	acceptReq := HTTPtransportAcceptRequest{B: b, Key: key, State: state}
	url := "http://" + ht.NodeAddrress + ":" + ht.NodePort + "/accept"
	acceptReqJSON, err := json.Marshal(acceptReq)
	if err != nil {
		return acceptedState, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(acceptReqJSON))
	if err != nil {
		return acceptedState, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	if err != nil {
		return acceptedState, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return acceptedState, errors.New(fmt.Sprintf("url:%v returned http status:%v instead of status:%v", url, resp.StatusCode, http.StatusOK))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return acceptedState, err
	}

	err = json.Unmarshal(body, &acceptedState)
	if err != nil {
		return acceptedState, err
	}

	fmt.Println("TransportAccept response body::", body, acceptedState)
	return acceptedState, nil
}