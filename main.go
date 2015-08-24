package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ErikDubbelboer/gspt"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
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
	ProgramName                 string `yaml:":program_name"`
	WaitBetweenSpawningChildren int    `yaml:":wait_between_spawning_children"`
	LogDir                      string `yaml:":log_dir"`
	PidDir                      string `yaml:":pid_dir"`
	SharedDir                   string `yaml:":shared_dir"`
	User                        string `yaml:":user"`
	Children                    int    `yaml:":children"`
	HTTPReadTimeout             int    `yaml:":http_read_timeout"`
	InstanceServiceRegion       string `yaml:":instance_service_region"`
	InstanceServiceEndpoint     string `yaml:":instance_service_endpoint"`
	InstanceServicePort         string `yaml:":instance_service_port"`
	WaitBetweenRuns             int    `yaml:":wait_between_runs"`
	WaitAfterError              int    `yaml:":wait_after_error"`
	CodedeployTestProfile       string `yaml:":codedeploy_test_profile"`
	OnPremisesConfigFile        string `yaml:":on_premises_config_file"`
}

func defaultConfig() Config {
	return Config{
		ProgramName:                 "codedeploy-agent",
		WaitBetweenSpawningChildren: 1,
		Children:                    1,
		HTTPReadTimeout:             80,
		WaitBetweenRuns:             30,
		WaitAfterError:              30,
		CodedeployTestProfile:       "prod",
		OnPremisesConfigFile:        "/etc/codedeploy-agent/conf/codedeploy.onpremises.yml",
	}
}

func readConfig(c *cli.Context) {
	data, err := ioutil.ReadFile(c.GlobalString("config_file"))
	if err != nil {
		log.Fatal("Could not read config file")
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error unmarshaling config file")
	}
}

func setup() {
	gspt.SetProcTitle(config.ProgramName)
}

var (
	config Config
)

func init() {
	config = defaultConfig()
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

	app.Before = func(c *cli.Context) {
		readConfig(c)
		setup()
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
					log.Fatal("AWS CodeDeploy agent is still running")
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
