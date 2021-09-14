package eoscli

import (
	"fmt"
	"strconv"

	eosc_args "github.com/eolinker/eosc/eosc-args"
	"github.com/eolinker/eosc/log"
	"github.com/eolinker/eosc/utils"
	"github.com/urfave/cli/v2"
)

func Start(x cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start goku server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "admin-ip",
				Aliases: []string{"ip"},
				Usage:   "ip for the admin process",
				Value:   "0.0.0.0",
			},
			&cli.IntFlag{
				Name:    "admin-port",
				Aliases: []string{"port", "p"},
				Usage:   "port for the admin process",
				Value:   9400,
			},
			&cli.BoolFlag{
				Name:  "join",
				Usage: "join the cluster if true",
			},
			&cli.StringFlag{
				Name:  "broadcast-ip",
				Usage: "ip for the node broadcast, required when join is true",
			},
			&cli.StringFlag{
				Name:  "protocol",
				Usage: "node listen protocol",
				Value: "http",
			},
			&cli.StringSliceFlag{
				Name:    "cluster-addr",
				Aliases: []string{"addr"},
				Usage:   "cluster addr",
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "eosc",
			},
			&cli.StringFlag{
				Name:    "group",
				Aliases: []string{"g"},
				Usage:   "eosc",
			},
		},
		Action: x,
	}
}

//StartFunc 开启节点
func StartFunc(c *cli.Context) error {
	// 判断程序是否存在
	if CheckPIDFILEAlreadyExists() {
		return fmt.Errorf("the app %s is running", eosc_args.AppName())
	}
	ClearPid()
	args := make([]string, 0, 20)
	ip := c.String("ip")
	port := c.Int("port")

	// 从文件中读取cli运行配置
	// 读取存在顺序，若值相同，后读取的会全量覆盖相关配置
	argsName := fmt.Sprintf("%s.args", eosc_args.AppName())
	//nodeName := fmt.Sprintf("%s_node.args", eosc_args.AppName())
	cfg := eosc_args.NewConfig(argsName)
	cfg.ReadFile(argsName)

	err := utils.IsListen(fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return err
	}

	cfg.Set(eosc_args.IP, ip)
	cfg.Set(eosc_args.Port, strconv.Itoa(port))

	protocol := c.String("protocol")
	if protocol == "" {
		protocol = getDefaultArg(cfg, eosc_args.Protocol, "http")
	}
	cfg.Set(eosc_args.Protocol, protocol)

	args = append(args, "start", fmt.Sprintf("--ip=%s", ip), fmt.Sprintf("--port=%d", port))
	cmd, err := StartMaster(args, nil)
	if err != nil {
		log.Errorf("start master error: %w", err)
		return err
	}

	//isJoin := c.Bool("join")
	//if !isJoin {
	//	isJoin, _ = strconv.ParseBool(getDefaultArg(cfg, eosc_args.IsCluster, "false"))
	//}
	//
	//if isJoin {
	//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//	ticket := time.NewTicker(1 * time.Second)
	//	defer ticket.Stop()
	//CheckPid:
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			{
	//				cancel()
	//				return errors.New("join cluster timeout")
	//			}
	//		case <-ticket.C:
	//			{
	//				pid, err := readPid()
	//				if err == nil {
	//					if processExists(pid) {
	//						break CheckPid
	//					}
	//				}
	//			}
	//		}
	//	}
	//	err = join(c, cfg)
	//	if err != nil {
	//		return err
	//	}
	//}
	cfg.Save()

	if eosc_args.IsDebug() {
		return cmd.Wait()
	}
	return nil
}
