module example.com/conta

go 1.16

replace example.com/cliente => ../clientes

require example.com/cliente v0.0.0-00010101000000-000000000000 // indirect
