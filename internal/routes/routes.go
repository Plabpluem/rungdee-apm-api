package routes

import (
	contractAdapters "rungdee-apm-api/internal/adapters/contract"
	customerAdapter "rungdee-apm-api/internal/adapters/customer"
	invoiceAdapter "rungdee-apm-api/internal/adapters/invoice"
	"rungdee-apm-api/internal/adapters/invoice/pdf"
	roomAdapters "rungdee-apm-api/internal/adapters/room"
	userAdapter "rungdee-apm-api/internal/adapters/user"
	contractUsecases "rungdee-apm-api/internal/usecases/contract"
	customeUsecases "rungdee-apm-api/internal/usecases/customer"
	invoiceUsecases "rungdee-apm-api/internal/usecases/invoice"
	roomUsecases "rungdee-apm-api/internal/usecases/room"
	userUsecases "rungdee-apm-api/internal/usecases/user"
	"rungdee-apm-api/pkg/middleware"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	userDb := userAdapter.NewGormUserRepository(db)
	userService := userUsecases.NewUserService(userDb)
	userHttp := userAdapter.NewHttpUserHandler(userService)

	app.Post("/user/login", userHttp.Login)
	app.Post("/user/signup", middleware.AuthRequired, middleware.RbacRequired([]string{"admin"}), userHttp.Signup)
}

func RoomRoutes(app fiber.Router, db *gorm.DB) {
	roomDb := roomAdapters.NewGormRoomRepository(db)
	roomService := roomUsecases.NewRoomService(roomDb)
	roomHttp := roomAdapters.NewHttpRoomHandler(roomService)

	app.Use("/room", middleware.AuthRequired, middleware.RbacRequired([]string{"admin", "employee"}))

	app.Get("/room", roomHttp.FindAll)
	app.Get("/room/dropdown", roomHttp.FindallDropdown)
	app.Get("/room/:id", roomHttp.Find)
	app.Post("/room", roomHttp.Create)
	app.Patch("/room/:id", roomHttp.Update)
}

func CustomerRoutes(app fiber.Router, db *gorm.DB) {
	customerDb := customerAdapter.NewGormCustomerRepository(db)
	customerService := customeUsecases.NewCustomerService(customerDb)
	customerHttp := customerAdapter.NewHttpCustomerHandler(customerService)

	app.Patch("/line-prescreen/:ref", customerHttp.UpdateLinePrescreen)

	app.Use("/customer", middleware.AuthRequired, middleware.RbacRequired([]string{"admin", "employee"}))

	app.Get("/customer", customerHttp.Findall)
	app.Get("/customer/dropdown", customerHttp.FindallDropdown)
	app.Get("/customer/:id", customerHttp.Find)
	app.Post("/customer", customerHttp.Create)
	app.Patch("/customer/:id", customerHttp.Update)
	app.Post("/customer/prescreen", customerHttp.GeneratePrescreen)
}

func ContractRoutes(app fiber.Router, db *gorm.DB) {
	contractDb := contractAdapters.NewGormContractRepository(db)
	contractService := contractUsecases.NewContractService(contractDb)
	contractHttp := contractAdapters.NewHttpContractHandler(contractService)

	app.Use("/contract", middleware.AuthRequired, middleware.RbacRequired([]string{"admin", "employee"}))

	app.Get("/contract", contractHttp.Findall)
	app.Get("/contract/:id", contractHttp.Find)
	app.Post("/contract", contractHttp.Create)
	app.Patch("/contract/:id", contractHttp.Update)

}

func InvoiceRoutes(app fiber.Router, db *gorm.DB) {
	invoiceDb := invoiceAdapter.NewGormInvoiceRepository(db)
	contractDb := contractAdapters.NewGormContractRepository(db)
	pdfGenerate := pdf.NewChromePdfGenerate()

	invoiceService := invoiceUsecases.NewInvoiceService(invoiceDb, contractDb, pdfGenerate)
	invoiceHttp := invoiceAdapter.NewHttpInvoiceHandler(invoiceService)

	app.Use("/invoice", middleware.AuthRequired, middleware.RbacRequired([]string{"admin", "employee"}))

	app.Get("/invoice", invoiceHttp.Findall)
	app.Get("/invoice/:id", invoiceHttp.Find)
	app.Post("/invoice", invoiceHttp.Create)
	app.Patch("/invoice/:id", invoiceHttp.Update)
	app.Post("/invoice/generate/:id", invoiceHttp.GeneratePdf)
}
