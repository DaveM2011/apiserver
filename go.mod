module github.com/davem2011/apiserver

go 1.13

require github.com/caddyserver/caddy/v2 v2.0.0-beta9

//replace github.com/davem2011/apiserver v0.0.0-20200119144700-92f87376b8f4 => ./

replace apiserver => ./

replace apiserver/modules => ./modules
