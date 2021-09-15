package container

import (
	"fmt"
	"os"
	"syscall"
)

/*
这里的 init 函数是在容器内部执行的，也就是说 ， 代码执行到这里后 ， 容器所在 的进程其实就已经创建出来了， 这是本容器执行的第 一个进程。
使用 mount 先去挂载 proc 文件系统，以便后面通过 ps 等系统命令去查看当前进程资源 的情况。
*/
func RunContainerInitProcess(command string, args []string) error {
	fmt.Printf("command %s \n", command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		fmt.Println(err)
	}
	return nil
}
