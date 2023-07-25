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
	"testing"
)

type Response struct {
	TotalVotes int
}

type Vote struct {
	CandidateID string
	VoterID     string
}

type Status struct {
	Code int
}

func httpClientRequest(method, url, path string, body io.Reader) (int, []byte, error) {
	// TODO: Implement this function
	return 0, nil, nil
}

func ballotTest() error {
	port := "8080" // TODO: Replace with your port
	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}
	log.Println("get ballot resp:", string(result))
	var initalRespData Response
	if err = json.Unmarshal(result, &initalRespData); err != nil {
		log.Printf("Failed to unmarshal get ballot response. %+v", err)
		return err
	}

	var ballotvotereq Vote
	ballotvotereq.CandidateID = fmt.Sprint(rand.Intn(10))
	ballotvotereq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotvotereq)
	if err != nil {
		log.Printf("Failed to marshall post ballot request %+v", err)
		return err
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}
	log.Println("post ballot resp:", string(result))
	var postballotResp Status
	if err = json.Unmarshal(result, &postballotResp); err != nil {
		log.Printf("Failed to unmarshal post ballot response. %+v", err)
		return err
	}
	if postballotResp.Code != 201 {
		return errors.New("post ballot resp status code")
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
		return err
	}
	log.Println("get final ballot resp:", string(result))
	var finalRespData Response
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		log.Printf("Failed to unmarshal get final ballot response. %+v", err)
		return err
	}
	if finalRespData.TotalVotes-initalRespData.TotalVotes != 1 {
		return errors.New("ballot vote count error")
	}
	return nil
}

func TestBallot(t *testing.T) {
	err := ballotTest()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("TestBallot passed")
	}
}

func TestBallotFailure(t *testing.T) {
	// TODO: Setup conditions for failure
	err := ballotTest()
	if err == nil {
		t.Error("Expected an error but didn't get one")
	} else {
		t.Log("TestBallot failed as expected")
	}
}
