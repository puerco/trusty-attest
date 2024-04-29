package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/stacklok/trusty-attest/internal/packages"
	"github.com/stacklok/trusty-attest/internal/trusty"
)

func main() {
	l := packages.NewLister()
	ctx := context.Background()
	if len(os.Args) == 1 {
		logrus.Fatal("no directory specified")
	}

	nodelist, err := l.ReadPackages(ctx, os.Args[1])
	if err != nil {
		logrus.Fatal(err)
	}

	tclient := trusty.NewClient()

	for _, n := range nodelist.Nodes {
		if n.Id == "root" {
			continue
		}
		fmt.Printf("package: %s version: %s\n", n.Name, n.Version)
		score, err := tclient.NodeScore(n)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("+%v", score)
		break
	}

}
