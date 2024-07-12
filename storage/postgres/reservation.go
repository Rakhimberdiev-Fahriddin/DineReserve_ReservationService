package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "reservation-service/generated/reservation_service"

	_ "github.com/lib/pq"
)

func (r *ReservationRepo) CreateReservation(req *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	query := `
		INSERT INTO reservations (
			user_id, 
			restaurant_id, 
			reservation_time, 
			status
		)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4
		)
		RETURNING 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status;
	`
	reservation := &pb.Reservation{}
	err := r.DB.QueryRow(query, req.UserId, req.RestaurantId, req.ReservationTime, req.Status).Scan(
		&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to create reservation: %v", err)
	}
	return &pb.CreateReservationResponse{Reservation: reservation}, nil
}

func (r *ReservationRepo) ListReservations(req *pb.ListReservationsRequest) (*pb.ListReservationsResponse, error) {
	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)
	query := `
		SELECT 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status 
		FROM 
			reservations 
		WHERE deleted_at = 0  `

	if req.RestaurantId != "" {
		params["restaurant_id"] = req.RestaurantId
		filter += " AND restaurant_id = :restaurant_id "
	}
	if req.ReservationTime != "" {
		params["reservation_time"] = req.ReservationTime
		filter += " AND reservation_time = :reservation_time "
	}
	if req.Status != "" {
		params["status"] = req.Status
		filter += " AND status = :status "
	}
	if req.Limit > 0{
		params["limit"] = req.Limit
		filter += " AND limit = :limit "
	}

	if req.Offset > 0{
		params["offset"] = req.Offset
		filter += " AND offset = :offset"
	}
	query += filter

	query, args = ReplaceQueryParams(query, params)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list reservations: %v", err)
	}
	defer rows.Close()

	var reservations []*pb.Reservation
	for rows.Next() {
		reservation := &pb.Reservation{}
		if err := rows.Scan(&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status); err != nil {
			return nil, fmt.Errorf("failed to scan reservations: %v", err)
		}
		reservations = append(reservations, reservation)
	}
	return &pb.ListReservationsResponse{Reservations: reservations}, nil
}

func (r *ReservationRepo) GetReservation(req *pb.GetReservationRequest) (*pb.GetReservationResponse, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status 
		FROM 
			reservations 
		WHERE id = $1 AND deleted_at = 0;
	`
	reservation := &pb.Reservation{}
	err := r.DB.QueryRow(query, req.Id).Scan(
		&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("reservation not found")
		}
		return nil, fmt.Errorf("failed to get reservation: %v", err)
	}
	return &pb.GetReservationResponse{Reservation: reservation}, nil
}

func (r *ReservationRepo) UpdateReservation(req *pb.UpdateReservationRequest) (*pb.UpdateReservationResponse, error) {
	query := `
		UPDATE 
			reservations 
		SET 
			user_id = $2, 
			restaurant_id = $3, 
			reservation_time = $4, 
			status = $5, 
			updated_at = CURRENT_TIMESTAMP
		WHERE 
			id = $1 AND deleted_at = 0
		RETURNING 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status;
	`
	reservation := &pb.Reservation{}
	err := r.DB.QueryRow(query, req.Id, req.UserId, req.RestaurantId, req.ReservationTime, req.Status).Scan(
		&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to update reservation: %v", err)
	}
	return &pb.UpdateReservationResponse{Reservation: reservation}, nil
}

func (r *ReservationRepo) DeleteReservation(req *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	query := `
		UPDATE 
			reservations 
		SET 
			deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)
		WHERE 
			id = $1 AND deleted_at = 0;
	`

	_, err := r.DB.Exec(query, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("reservation not found")
		}
		return nil, fmt.Errorf("failed to delete reservation: %v", err)
	}
	return &pb.DeleteReservationResponse{Message: "Reservation deleted successfully"}, nil
}

func (r *ReservationRepo) CheckReservation(ctx context.Context, in *pb.CheckReservationRequest) (*pb.CheckReservationResponse, error) {
	var exists bool
	err := r.DB.QueryRow(`
		SELECT 
			EXISTS (
				SELECT
					1
				FROM 
					Reservations 
				WHERE 
					restaurant_id = $1 
			)
	`, in.RestaurantId).Scan(&exists)

	if err != nil {
		return nil, err
	}

	return &pb.CheckReservationResponse{Available: exists}, nil
}

func (r *ReservationRepo) OrderMeals(ctx context.Context, in *pb.OrderMealsRequest) (*pb.OrderMealsResponse, error) {
	var reservationTime time.Time
	
	err := r.DB.QueryRow(`
		SELECT
			reservation_time
		FROM
			reservations
		WHERE 
			deleted_at = 0 and id = $1
	`, in.ReservationId).Scan(&reservationTime)
	if err != nil {
		return nil, err
	}

	// Ovqat buyurtmalarini Redis-ga bir marta HSet bilan qo'shish
	mealData := make(map[string]interface{})
	for _, meal := range in.Meals {
		mealData[meal.MenuItemId] = meal.Quantity
	}
	err = r.R.HSet(ctx, in.ReservationId, mealData).Err()
	if err != nil {
		return nil, err
	}

	// Redisda saqlash muddatini belgilash (expiring key)
	expiration := time.Until(reservationTime)
	err = r.R.Expire(ctx, in.ReservationId, expiration).Err()
	if err != nil {
		return nil, err
	}
	
	return &pb.OrderMealsResponse{Status: "success"}, nil
}