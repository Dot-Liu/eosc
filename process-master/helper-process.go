package process_master

import (
	"bytes"
	"errors"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/eolinker/eosc/service"

	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/process"
)

func newHelperProcess(data []byte) (*service.ExtendsResponse, error) {
	cmd, err := process.Cmd(eosc.ProcessHelper, nil)
	if err != nil {
		return nil, err
	}
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	cmd.Wait()
	var body []byte
	_, err = cmd.Stdout.Write(body)
	if err != nil {
		return nil, err
	}
	response := new(service.ExtendsResponse)
	err = proto.Unmarshal(body, response)

	return response, err
}

func checkExtends(exts []*service.ExtendsBasicInfo) ([]*service.ExtendsInfo, []*service.ExtendsBasicInfo, error) {
	request := &service.ExtendsRequest{Extends: make([]*service.ExtendsBasicInfo, 0, len(exts))}
	request.Extends = append(request.Extends, exts...)
	data, err := proto.Marshal(request)
	if err != nil {
		return nil, nil, err
	}
	cmd, err := process.Cmd(eosc.ProcessHelper, nil)
	if err != nil {
		return nil, nil, err
	}
	cmd.Stdin = bytes.NewReader(data)
	buff := &bytes.Buffer{}
	cmd.Stdout = buff
	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}
	cmd.Wait()

	response := new(service.ExtendsResponse)

	err = proto.Unmarshal(buff.Bytes(), response)
	if err != nil {
		return nil, nil, err
	}
	if response.Code != "000000" {
		return nil, nil, errors.New(response.Msg)
	}
	return response.Extends, response.FailExtends, nil
}
