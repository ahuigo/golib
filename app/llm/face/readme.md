#  goface 模型
在 models 目录下有4个模型文件：

dlib_face_recognition_resnet_model_v1.dat

    这是 ResNet-based 人脸识别模型
    用于生成128维的人脸特征向量（face descriptor）
    基于深度卷积神经网络（CNN）ResNet架构

mmod_human_face_detector.dat

    这是 Max-Margin Object Detection (MMOD) 人脸检测器
    用于在图片中检测和定位人脸位置
    基于深度学习的现代人脸检测模型

shape_predictor_5_face_landmarks.dat

    5点面部关键点预测器
    检测眼睛中心点和鼻尖等5个关键点
    用于人脸对齐和预处理

shape_predictor_68_face_landmarks.dat

    68点面部关键点预测器
    检测更详细的面部轮廓点（眉毛、眼睛、鼻子、嘴巴、脸部轮廓）
    提供更精确的人脸特征定位
