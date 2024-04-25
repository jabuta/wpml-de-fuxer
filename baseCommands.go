package main

import "os"

func exitDefuxer(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
