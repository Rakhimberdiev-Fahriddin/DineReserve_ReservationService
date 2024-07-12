package postgres

import (
	pb "reservation-service/generated/reservation_service"
	"reservation-service/storage/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMenuItem(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	r := redis.ConnectR()
	menuRepo := NewRRestaurantRepo(db, r)

	menu := pb.CreateMenuItemRequest{
		RestaurantId: "a9a9858a-def9-4ab0-9925-a40177cd9b7d",
		Name:         "Osh",
		Description:  "Very good",
		Price:        25000,
	}
	res, err := menuRepo.CreateMenuItem(&menu)
	if err != nil {
		t.Errorf("failed to created menuItem : %v", err)
		return
	}
	// assert.NoError(t, err)

	assert.Equal(t, menu.RestaurantId, res.MenuItem.RestaurantId)
	assert.Equal(t, menu.Name, res.MenuItem.Name)
	assert.Equal(t, menu.Description, res.MenuItem.Description)
	assert.Equal(t, menu.Price, res.MenuItem.Price)

	assert.NotEmpty(t, res.MenuItem.Id)
}

func TestListMenuItems(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	r := redis.ConnectR()
	menuRepo := NewRRestaurantRepo(db, r)
	reqMenu := pb.ListMenuItemsRequest{
		RestaurantId: "6751a219-6c1c-4676-b2fb-34dac8bfe41a",
		Name:         "Osh",
	}
	listMenu, err := menuRepo.ListMenuItems(&reqMenu)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}

	// assert.NoError(t,err)

	assert.Equal(t, reqMenu.RestaurantId, listMenu.MenuItems[0].RestaurantId)
	assert.Equal(t, reqMenu.Name, listMenu.MenuItems[0].Name)

	assert.NotEmpty(t, listMenu.MenuItems[0].Id)
}

func TestGetMenuItem(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	r := redis.ConnectR()
	menuRepo := NewRRestaurantRepo(db, r)
	id := pb.GetMenuItemRequest{
		Id: "be933d44-3822-43f0-bcde-940bfb724dff",
	}
	menu, err := menuRepo.GetMenuItem(&id)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}

	// assert.NoError(t,err)

	assert.NotEmpty(t, menu)
}

func TestUpdateMenuItem(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	r := redis.ConnectR()
	menuRepo := NewRRestaurantRepo(db, r)
	menu := pb.UpdateMenuItemRequest{
		Id:           "be933d44-3822-43f0-bcde-940bfb724dff",
		RestaurantId: "6751a219-6c1c-4676-b2fb-34dac8bfe41a",
		Name:         "Mastava",
		Description:  "dsfa",
		Price:        13000,
	}
	updateMenu, err := menuRepo.UpdateMenuItem(&menu)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}

	assert.Equal(t, menu.Id, updateMenu.MenuItem.Id)
	assert.Equal(t, menu.RestaurantId, updateMenu.MenuItem.RestaurantId)
	assert.Equal(t, menu.Name, updateMenu.MenuItem.Name)
	assert.Equal(t, menu.Description, updateMenu.MenuItem.Description)
	assert.Equal(t, menu.Price, updateMenu.MenuItem.Price)
}

func TestDeleteMenuItem(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	r := redis.ConnectR()
	menuRepo := NewRRestaurantRepo(db, r)
	id := pb.DeleteMenuItemRequest{
		Id: "903cca44-1f9e-487f-9529-ecc06173f042",
	}
	res, err := menuRepo.DeleteMenuItem(&id)
	if err != nil {
		t.Errorf("ERROR : %v", err)
		return
	}

	assert.NotEmpty(t, res)
}
