package main

type address struct {
	version string
	ip      string
	port    string
}

type sodagar struct {
	version string

	http  address
	https address

	uuid string

	sg_type string
}

func main() {

	sodagars := []sodagar{
		sodagar{},
	}

}
