package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"RaribleAPI/internal/model"
)

func TestGetNFTOwnerships_Success(t *testing.T) {

	fakeResp := model.OwnershipResponse{ID: "test-id", TokenID: "123", Value: "1"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(fakeResp)
	}))
	defer server.Close()

	client := NewRaribleClient(server.URL, "")
	resp, err := client.GetNFTOwnerships(context.Background(), "test-id")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.ID != "test-id" {
		t.Errorf("expected ID 'test-id', got %s", resp.ID)
	}
}

func TestGetNFTOwnerships_BadStatus(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewRaribleClient(server.URL, "")
	_, err := client.GetNFTOwnerships(context.Background(), "test-id")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetTraitRaritiesPOST_Success(t *testing.T) {

	fakeResp := model.RarityResponse{Continuation: "", Traits: []model.RarityTrait{{Key: "Hat", Value: "Halo", Rarity: "0.1"}}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(fakeResp)
	}))
	defer server.Close()

	client := NewRaribleClient(server.URL, "")
	body := model.RarityRequest{CollectionId: "col", Properties: []model.Property{{Key: "Hat", Value: "Halo"}}}
	resp, err := client.GetTraitRaritiesPOST(context.Background(), body)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(resp.Traits) != 1 || resp.Traits[0].Key != "Hat" {
		t.Errorf("expected trait 'Hat', got %+v", resp.Traits)
	}
}

func TestGetTraitRaritiesPOST_BadStatus(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := NewRaribleClient(server.URL, "")
	body := model.RarityRequest{CollectionId: "col", Properties: []model.Property{{Key: "Hat", Value: "Halo"}}}
	_, err := client.GetTraitRaritiesPOST(context.Background(), body)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
