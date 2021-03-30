module example.com/banco

go 1.16

replace example.com/conta => /contas

replace example.com/cliente => /clientes

require example.com/conta v0.0.0-00010101000000-000000000000
