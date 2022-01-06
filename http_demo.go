package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	mv2ray "main/miniv2ray"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"

	"log"
	"net/http"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

const (
	accessKey    = ""
	secretKey    = ""
	region       = "ap-southeast-1"
	instanceName = "CentOS-1-V2ray"
)

func main() {
	fc.StartHttp(HandleHttpRequest)
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")

	w.Write([]byte(getRespWithUrl(req) + "\r\n"))
	return nil
}

func getRespWithUrl(req *http.Request) string {
	if strings.Contains(req.URL.String(), "instance") {
		creds := credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
			},
		}
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithCredentialsProvider(creds))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}
		svc := lightsail.NewFromConfig(cfg)
		resp, err := svc.GetInstance(context.TODO(), &lightsail.GetInstanceInput{
			InstanceName: aws.String(instanceName),
		})

		if err != nil {
			log.Fatalf("failed to list tables, %v", err)
		}
		instance := resp.Instance
		vmess, err := json.Marshal(map[string]string{
			"v":    "2",
			"ps":   "aws-tcp",
			"add":  *instance.PublicIpAddress,
			"port": "3306",
			"id":   "5d4893a0-18d5-11eb-a501-029405bb920e",
			"aid":  "0",
			"net":  "tcp",
			"type": "none",
			"host": "",
			"path": "",
			"tls":  "",
		})
		result := map[string]string{"Arn": *instance.Arn, "BlueprintId": *instance.BlueprintId,
			"BlueprintName":   *instance.BlueprintName,
			"PublicIpAddress": *instance.PublicIpAddress,
			"State":           *instance.State.Name, "instanceName": *instance.Name,
			"vmess": base64.StdEncoding.EncodeToString(vmess)}
		return MapToJson(result)
	}

	if strings.Contains(req.URL.String(), "vmess") {
		vmess := "vmess://" + req.FormValue("vmess")
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
		ps, _ := Ping(vmess, uint(3), "http://www.google.com/gen_204", uint(3), uint(1), uint(0), osSignals, false, false, false)
		result := map[string]string{"vmess": vmess, "counter": strconv.Itoa(int(ps.ReqCounter)), "success": strconv.Itoa(len(ps.Delays))}
		return MapToJson(result)
	}
	return req.FormValue("key")
}
func MapToJson(param map[string]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func Ping(vmess string, count uint, dest string, timeoutsec, inteval, quit uint, stopCh <-chan os.Signal, showNode, verbose, usemux bool) (*PingStat, error) {
	server, err := mv2ray.StartV2Ray(vmess, verbose, usemux)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if err := server.Start(); err != nil {
		fmt.Println("Failed to start", err)
		return nil, err
	}
	defer server.Close()

	if showNode {
		go func() {
			info, err := mv2ray.GetNodeInfo(server, time.Second*10)
			if err != nil {
				return
			}

			fmt.Printf("Node Outbound: %s/%s\n", info["loc"], info["ip"])
		}()
	}

	ps := &PingStat{}
	ps.StartTime = time.Now()
	round := count
L:
	for round > 0 {
		seq := count - round + 1
		ps.ReqCounter++

		chDelay := make(chan int64)
		go func() {
			delay, err := mv2ray.MeasureDelay(server, time.Second*time.Duration(timeoutsec), dest)
			if err != nil {
				ps.ErrCounter++
				fmt.Printf("Ping %s: seq=%d err %v\n", dest, seq, err)
			}
			chDelay <- delay
		}()

		select {
		case delay := <-chDelay:
			if delay > 0 {
				ps.Delays = append(ps.Delays, delay)
				fmt.Printf("Ping %s: seq=%d time=%d ms\n", dest, seq, delay)
			}
		case <-stopCh:
			break L
		}

		if quit > 0 && ps.ErrCounter >= quit {
			break
		}

		if round--; round > 0 {
			select {
			case <-time.After(time.Second * time.Duration(inteval)):
				continue
			case <-stopCh:
				break L
			}
		}
	}

	ps.CalStats()
	return ps, nil
}

func (p *PingStat) CalStats() {
	for _, v := range p.Delays {
		p.SumMs += uint(v)
		if p.MaxMs == 0 || p.MinMs == 0 {
			p.MaxMs = uint(v)
			p.MinMs = uint(v)
		}
		if uv := uint(v); uv > p.MaxMs {
			p.MaxMs = uv
		}
		if uv := uint(v); uv < p.MinMs {
			p.MinMs = uv
		}
	}
	if len(p.Delays) > 0 {
		p.AvgMs = uint(float64(p.SumMs) / float64(len(p.Delays)))
	}
}

type PingStat struct {
	StartTime  time.Time
	SumMs      uint
	MaxMs      uint
	MinMs      uint
	AvgMs      uint
	Delays     []int64
	ReqCounter uint
	ErrCounter uint
}
