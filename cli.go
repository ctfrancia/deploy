package main

import (
	"flag"
)

type serviceBranch struct {
	serviceName string
	branch      string
	test        bool
}

func newServiceBranch() *serviceBranch {
	sb := new(serviceBranch)

	flag.StringVar(&sb.serviceName, "s", "", "The service that you would like to be deployed")
	flag.StringVar(&sb.branch, "b", "", "The branch that you would like to deploy (master, uat, staging)")
	flag.BoolVar(&sb.test, "t", false, "Whether or not the tests should be ran for the project")

	flag.Parse()

	return sb
}
