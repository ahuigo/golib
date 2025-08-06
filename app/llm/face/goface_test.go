package face

/*
人脸识别方案
1. go-face(基于dlib 一个强大的C++库，提供人脸检测和识别功能。可以通过CGO调用): 专注人脸。
	- 使用深度卷积神经网络（CNN）
	- 简单直接的工作流程：输入图片 -> 得到特征向量 -> 对比距离。
	- 参考实现案例：https://github.com/AndriyKalashnykov/go-face-recognition/tree/main/models
	- models 下载： https://dlib.net/files/
2. gocv: 基于opencv的Go语言绑定
	- 它不仅支持人脸检测, 还有能进行图像滤波、色彩空间转换、物体跟踪、视频流处理（全栈）
	- 现代深度学习方法 (DNN模块), 可以加载预先训练好的深度学习模型（如来自 TensorFlow, Caffe, PyTorch 的模型
*/
import (
	"fmt"
	"math"
	"testing"

	"github.com/Kagami/go-face"
)

func isSameFace(facePath1, facePath2 string) (bool, error) {
	// 初始化人脸识别器，需要提供模型文件路径
	// 注意：需要下载 dlib 的人脸识别模型文件
	rec, err := face.NewRecognizer("models")
	if err != nil {
		return false, fmt.Errorf("无法初始化人脸识别器: %w", err)
	}
	defer rec.Close()

	// 从第一张图片中识别人脸
	faces1, err := rec.RecognizeFile(facePath1)
	if err != nil {
		return false, fmt.Errorf("无法识别第一张图片中的人脸: %w", err)
	}
	if len(faces1) == 0 {
		return false, fmt.Errorf("第一张图片中未检测到人脸")
	}
	if len(faces1) > 1 {
		return false, fmt.Errorf("第一张图片中检测到多张人脸")
	}

	// 从第二张图片中识别人脸
	faces2, err := rec.RecognizeFile(facePath2)
	if err != nil {
		return false, fmt.Errorf("无法识别第二张图片中的人脸: %w", err)
	}
	if len(faces2) == 0 {
		return false, fmt.Errorf("第二张图片中未检测到人脸")
	}
	if len(faces2) > 1 {
		return false, fmt.Errorf("第二张图片中检测到多张人脸")
	}

	// 计算两个人脸特征向量的欧几里得距离
	distance := euclideanDistance(faces1[0].Descriptor, faces2[0].Descriptor)

	// 设置阈值，通常小于一个值认为是同一个人
	fmt.Println("distance:", distance)
	threshold := 0.5
	return distance < threshold, nil
}

func isRealHumanFace(facePath string) (bool, error) {
	// 初始化人脸识别器
	rec, err := face.NewRecognizer("models")
	if err != nil {
		return false, fmt.Errorf("无法初始化人脸识别器: %w", err)
	}
	defer rec.Close()

	// 检测图片中的人脸
	faces, err := rec.RecognizeFile(facePath)
	if err != nil {
		return false, fmt.Errorf("无法识别图片中的人脸: %w", err)
	}

	// 如果检测到人脸，认为是真人脸（基础实现）
	// 更复杂的活体检测需要额外的模型和算法
	if len(faces) > 0 {
		return true, nil
	}

	return false, nil
}

// euclideanDistance 计算两个特征向量之间的欧几里得距离
func euclideanDistance(a, b face.Descriptor) float64 {
	var sum float64
	for i := range a {
		diff := a[i] - b[i]
		sum += float64(diff) * float64(diff)
	}
	return math.Sqrt(sum)
}
func TestIsSameFace(t *testing.T) {
	p1 := "./pics/lili.jpg"
	p1 = "./pics/leijun.jpg"
	p1 = "./pics/chenshuai1.jpg"
	p2 := "./pics/wanpeng.jpg"
	// p2 = "./pics/lili2.jpg"
	p2 = "./pics/leijun.jpg"
	p2 = "./pics/chenshuai3.jpg"
	p2 = "./pics/glna2.jpg"
	result, err := isSameFace(p1, p2)
	if err != nil {
		t.Fatalf("isSameFace(%s, %s) returned error: %v", p1, p2, err)
	}
	if result {
		t.Fatalf("isSameFace(%s, %s) expected false, got true", p1, p2)
	}
	t.Log("Result of isSameFace:", result)
}

func TestRealHumanFace(t *testing.T) {
	p1 := "/opt/tmp/pics/t1.jpg"
	result, err := isRealHumanFace(p1)
	if err != nil {
		t.Fatalf("isRealHumanFace(%s) returned error: %v", p1, err)
	}
	if !result {
		t.Fatalf("isRealHumanFace(%s) expected true, got false", p1)
	}
	t.Log("Result of isRealHumanFace:", result)
}
