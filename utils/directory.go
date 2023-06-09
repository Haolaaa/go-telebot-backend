package utils

import (
	"admin/global"
	"errors"
	"os"

	"go.uber.org/zap"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("file exists")
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.LOG.Debug("create directory" + v)
			if err := os.Mkdir(v, os.ModePerm); err != nil {
				global.LOG.Error("create directory"+v, zap.Any(" error", err))
				return err
			}
		}
	}
	return err
}
