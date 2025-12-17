package util

import (
	"os"
	"path/filepath"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

func GetLastedRmpFile() *os.File {
	yymmdd := time.Now().Format("20060102")
	dirPath := "/root/hl/data/periodic_abci_state_statuses/" + yymmdd
	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	var latestFile string
	for _, file := range files {
		if !file.IsDir() {
			if file.Name() > latestFile {
				latestFile = file.Name()
			}
		}
	}

	if latestFile != "" {
		f, err := os.Open(filepath.Join(dirPath, latestFile))
		if err != nil {
			panic(err)
		}
		return f
	}
	panic("no files found")
}

func ExtractUser() []string {
	decoder := msgpack.NewDecoder(GetLastedRmpFile())

	// 设置允许非字符串键的选项，等同于 Python 的 strict_map_key=False
	decoder.SetMapDecoder(func(d *msgpack.Decoder) (interface{}, error) {
		n, err := d.DecodeMapLen()
		if err != nil {
			return nil, err
		}
		if n == -1 {
			return nil, nil
		}

		m := make(map[interface{}]interface{}, n)
		for i := 0; i < n; i++ {
			k, err := d.DecodeInterface()
			if err != nil {
				return nil, err
			}
			v, err := d.DecodeInterface()
			if err != nil {
				return nil, err
			}
			m[k] = v
		}
		return m, nil
	})

	// 流式解码根对象
	// 首先获取 map 长度
	n, err := decoder.DecodeMapLen()
	if err != nil {
		panic(err)
	}

	if n == -1 {
		panic("Root is not a map")
	}

	var exchangeFound bool

	users := []string{}

	// 遍历 map 中的每个键值对
	for i := 0; i < n; i++ {
		k, err := decoder.DecodeInterface()
		if err != nil {
			panic(err)
		}

		// 检查是否是 "exchange" 键
		if keyStr, ok := k.(string); ok && keyStr == "exchange" {
			exchangeFound = true

			// 对 "exchange" 值进行流式处理
			// 获取 exchange map 长度
			exchangeMapLen, err := decoder.DecodeMapLen()
			if err != nil {
				panic(err)
			}

			if exchangeMapLen == -1 {
				panic("exchange is not a map")
			}

			// 查找 "user_states" 键
			var userStatesFound bool
			for j := 0; j < exchangeMapLen; j++ {
				exchangeKey, err := decoder.DecodeInterface()
				if err != nil {
					panic(err)
				}

				if exchangeKeyStr, ok := exchangeKey.(string); ok && exchangeKeyStr == "user_states" {
					userStatesFound = true

					// 处理 user_states 数组
					arrayLen, err := decoder.DecodeArrayLen()
					if err != nil {
						panic(err)
					}

					if arrayLen == -1 {
						panic("user_states is not an array")
					}

					// 流式遍历数组元素
					for idx := 0; idx < arrayLen; idx++ {
						// 解码数组中的每个元素
						userElementLen, err := decoder.DecodeArrayLen()
						if err != nil {
							panic(err)
						}

						if userElementLen <= 0 {
							continue
						}

						// 只解码第一个元素来检查是否为目标用户
						firstElement, err := decoder.DecodeInterface()
						if err != nil {
							panic(err)
						}

						if firstElementStr, ok := firstElement.(string); ok {
							users = append(users, firstElementStr)
						}

						// 跳过该用户数据的其余部分
						for k := 1; k < userElementLen; k++ {
							err := decoder.Skip()
							if err != nil {
								panic(err)
							}
						}
					}
					break
				} else {
					// 跳过这个键对应的值
					err = decoder.Skip()
					if err != nil {
						panic(err)
					}
				}
			}

			if !userStatesFound {
				panic("user_states not found")
			}
			break
		} else {
			// 跳过这个键对应的值
			err = decoder.Skip()
			if err != nil {
				panic(err)
			}
		}
	}

	if !exchangeFound {
		panic("exchange field not found")
	}

	return users
}
