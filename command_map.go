package main

import "fmt"

func commandMap(cfg *config) error {
	var url *string = nil
	if cfg.next != nil {
		url = cfg.next
	}
	res, err := cfg.client.ListLocations(url)
	if err != nil {
		return err
	}
	cfg.next = res.Next
	cfg.previous = res.Previous

	for _, result := range res.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	var url *string = nil
	if cfg.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	url = cfg.previous
	res, err := cfg.client.ListLocations(url)
	if err != nil {
		return err
	}
	cfg.next = res.Next
	cfg.previous = res.Previous

	for _, result := range res.Results {
		fmt.Println(result.Name)
	}
	return nil
}
