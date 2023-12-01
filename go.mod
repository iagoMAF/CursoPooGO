module banco

go 1.21.4

replace contas => ./pkg/contas

replace clientes => ./pkg/clientes

require contas v0.0.0-00010101000000-000000000000

require clientes v0.0.0-00010101000000-000000000000 // indirect
