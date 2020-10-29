package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	fmt.Println("vim-go")
	src := os.Args[1]
	dst := os.Args[2]
	fmt.Println(CpFileAtomic(src, dst, 0755))
}

func WriteFileAtomic(filename string, data []byte, perm os.FileMode) error {
	dir, name := path.Split(filename)
	f, err := ioutil.TempFile(dir, name)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err == nil {
		err = f.Sync()
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if permErr := os.Chmod(f.Name(), perm); err == nil {
		err = permErr
	}
	if err == nil {
		err = os.Rename(f.Name(), filename)
	}
	// Any err should result in full cleanup.
	if err != nil {
		os.Remove(f.Name())
	}
	return err
}

func CpFileAtomic(srcFile string, destFile string, perm os.FileMode) error {
	data, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}
	return WriteFileAtomic(destFile, data, perm)
}
func CmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		// timeout
		if err = cmd.Process.Kill(); err != nil {
			glog.Errorf("failed to kill: %s, error: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		glog.Infof("process: `%s %s` killed", cmd.Path, strings.Join(cmd.Args, " "))
		//fmt.Printf("process:`%s %s` killed\n", cmd.Path, strings.Join(cmd.Args, " "))
		return err, true
	case err = <-done:
		return err, false
	}
}

func CmdCopy(src, dest string) (string, error) {
	if !dir.IsExist(src) {
		return "", fmt.Errorf("src [%s] is not exist.", src)
	}
	if dir.IsExist(dest) {
		if err := os.RemoveAll(dest); err != nil {
			return "", err
		}
	}

	cmdStr := fmt.Sprintf("cp -r %s %s", src, dest)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Start() // attention!

	var isTimeout bool
	timeout := 60
	err, isTimeout := CmdRunWithTimeout(cmd, time.Duration(timeout)*time.Second)
	cmdOutput := "stdout:[" + stdout.String() + "]stderr:[" + stderr.String() + "]"

	if err != nil {
		return cmdOutput, err
	}
	if isTimeout {
		return cmdOutput, fmt.Errorf("err: Exec timeout %s", cmdStr)
	}

	return cmdOutput, nil
}
