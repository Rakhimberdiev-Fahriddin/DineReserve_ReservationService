package service

import (
	"context"
	"log/slog"
	pb "reservation-service/generated/reservation_service"
	"reservation-service/storage/postgres"
)

type ReservationService struct{
	pb.UnimplementedReservationServiceServer
	Reservation postgres.ReservationRepo
	Logger *slog.Logger
}

func NewRRestaurantService(reservation postgres.ReservationRepo)*ReservationService{
	return &ReservationService{Reservation: reservation}
}

func (r *ReservationService) CreateRestaurant(ctx context.Context,restaurant *pb.CreateRestaurantRequest)(*pb.CreateRestaurantResponse,error){
	r.Logger.Info("Create Restaurant")
	res,err := r.Reservation.CreateRestaurant(restaurant)
	if err != nil{
		r.Logger.Error("Failed created to restaurant", "error", err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) ListRestaurants(ctx context.Context, listRestaurant *pb.ListRestaurantsRequest)(*pb.ListRestaurantsResponse,error){
	r.Logger.Info("List Restaurant")
	res,err := r.Reservation.ListRestaurants(listRestaurant)
	if err != nil{
		r.Logger.Error("Failed get restaurants","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) GetRestaurant(ctx context.Context, id *pb.GetRestaurantRequest)(*pb.GetRestaurantResponse,error){
	r.Logger.Info("Get Restaurant")
	res,err := r.Reservation.GetRestaurant(id)
	if err != nil{
		r.Logger.Error("Failed get restaurant","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) UpdateRestaurant(ctx context.Context,updateRestaurant *pb.UpdateRestaurantRequest)(*pb.UpdateRestaurantResponse,error){
	r.Logger.Info("Update to Restaurant")
	res,err := r.Reservation.UpdateRestaurant(updateRestaurant)
	if err != nil{
		r.Logger.Error("Failed update to restaurant","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) DeleteRestaurant(ctx context.Context, id *pb.DeleteRestaurantRequest)(*pb.DeleteRestaurantResponse,error){
	r.Logger.Info("Delete for Restaurant")
	res,err := r.Reservation.DeleteRestaurant(id)
	if err != nil{
		r.Logger.Error("Failed delete to restaurant","error",err.Error())
		return nil,err
	}
	return res,nil
}






func (r *ReservationService) CreateReservation(ctx context.Context,reservation *pb.CreateReservationRequest)(*pb.CreateReservationResponse,error){
	r.Logger.Info("Create to Reservation")
	res,err := r.Reservation.CreateReservation(reservation)
	if err != nil{
		r.Logger.Error("Failed create to reservation","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) ListReservations(ctx context.Context,listReservation *pb.ListReservationsRequest)(*pb.ListReservationsResponse,error){
	r.Logger.Info("List Reservation")
	res,err :=r.Reservation.ListReservations(listReservation)
	if err != nil{
		r.Logger.Error("Failed get reservations","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) GetReservation(ctx context.Context,Reservation *pb.GetReservationRequest)(*pb.GetReservationResponse,error){
	r.Logger.Info("Get Reservation")
	res,err := r.Reservation.GetReservation(Reservation)
	if err != nil{
		r.Logger.Error("Failed get reservation","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) UpdateReservation(ctx context.Context, updateReservation *pb.UpdateReservationRequest)(*pb.UpdateReservationResponse,error){
	r.Logger.Error("Update Reservation")
	res,err := r.Reservation.UpdateReservation(updateReservation)
	if err != nil{
		r.Logger.Error("Failed update to reservation","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) DeleteReservation(ctx context.Context, id *pb.DeleteReservationRequest)(*pb.DeleteReservationResponse,error){
	r.Logger.Info("Delete Rservation")
	res,err := r.Reservation.DeleteReservation(id)
	if err != nil{
		r.Logger.Error("Failed delete to reservation","error",err.Error())
		return nil,err
	}
	return res,nil
}







func (r * ReservationService) CreateMenuItem(ctx context.Context, menu *pb.CreateMenuItemRequest)(*pb.CreateMenuItemResponse,error){
	r.Logger.Info("Create MenuItem")
	res,err := r.Reservation.CreateMenuItem(menu)
	if err != nil{
		r.Logger.Error("Failed create to menuitem","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) ListMenuItems(ctx context.Context, listMenu *pb.ListMenuItemsRequest)(*pb.ListMenuItemsResponse,error){
	r.Logger.Info("List MenuItem")
	res,err := r.Reservation.ListMenuItems(listMenu)
	if err != nil{
		r.Logger.Error("Failed get menuitems","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService)  GetMenuItem(ctx context.Context,id *pb.GetMenuItemRequest)(*pb.GetMenuItemResponse,error){
	r.Logger.Info("Get MenuItem")
	res,err := r.Reservation.GetMenuItem(id)
	if err != nil{
		r.Logger.Error("Failed get menuitem","error",err.Error())
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) UpdateMenuItem(ctx context.Context, menu *pb.UpdateMenuItemRequest)(*pb.UpdateMenuItemResponse,error){
	r.Logger.Info("Update MenuItem")
	res,err := r.Reservation.UpdateMenuItem(menu)
	if err != nil{
		r.Logger.Error("Failed update menuitem","error",err.Error())
		 return nil,err
	}
	return res,nil
}

func	(r *ReservationService) DeleteMenuItem(ctx context.Context, id *pb.DeleteMenuItemRequest)(*pb.DeleteMenuItemResponse,error){
	r.Logger.Info("Delete MenuItem")
	res,err := r.Reservation.DeleteMenuItem(id)
	if err != nil{
		r.Logger.Error("Failed delete to menuitem","error",err.Error())
		return nil,err
	}
	return res,nil
}