package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	err := SaveData1("output/database/test.data", []byte("Hello, World!"))
	if err == nil {
		err := SaveData2("output/database/test", []byte("Hello, World!"))
		if err != nil {
			return
		}
	}
}

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {

		}
	}(fp)

	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return fp.Sync() // fsync
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer func() {
		err = fp.Close()
		if err != nil {
			err := os.Remove(tmp)
			if err != nil {
				return
			}
		}
	}()
	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	err = fp.Sync() // fsync
	if err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func randomInt() any {
	return rand.IntN(100)
}
