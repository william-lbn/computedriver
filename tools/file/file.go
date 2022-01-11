package file

import (
	"io"
	"k8s.io/klog/v2"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		klog.Warningf("%s is not exist and it`s unable to join the cluster. ", path)
		return false
	}
	klog.Warningf("%s is not exist please check it and it`s unable to join the cluster. ", path)
	return false
}

func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		klog.Errorf("file create failed. err: %v", err.Error())
	} else {
		// offset
		// os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, io.SeekEnd)
		_, err = f.WriteAt([]byte(content), n)
		klog.Infof("write succeed!")
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				klog.Errorf("file close error!")
			}
		}(f)
	}
	return err
}
