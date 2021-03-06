package main

import (
	"fmt"
	"log"

	"github.com/kubaj/kubeauth/providers"
	"gopkg.in/AlecAivazis/survey.v1"
)

var (
	// Version is the current version found in `VERSION` file.
	Version string
	// Build is the output of `git rev-parse HEAD`
	Build string
)

func main() {
	fmt.Printf("version=%s, build=%s\n", Version, Build)
	gcloud := providers.GCloudProvider{}
	kube := providers.KubeConfig{}
	accs, err := gcloud.ReadAccounts()
	if err != nil {
		log.Fatalln(err)
	}
	if len(accs) == 0 {
		if err := gcloud.Authenticate(); err != nil {
			log.Fatalln(err)
		}
		// Now that we've authenticated, try to Re-Read the accounts again and exit if
		// we got an error.
		accs, err = gcloud.ReadAccounts()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if len(accs) > 1 {
		account := ""
		survey.AskOne(
			&survey.Select{
				Message: "Choose an account:",
				Options: accs,
			},
			&account,
			nil,
		)
		gcloud.SelectAccount(account)
	}

	projects, err := gcloud.ReadProjects()
	if err != nil {
		log.Fatalln(err)
	}

	if len(projects) == 0 {
		log.Fatalln("Selected account doesn't have access to any projects")
	}
	project := ""
	survey.AskOne(
		&survey.Select{
			Message: "Choose a project:",
			Options: projects,
		},
		&project,
		nil,
	)

	err = gcloud.SelectProject(project)
	if err != nil {
		log.Fatalln(err)
	}

	clusters, err := gcloud.ReadClusters()
	if err != nil {
		log.Fatalln(err)
	}

	cluster := ""
	survey.AskOne(
		&survey.Select{
			Message: "Choose a cluster:",
			Options: clusters,
		},
		&cluster,
		nil,
	)

	if len(clusters) == 0 {
		log.Fatalln("Selected project doesn't contain any clusters")
	}
	err = gcloud.SelectCluster(cluster)
	if err != nil {
		log.Fatalln(err)
	}

	contexts, err := kube.ReadContexts(cluster)
	if err != nil {
		log.Fatalln(err)
	}

	context := ""
	survey.AskOne(
		&survey.Select{
			Message: "Choose a context:",
			Options: contexts,
		},
		&context,
		nil,
	)

	err = kube.SelectContext(context)
	if err != nil {
		log.Fatalln(err)
	}

}
