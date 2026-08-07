# Items API

Objetivo del ejercicio final:

- completar el `Repositorio`
- completar el `Service`
- completar los `Controllers`
- usar MongoDB como persistencia

## Sugerencia de estructura

```text
items-api/
  dao/
    items/
      items_dao.go
  repositories/
    items/
      items_mock.go
      items_mongo.go
  services/
    items/
      items_service.go
      items_service_test.go
  controllers/
    items/
      items_controller.go
  main.go
```

## Operaciones mínimas sugeridas

- `CreateItem`
- `GetItemByID`
- `ListItems`
- `DeleteItemByID`
