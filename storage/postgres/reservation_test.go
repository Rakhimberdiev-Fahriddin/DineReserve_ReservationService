package postgres

import (
	"context"
	"reflect"
	"testing"
	"time"

	pb "reservation-service/generated/reservation_service"

	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := ReservationRepo{DB: db}

	req := &pb.CreateReservationRequest{
		UserId:          "67188541-6344-42bd-8be2-a14c558d30aa",
		RestaurantId:    "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
		ReservationTime: time.Now().Format("2006-01-02 15:04:05"),
		Status:          "Confirmed",
	}

	resp, err := repo.CreateReservation(req)

	expectedRespnce := &pb.CreateReservationResponse{
		Reservation: &pb.Reservation{
			UserId:          "67188541-6344-42bd-8be2-a14c558d30aa",
			RestaurantId:    "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
			ReservationTime: "",
			Status:          "Confirmed",
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, resp.Reservation.UserId, expectedRespnce.Reservation.UserId)
	assert.Equal(t, resp.Reservation.RestaurantId, expectedRespnce.Reservation.RestaurantId)
	assert.Equal(t, resp.Reservation.Status, expectedRespnce.Reservation.Status)
	assert.NotEmpty(t, resp.Reservation.ReservationTime)
	assert.NotEmpty(t, resp.Reservation.Id)
}

func TestListReservations(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := ReservationRepo{DB: db}

	req := &pb.ListReservationsRequest{
		RestaurantId: "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
	}

	resp, err := repo.ListReservations(req)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := &pb.ListReservationsResponse{
		Reservations: []*pb.Reservation{
			{
				Id:              "e93dc146-97dc-417c-b312-f0f1f349ee78",
				UserId:          "67188541-6344-42bd-8be2-a14c558d30aa",
				RestaurantId:    "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
				ReservationTime: "2024-07-10T11:41:40Z",
				Status:          "Confirmed",
			},
		},
	}

	if !reflect.DeepEqual(resp, expectedResponse) {
		t.Errorf("have %v , wont %v", resp, &expectedResponse)
	}
}

func TestGetReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := ReservationRepo{DB: db}

	req := &pb.GetReservationRequest{Id: "e93dc146-97dc-417c-b312-f0f1f349ee78"}
	resp, err := repo.GetReservation(req)
	assert.NoError(t, err)
	
	expectedResponse := pb.GetReservationResponse{
		Reservation: &pb.Reservation{
			Id:              "e93dc146-97dc-417c-b312-f0f1f349ee78",
			UserId:          "67188541-6344-42bd-8be2-a14c558d30aa",
			RestaurantId:    "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
			ReservationTime: "2024-07-10T11:41:40Z",
			Status:          "Confirmed",	
		},
	}

	if !reflect.DeepEqual(resp, &expectedResponse) {
		t.Errorf("have %v , wont %v", resp, &expectedResponse)
	}
}

func TestUpdateReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := ReservationRepo{DB: db}

	req := &pb.UpdateReservationRequest{
		Id:              "e93dc146-97dc-417c-b312-f0f1f349ee78",
		UserId:          "67188541-6344-42bd-8be2-a14c558d30aa",
		RestaurantId:    "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
		ReservationTime: "2024-07-10 11:41:40",
		Status:          "Confirmed",
	}
	resp, err := repo.UpdateReservation(req)
	assert.NoError(t, err)
	
	expectedResponse := &pb.UpdateReservationResponse{
		Reservation: &pb.Reservation{
			Id:              "e93dc146-97dc-417c-b312-f0f1f349ee78",
			UserId:          "67188541-6344-42bd-8be2-a14c558d30aa",
			RestaurantId:    "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
			ReservationTime: "2024-07-10T11:41:40Z",
			Status:          "Confirmed",	
		},
	}

	if !reflect.DeepEqual(resp, expectedResponse) {
		t.Errorf("have %v , wont %v", resp, &expectedResponse)
	}
}

func TestDeleteReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := ReservationRepo{DB: db}

	req := &pb.DeleteReservationRequest{Id: "e93dc146-97dc-417c-b312-f0f1f349ee78"}
	resp, err := repo.DeleteReservation(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Reservation deleted successfully", resp.Message)
}

func TestCheckReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := ReservationRepo{DB: db}

	restaurantId := "a9a9858a-def9-4ab0-9925-a40177cd9b7d"

	req := &pb.CheckReservationRequest{
		RestaurantId: restaurantId,
	}

	resp, err := repo.CheckReservation(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Available)
}

func TestOrderMeals(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := ReservationRepo{DB: db}

	reservationId := "2d73a102-06a9-4d39-a442-93aff8b61b35"
	meals := []*pb.MealOrder{
		{MenuItemId: "ec8c7271-8a29-4f75-9588-fba18a1cf056", Quantity: 2},
	}

	req := &pb.OrderMealsRequest{
		ReservationId: reservationId,
		Meals:         meals,
	}

	resp, err := repo.OrderMeals(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "success", resp.Status)
}
