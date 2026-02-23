# Rungdee APM API - Workflow & Task Tracker
# อัพเดทล่าสุด: 23/02/2026

================================================================================
## สถานะแต่ละ Module
================================================================================

### USER MODULE                                                    [DONE]
  Entity            : user.go
  UseCase           : user_usecases.go
  Repository I/F    : user_repository.go
  GORM Adapter      : gorm_user_repository.go
  HTTP Handler      : http_user_repository.go
  DTOs              : signup, login, find-user, update-user
  Routes            : POST /user/login, POST /user/signup
  ---
  CRUD: Create(Signup), FindAll, Find, Update, Login, FindByUsername

### ROOM MODULE                                                    [DONE]
  Entity            : room.go
  UseCase           : room_usecase.go
  Repository I/F    : room_repository.go
  GORM Adapter      : gorm_room_repository.go
  HTTP Handler      : http_room_repository.go
  DTOs              : create_room, find_room, update_room
  Routes            : GET /room, GET /room/:id, POST /room, PATCH /room
  ---
  CRUD: Create, FindAll, Find, Update

### CUSTOMER MODULE                                                [90%]
  Entity            : customer.go
  UseCase           : customer_usecase.go
  Repository I/F    : customer_repository.go
  GORM Adapter      : gorm_customer_repository.go
  HTTP Handler      : http_customer_repository.go
  DTOs              : create, find, filter, update
  Response          : customer_response.go (pagination)
  Routes            : CustomerRoutes() defined BUT NOT registered in main.go  <-- FIX
  ---
  CRUD: Create, Findall, Find, Update

### CONTRACT MODULE                                                [70%]
  Entity            : contract.go
  UseCase           : contract_usecase.go
  Repository I/F    : contract_repository.go
  GORM Adapter      : gorm_contract_repository.go (+FindByUuid, +FindById)
  HTTP Handler      : http_contract.repository.go
  DTOs              : create, find, filter, update
  Response          : contract_response.go (pagination)
  Routes            : NOT defined in routes.go                     <-- TODO
  ---
  CRUD: Create, Findall, Findone, Update, FindByUuid, FindById

### INVOICE MODULE                                                 [50%]
  Entity            : invoice.go
  UseCase           : invoice_usecase.go (business logic done)
  Repository I/F    : invoice_repository.go (+ContractReader)
  GORM Adapter      : gorm_invoice_repository.go
  HTTP Handler      : NOT IMPLEMENTED                              <-- TODO
  DTOs              : create, find, filter, update
  Response          : invoice_response.go (pagination)
  Routes            : NOT defined in routes.go                     <-- TODO
  ---
  CRUD: Create, Findall, Find, Update

================================================================================
## Middleware & Security
================================================================================

  Auth Middleware    : pkg/middleware/authentication.go              [DONE]
  RBAC Middleware   : pkg/middleware/rbac.go                         [DONE]
  Pagination        : pkg/pagination.go                             [DONE]

================================================================================
## TODO - งานที่ยังต้องทำ (เรียงตามลำดับ)
================================================================================

### Priority 1 - Integration (ทำให้ระบบใช้งานได้)
-----------------------------------------------
[ ] 1.1 main.go           : เพิ่ม routes.CustomerRoutes(api, db)
[ ] 1.2 routes.go         : เพิ่ม ContractRoutes() function + register ใน main.go
[ ] 1.3 routes.go         : เพิ่ม InvoiceRoutes() function + register ใน main.go
[ ] 1.4 invoice adapter   : สร้าง http_invoice_handler.go (Create, Findall, Find, Update)
[ ] 1.5 invoice routes    : wire InvoiceService(invoiceDb, contractDb) ส่ง 2 repos

### Priority 2 - RBAC & Security
-----------------------------------------------
[ ] 2.1 routes.go         : Signup ต้องเพิ่ม AuthRequired + RbacRequired(admin)
[ ] 2.2 routes.go         : Room POST/PATCH ต้องแยก RBAC admin only (ตอนนี้ employee ก็สร้างได้)
[ ] 2.3 rbac.go           : เปลี่ยน .JSON("string") -> .JSON(fiber.Map{"error": "..."})
[ ] 2.4 authentication.go : เปลี่ยน SendStatus -> .JSON(fiber.Map{"error": "unauthorized"})

### Priority 3 - Bug Fixes
-----------------------------------------------
[ ] 3.1 invoice_usecase.go:14  : NewInvoiceService ไม่ได้ assign contractRepo
      แก้: return &InvoiceService{repo: repo, contractRepo: contractRepo}
[ ] 3.2 invoice_usecase.go:54  : TotalAmount ใช้ค่าสลับ (water*elec, elec*water)
      แก้: WaterPerUnit * water_unit, ElecPerUnit * elec_unit
[ ] 3.3 invoice_usecase.go:100 : TotalAmount เดียวกัน ใน Update
[ ] 3.4 gorm_invoice_repo:29   : return dto, err -> return dto, nil
[ ] 3.5 gorm_contract_repo:76  : return &contract, err -> return &contract, nil
[ ] 3.6 gorm_contract_repo:115 : return &contract, err -> return &contract, nil
[ ] 3.7 gorm_contract_repo:129 : return &contract, err -> return &contract, nil

### Priority 4 - Naming Convention
-----------------------------------------------
[ ] 4.1 Rename: http_user_repository.go     -> http_user_handler.go
[ ] 4.2 Rename: http_room_repository.go     -> http_room_handler.go
[ ] 4.3 Rename: http_customer_repository.go -> http_customer_handler.go
[ ] 4.4 Rename: http_contract.repository.go -> http_contract_handler.go

### Priority 5 - Future Features
-----------------------------------------------
[ ] 5.1 PDF Generation   : สร้าง endpoint GET /contract/:uuid/pdf
[ ] 5.2 Line Notify      : ส่ง invoice ผ่าน Line Messaging API
[ ] 5.3 Refresh Token     : JWT access 15min + refresh token

================================================================================
## Clean Architecture Checklist
================================================================================

  [x] Entity ไม่ import framework (ยกเว้น GORM tags - acceptable trade-off)
  [x] UseCase interface อยู่ใน usecase layer
  [x] Dependency injection ผ่าน constructor
  [x] DTO แยกจาก Entity
  [x] Business logic อยู่ใน UseCase (JWT, bcrypt, invoice calc)
  [x] HTTP handler ทำแค่ parse -> call usecase -> return
  [x] ContractReader interface อยู่ใน invoice package (ISP)
  [ ] Response struct อยู่ใน adapter layer
  [ ] Naming convention สม่ำเสมอ (handler vs repository)
  [ ] RBAC ครอบคลุมทุก endpoint
