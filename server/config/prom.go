package config

import "fmt"

type Prom struct {
	Services []Service `yaml:"services"`
	Day      string    `yaml:"day"`
}

type Service struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}

func (pc *Prom) String() {
	fmt.Println("PromConfig:")
	fmt.Println("  Services:")
	for _, value := range pc.Services {
		fmt.Println("Name:", value.Name)
		fmt.Println("address:", value.Address)
		fmt.Println("account:", value.Account)
		fmt.Println("password:", value.Password)
	}
	fmt.Println("day:", pc.Day)

}
