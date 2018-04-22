package main

import (
	"fmt"

	"github.com/StackExchange/blackbox/pkg/bbutil"
)

func cmdStatus() error {
	bbu, err := bbutil.New()
	if err != nil {
		return err
	}
	names, err := bbu.RegisteredFiles()
	if err != nil {
		return err
	}

	for _, item := range names {
		s := bbutil.FileStatus(bbu.RepoBaseDir, item.Name)
		fmt.Printf("%s\t%s\n", s, item.Name)
	}
	return nil
}
