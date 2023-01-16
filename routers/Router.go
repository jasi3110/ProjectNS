package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	route := mux.NewRouter().StrictSlash(false)
	route = UserRoutes(route)
	route = ProductRoutes(route)
	route = CategoryRoutes(route)
	route = UnitRoutes(route)
	route = RoleRoutes(route)
	return route
}
