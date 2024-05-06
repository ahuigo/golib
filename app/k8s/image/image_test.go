package images

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"flag"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetNodes(t *testing.T) {
	kubeconfig := flag.String("kubeconfig", "/path/to/your/kubeconfig", "kubeconfig file")
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, _ := kubernetes.NewForConfig(config)

	node, _ := clientset.CoreV1().Nodes().Get(context.Background(), "your-node-name", metav1.GetOptions{})
	for _, image := range node.Status.Images {
		for _, name := range image.Names {
			fmt.Println(name)
		}
	}
}

var cli *client.Client

func init() {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
	if err != nil {
		panic(err)
	}

}

func TestImageLayers(t *testing.T) {
	// cli, err := client.NewClientWithOpts(client.FromEnv) //创建了一个新的 Docker 客户端
	//1. get history layers of image
	const imageID = "45d951d3fc77"
	history, err := cli.ImageHistory(context.Background(), imageID)
	if err != nil {
		log.Fatal(err)
	}
	for _, layer := range history {
		size, err := strconv.ParseFloat(fmt.Sprintf("%d", layer.Size), 32)
		if err != nil {
			log.Fatal(err)
		}
		t.Log("size: ", size, "\tshaid: ", layer.ID)
	}
}

func TestImageList(t *testing.T) {
	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, image := range images {
		fmt.Println("ID: ", image.ID, "RepoTags: ", image.RepoTags, image.Labels)
	}

}
