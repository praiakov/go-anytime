module example.com/main

go 1.16

replace example.com/models => /models

replace example.com/db => /db

replace example.com/controllers => /controllers

replace example.com/routes => /routes

require (
	example.com/controllers v0.0.0-00010101000000-000000000000 // indirect
	example.com/db v0.0.0-00010101000000-000000000000 // indirect
	example.com/models v0.0.0-00010101000000-000000000000 // indirect
	example.com/routes v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.0 // indirect
)
