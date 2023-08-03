// 这个文件夹是专门处理视频封面的
package cover

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

func ossimage(filename string, file string) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	fmt.Println("OSS Go SDK Version: ", oss.Version)
	client, err := oss.New("http://oss-cn-beijing.aliyuncs.com", "LTAI5tSkrGE4gJHrPwZGMJv5", "D5eWGbbYYegGjR70c4zfyGhOYaKTpH")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	//列举所有的存储空间
	marker := ""
	for {
		lsRes, err := client.ListBuckets(oss.Marker(marker))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// 默认情况下一次返回100条记录。
		for _, bucket := range lsRes.Buckets {
			fmt.Println("Bucket: ", bucket.Name)
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	//判断存储空间是否存在
	//yourBucketName存储空间名称
	BucketName := "graphatk1141087952"
	isExist, err := client.IsBucketExist(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("IsBucketExist result : ", isExist)
	//文件上传，文件上传有简单上传，追加上传，断点续传上传，分片上传
	if !isExist {
		os.Exit(-1)
	}
	bucket, err := client.Bucket(BucketName) //注意此处不要写错，写错的话，err让然是nil，我们应该需要先判断一下是否存在
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = bucket.PutObjectFromFile(filename, file) // 这里的
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".jpg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	//names := strings.Split(snapshotPath, "\\")
	snapshotName = snapshotPath + ".jpg"
	fmt.Println(snapshotName)
	ossimage(snapshotName, snapshotName)
	return
}

func Get_cover(filename string, path string, id int64) string {
	photoname, err := GetSnapshot(filename, path, int(id))
	if err != nil {
		fmt.Println(err)
	}
	return photoname
}
