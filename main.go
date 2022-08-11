package main

import "github.com/felipefadoni/boilerplate-golang/src/routes"

func main() {

	r := routes.Routes()

	r.Run(":8001")
}
