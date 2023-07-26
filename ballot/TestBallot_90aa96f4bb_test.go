package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
	"testing"
)

type Response struct {
	TotalVotes int
}

type Status struct {
	Code int
}

type Vote struct {
	CandidateID string
	VoterID     string
}

func httpClientRequest(method, url, path string, body io.Reader) (*http.Response, []byte, error) {
	// TODO: Implement this function
	return nil, nil, nil
}

func TestBallot(t *testing.T) {
	// TODO: Define port
	port := ""
	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		t.Error(err)
	}
	log.Println("get ballot resp:", string(result))
	var initalRespData Response
	if err = json.Unmarshal(result, &initalRespData); err != nil {
		log.Printf("Failed to unmarshal get ballot response. %+v", err)
		t.Error(err)
	}

	var ballotvotereq Vote
	ballotvotereq.CandidateID = fmt.Sprint(rand.Intn(10))
	ballotvotereq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotvotereq)
	if err != nil {
		log.Printf("Failed to marshall post ballot request %+v", err)
		t.Error(err)
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		t.Error(err)
	}
	log.Println("post ballot resp:", string(result))
	var postballotResp Status
	if err = json.Unmarshal(result, &postballotResp); err != nil {
		log.Printf("Failed to unmarshal post ballot response. %+v", err)
		t.Error(err)
	}
	if postballotResp.Code != 201 {
		t.Error(errors.New("post ballot resp status code"))
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
		t.Error(err)
	}
	log.Println("get final ballot resp:", string(result))
	var finalRespData Response
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		log.Printf("Failed to unmarshal get final ballot response. %+v", err)
		t.Error(err)
	}
	if finalRespData.TotalVotes-initalRespData.TotalVotes != 1 {
		t.Error(errors.New("ballot vote count error"))
	}
}

func TestTestBallot_90aa96f4bb_Failure(t *testing.T) {
	// TODO: Manipulate conditions for TestBallot to fail
	err := TestBallot()
	if err == nil {
		t.Error("TestBallot was expected to fail, but it didn't")
	} else {
		t.Log("TestBallot failed as expected")
	}
}
