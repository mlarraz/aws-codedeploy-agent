package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func startAgent() {

}

func stopAgent() {

}

func restartAgent() {

}

func agentStatus() string {
	return ""
}

type Config struct {
	ProgramName                 string
	WaitBetweenSpawningChildren int
	LogDir                      string
	PidDir                      string
	SharedDir                   string
	User                        string
	Children                    int
	HTTPReadTimeout             int
	InstanceServiceRegion       string
	InstanceServiceEndpoint     string
	InstanceServicePort         string
	WaitBetweenRuns             int
	WaitAfterError              int
	CodedeployTestProfile       string
	OnPremisesConfigFile        string
}

func main() {
	app := cli.NewApp()

	app.Name = "codedeploy-agent"
	app.Usage = "AWS CodeDeploy Agent"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config_file",
			Value: "/etc/codedeploy-agent/conf/codedeployagent.yml",
			Usage: "Path to agent config file",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the AWS CodeDeploy agent",
			Action: func(c *cli.Context) {
				startAgent()
			},
		},
		{
			Name:  "stop",
			Usage: "stop the AWS CodeDeploy agent",
			Action: func(c *cli.Context) {
				stopAgent()

				pid := agentStatus()

				if len(pid) == 0 {
				}
			},
		},
		{
			Name:  "restart",
			Usage: "restart the AWS CodeDeploy agent",
			Action: func(c *cli.Context) {
				restartAgent()
			},
		},
		{
			Name:  "status",
			Usage: "Report running status of the AWS CodeDeploy agent",
			Action: func(c *cli.Context) {
				pid := agentStatus()

				if len(pid) == 0 {
					log.Printf("The AWS CodeDeploy agent is running as PID %s", pid)
				} else {
					log.Fatal("No AWS CodeDeploy agent is running")
				}
			},
		},
	}

	app.Run(os.Args)
}
