package postgres

import (
	pb "reservation-service/generated/reservation_service"
	"reservation-service/storage/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRestaurant(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("Failed database connection")
		return
	}
	r := redis.ConnectR()
	restaurantRepo := NewRRestaurantRepo(db, r)
	Newrestaurant := pb.CreateRestaurantRequest{
		Name:        "S",
		Address:     "Chilonzor",
		PhoneNumber: "991234567",
		Description: "jwlejf",
	}
	restaurant, err := restaurantRepo.CreateRestaurant(&Newrestaurant)
	if err != nil {
		t.Errorf("Failed Created Restaurant : %v", err)
		return
	}

	// assert.NoError(t, err)

	assert.Equal(t, Newrestaurant.Name, restaurant.Restaurant.Name)
	assert.Equal(t, Newrestaurant.Address, restaurant.Restaurant.Address)
	assert.Equal(t, Newrestaurant.PhoneNumber, restaurant.Restaurant.PhoneNumber)
	assert.Equal(t, Newrestaurant.Description, restaurant.Restaurant.Description)

	assert.NotEmpty(t, restaurant.Restaurant.Id)
}

func TestListRestaurants(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("Failed database connection")
		return
	}
	r := redis.ConnectR()
	restaurantRepo := NewRRestaurantRepo(db, r)
	reqMenu := pb.ListRestaurantsRequest{
		Name:    "S",
		Address: "Chilonzor",
	}
	listRestaurant, err := restaurantRepo.ListRestaurants(&reqMenu)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}
	// assert.NoError(t, err)

	assert.Equal(t, listRestaurant.Restaurants[0].Name, reqMenu.Name)
	assert.Equal(t, listRestaurant.Restaurants[0].Address, reqMenu.Address)

	assert.NotEmpty(t, listRestaurant.Restaurants[0].Id)
}

func TestGetRestaurant(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("Failed database connection")
		return
	}
	r := redis.ConnectR()
	restaurantRepo := NewRRestaurantRepo(db, r)
	id := pb.GetRestaurantRequest{
		Id: "6751a219-6c1c-4676-b2fb-34dac8bfe41a",
	}
	restaurant, err := restaurantRepo.GetRestaurant(&id)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}
	// assert.NoError(t,err)

	assert.NotEmpty(t, restaurant)
}

func TestUpdateRestaurant(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("Failed database connection")
		return
	}
	r := redis.ConnectR()
	restaurantRepo := NewRRestaurantRepo(db, r)
	restaurant := pb.UpdateRestaurantRequest{
		Id:          "207815e3-0b01-46bb-952c-2ea8b8d728e5",
		Name:        "sS",
		Address:     "Qatortol",
		PhoneNumber: "991234546",
		Description: "",
	}
	updateRes, err := restaurantRepo.UpdateRestaurant(&restaurant)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}

	// assert.NoError(t,err)

	assert.Equal(t, restaurant.Name, updateRes.Restaurant.Name)
	assert.Equal(t, restaurant.Address, updateRes.Restaurant.Address)
	assert.Equal(t, restaurant.PhoneNumber, updateRes.Restaurant.PhoneNumber)
	assert.Equal(t, restaurant.Description, updateRes.Restaurant.Description)

	assert.NotEmpty(t, updateRes.Restaurant.Id)
}

func TestDeleteRestaurant(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("Failed database connection")
		return
	}
	r := redis.ConnectR()
	restaurantRepo := NewRRestaurantRepo(db, r)
	id := pb.DeleteRestaurantRequest{
		Id: "207815e3-0b01-46bb-952c-2ea8b8d728e5",
	}
	res, err := restaurantRepo.DeleteRestaurant(&id)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}
	// assert.NoError(t,err)

	assert.NotEmpty(t, res)
}
