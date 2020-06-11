package thread

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
)

type Envs map[string]string

type Runner struct {
	Name        string
	Path        string
	Args        []string
	WorkingDir  string
	Envs        []string
	Output      string
	proc        *os.Process
	killManual  bool
	killSignal  chan interface{}
	exitSignal  chan *os.ProcessState
	initialized bool
}

func (r *Runner) fork(from string) error {
	var (
		flag = os.O_WRONLY | os.O_APPEND | os.O_CREATE
		path = fmt.Sprintf("/mnt/sdcard/log/data/daemon-%s-%v.log",
			r.Name, time.Now().Format("20060102"))
		stdout, stderr = os.Stdout, os.Stderr
		file, err      = os.OpenFile(path, flag, 0666)
		pid            int
	)
	if err == nil {
		stdout, stderr = file, file
	}

	pid, err = syscall.ForkExec(r.Path, r.Args, &syscall.ProcAttr{
		Dir:   r.WorkingDir,
		Env:   append(append(os.Environ(), "forked="+from), r.Envs...),
		Files: []uintptr{0, stdout.Fd(), stderr.Fd()},
		Sys:   nil,
	})
	if err != nil {
		return err
	}
	r.proc, err = os.FindProcess(pid)
	if err != nil {
		return err
	}
	if r.proc == nil {
		return errors.New("process context is nil")
	}
	go r.wait()
	return nil
}

func (r *Runner) init() {
	if !r.initialized {
		r.killSignal = make(chan interface{}, 1)
		r.exitSignal = make(chan *os.ProcessState, 1)
		r.Args = append([]string{r.Path}, r.Args...)
		r.initialized = true
	}
	r.killManual = false
}

func (r *Runner) wait() {
	var result, _ = r.proc.Wait()
	r.exitSignal <- result
}

func (r *Runner) kill() {
	if r.proc != nil {
		_ = syscall.Kill(r.proc.Pid, syscall.SIGKILL)
	}
}

func (r *Runner) Stop() {
	r.killManual = true
	r.kill()
	<-r.killSignal
	return
}

func (r *Runner) Start() {
	r.init()

	var err error

	err = r.clean()
	if err != nil {
		fmt.Printf("%s: kill previous failed: %v", r.Name, err)
	}

	if err = r.fork("A"); err != nil {
		fmt.Printf("%s: fork exec failed: %v", r.Name, err)
		return
	}
	fmt.Printf("%s: process started: %d", r.Name, r.proc.Pid)

	go func() {
		for {
			select {
			case v := <-r.exitSignal:
				fmt.Printf("monitor %s exit code: %d", r.Name, v.ExitCode())
				if r.killManual {
					fmt.Printf("monitor %s: manual killed", r.Name)
					r.killSignal <- 1
					return
				}
				if err = r.fork("A"); err != nil {
					fmt.Printf("%s: fork exec failed: %v", r.Name, err)
					r.killSignal <- 1
					return
				}
				fmt.Printf("monitor %s: process started: %d", r.Name, r.proc.Pid)
				continue
			}
		}
	}()
}

func (r *Runner) Restart() error {
	var err error

	r.Stop()
	time.Sleep(5 * time.Second)
	if err = r.fork("B"); err != nil {
		return err
	}

	select {
	case v := <-r.exitSignal:
		return fmt.Errorf("restart %s exit code: %d", r.Name, v.ExitCode())
	case <-time.After(30 * time.Second):
		r.kill()
		<-r.exitSignal
		r.Start()
		return nil
	}
}

func (r *Runner) clean() error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, v := range processes {
		name, _ := v.Name()
		if strings.Contains(name, r.Name) {
			if err = syscall.Kill(int(v.Pid), syscall.SIGKILL); err != nil {
				return err
			}
		}
	}
	return nil
}
