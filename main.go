package main

import (
	"fmt"
	"os"

	"github.com/hyperform-tech/rtsp-proxy/pkg/proxy"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	version string = "v0.0.0"
)

type args struct {
	version  bool
	confPath string
}

func main() {
	fmt.Println("RTSP proxy")
	kingpin.CommandLine.Help = "rtsp-simple-proxy " + version + "\n\n" +
		"RTSP proxy."

	argVersion := kingpin.Flag("version", "print version").Bool()
	argConfPath := kingpin.Arg("confpath", "path of a config file. The default is conf.yml. Use 'stdin' to read config from stdin").Default("conf.yml").String()

	kingpin.MustParse(kingpin.CommandLine.Parse(os.Args[1:]))

	args := args{
		version:  *argVersion,
		confPath: *argConfPath,
	}

	if args.version == true {
		fmt.Println(version)
		os.Exit(0)
	}

	config, err := proxy.ParseConf("conf.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	program, err := proxy.NewProgram(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	program.Run()
}
