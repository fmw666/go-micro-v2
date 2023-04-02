package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	Green   = "\033[32m"
	Red     = "\033[31m"
	GreenBG = "\033[42;37m"
	RedBG   = "\033[41;37m"
	Font    = "\033[0m"
)
const (
	OK    = Green + "[OK]" + Font
	Error = Red + "[Error]" + Font
)

type Script struct {
}

// 打印帮助文档
func (s *Script) Help() {
	fmt.Println("Usage: script <command>")
	fmt.Println("Commands:")
	fmt.Println("  help\t\tShow this help")
	fmt.Println("  init\t\tInit swagger file and go mod")
	fmt.Println("  start\t\tBuild the script")
	fmt.Println("  stop\t\tRun the script")
	fmt.Println("  ls\t\tList screen sessions")
	os.Exit(0)
}

// 初始化功能
func (s *Script) Init() {
	// 打印当前目录下的所有目录
	dirs, _ := ioutil.ReadDir(".")
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("开始初始化 " + dir.Name() + " 服务...")
			// 切换当前工作目录
			os.Chdir(dir.Name())
			if dir.Name() == "api-gateway" {
				// 初始化 swagger 文件
				err := runCmd("swag init")
				if err != nil {
					err = runCmd("go install github.com/swaggo/swag/cmd/swag@latest")
					runCmd("swag init")
				}
			}
			// 初始化 go mod
			runCmd("go mod tidy")
			fmt.Println(OK + " " + GreenBG + dir.Name() + " 服务初始化完成" + Font)
		}
	}
	os.Chdir("..")
}

// 启动所有微服务
func (s *Script) Start() {
	// 得到当前路径
	pwdDir, _ := os.Getwd()
	dirs, _ := ioutil.ReadDir(".")
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("开始启动 " + dir.Name() + " ...")

			serviceDir := filepath.Join(pwdDir, dir.Name())
			mainFile := filepath.Join(serviceDir, "main.go")
			// 判断 main.go 文件是否存在
			if _, err := os.Stat(mainFile); os.IsNotExist(err) {
				fmt.Println(Error + " " + RedBG + dir.Name() + " 服务不存在 main.go 文件" + Font)
				continue
			}

			cmdArgs := []string{"cd", serviceDir, "&&", "go", "run", "main.go"}
			sessionName := dir.Name() + "-service"

			runCmd(fmt.Sprintf("screen -S %s -X quit", sessionName))
			runCmd(fmt.Sprintf("screen -dmS %s", sessionName))
			err := runCmd(fmt.Sprintf("screen -S %s -p 0 -X stuff \"%s\n\"", sessionName, strings.Join(cmdArgs, " ")))
			if err != nil {
				fmt.Println(Error + " " + RedBG + dir.Name() + " 启动失败" + Font)
			} else {
				fmt.Println(OK + " " + GreenBG + dir.Name() + " 启动完成" + Font)
			}
		}
	}
}

// 关闭所有微服务及网关
func (s *Script) Stop() {
	dirs, _ := ioutil.ReadDir(".")
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("开始关闭 " + dir.Name() + " ...")
			sessionName := dir.Name() + "-service"
			runCmd(fmt.Sprint("screen -S " + sessionName + " -X quit"))
			fmt.Println(OK + " " + GreenBG + dir.Name() + " 关闭完成" + Font)
		}
	}
}

// 展示所有的 session
func (s *Script) Ls() {
	runCmd("screen -ls")
}

// 调用 Linux 命令
func runCmd(cmdStr string) error {
	//需要执行的命令
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	// 获取管道输入
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取命令的标准输出管道", err.Error())
		return err
	}
	// 执行Linux命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
		return err
	}
	// 读取所有输出
	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Println("打印异常，请检查")
		return err
	}
	if err := cmd.Wait(); err != nil {
		// fmt.Println("Wait", err.Error())
		return err
	}
	fmt.Printf("%s", bytes)
	return nil
}

// 检查是否有 screen 命令
func checkScreen() {
	cmd := exec.Command("screen", "-v")
	err := cmd.Run()
	if err != nil {
		fmt.Println(Error + " " + Red + "未安装 screen !" + Font)
		os.Exit(1)
	}
}

// 将首字母大写
func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func main() {
	checkScreen()
	args := os.Args
	script := Script{}
	if len(args) == 2 {
		fv := reflect.ValueOf(&script)
		fc := fv.MethodByName(firstUpper(args[1]))
		fc.Call([]reflect.Value{})
	} else {
		script.Help()
	}
}
