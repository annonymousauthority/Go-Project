module goproject/usecases

go 1.20

replace goproject/web_crawler/crawler => ./web_crawler

require goproject/web_crawler/crawler v0.0.0-00010101000000-000000000000

require golang.org/x/net v0.6.0 // indirect
