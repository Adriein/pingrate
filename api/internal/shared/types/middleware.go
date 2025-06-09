package types

type Middleware func(next PingrateHttpHandler) PingrateHttpHandler
