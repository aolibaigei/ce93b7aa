package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	//BASE INFO
	info, err := client.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println("容器总数:", info.Containers)
	fmt.Println("正在使用中:", info.ContainersRunning)
	fmt.Println("运行暂停中:", info.ContainersPaused)
	fmt.Println("运行停止:", info.ContainersStopped)
	fmt.Println("容器镜像总数:", info.Images)
	fmt.Println("操作系统:", info.OperatingSystem)
	fmt.Println("内核版本:", info.KernelVersion)
	fmt.Println("操作系统类型:", info.OSType)
	fmt.Println("操作系统体系架构:", info.Architecture)
	fmt.Println("Docker根目录:", info.DockerRootDir)
	fmt.Println("Docker版本:", info.ServerVersion)
	fmt.Println("Dokcer运行时:", info.DefaultRuntime)
	fmt.Println("Registry:", info.IndexServerAddress)
	fmt.Println("Runtimes:", info.Runtimes)
	fmt.Println("标签:", info.Labels)
	fmt.Println("标签:", info.Name)
	//
	//
	//
	//Container info
	// containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	// if err != nil {
	// 	panic(err)
	// }

	// for _, b := range containers {
	// 	fmt.Println("容器ID:", b.ID)
	// 	fmt.Println("使用的镜像名称:", b.Image)
	// 	fmt.Println("容器名称:", b.Names)
	// 	fmt.Println("容器创建时间:", b.Created)
	// 	fmt.Println("容器网络配置:", b.Networks)
	// 	fmt.Println("占用端口号:", b.Ports)
	// 	fmt.Println("运行状态:", b.Status)
	// 	fmt.Println(b.State)
	// 	fmt.Println(b.Mounts)
	// 	fmt.Println(b.Labels)
	// 	fmt.Println(b.Command)
	// }
	//
	//image info
	// imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	// if err != nil {
	// 	panic(err)
	// }
	// for _, img := range imgs {
	// 	fmt.Println("ID: ", img.ID)
	// 	fmt.Println("RepoTags: ", img.RepoTags)
	// 	fmt.Println("Created: ", img.Created)
	// 	fmt.Println("Size: ", img.Size)
	// 	fmt.Println("VirtualSize: ", img.VirtualSize)
	// 	fmt.Println("ParentId: ", img.ParentID)
	// }
	// Networks_info, err = client.ListNetworks()

	// fmt.Println(Networks_info)

}
