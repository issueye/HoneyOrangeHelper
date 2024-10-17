package helper_cmd

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type RunFunc func(ctx context.Context, isSetWorkSpace bool, path string, args ...string) (chan string, error)

func Run(chanLen int) RunFunc {
	return func(ctx context.Context, isSetWorkSpace bool, path string, args ...string) (chan string, error) {
		cmd := exec.CommandContext(ctx, path, args...)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		// 使用传入进来的path 设置为工作区
		if isSetWorkSpace {
			workSpace := path
			// 判断是文件还是文件夹，如果是文件，则获取其父目录
			if info, err := os.Stat(workSpace); err == nil && !info.IsDir() {
				workSpace = filepath.Dir(workSpace)
			}

			cmd.Dir = workSpace
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return nil, err
		}

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}

		err = cmd.Start()
		if err != nil {
			return nil, err
		}

		if chanLen < 20 {
			chanLen = 20
		}

		msgChannel := make(chan string, chanLen)
		go io.Copy(&appLogWriter{msgChannel: msgChannel}, stderr)
		go io.Copy(&appLogWriter{msgChannel: msgChannel}, stdout)

		return msgChannel, nil
	}
}
