# Response Mappers

This package provides mapping logic from internal domain records to response models used in APIs and services.

It is divided into two parts:

- `api/`: Mappers used in the handler/controller layer to transform domain data into API responses.
- `service/`: Mappers used inside service/business logic to produce clean and consistent outputs for other internal services or components.

These mappers ensure a clear separation between internal representations (`record.*`) and external response formats.
