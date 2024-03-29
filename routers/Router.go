package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	route := mux.NewRouter().StrictSlash(false)
	route = UserRoutes(route)
	route = ProductRoutes(route)
	route = CategoryRoutes(route)
	route = UnitRoutes(route)
	route = RoleRoutes(route)
	route = PricesRoutes(route)
	route = SaleRoutes(route)
	route = UserAddressRoutes(route)
	route = DiscountRoutes(route)
	route = CartRoutes(route)
	route= DashBoardRoutes(route)
	route=AdImagesRoutes(route)
	return route
}
