module example.com/banco

go 1.16

replace example.com/conta => /contas

require (
	example.com/cliente v0.0.0-00010101000000-000000000000 // indirect
	example.com/conta v0.0.0-00010101000000-000000000000
)

replace example.com/cliente => /clientes
