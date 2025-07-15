package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	b := make(map[int]models.Book)
	m := make(map[int]models.Member)
	lib := services.Library{Books : b,Members : m}
	user_choice := 1
	for user_choice != 7{
		user_choice = controllers.User_options()
		services.Service(user_choice, lib)
	}
}