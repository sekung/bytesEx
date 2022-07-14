package bytesEx

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"math"
	"time"
)

// Reversed 将源字节切片翻转,内存地址不变
func Reversed(buf []byte) {
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
}

// Reverse 字节翻转，生成一个新字节切片
func Reverse(buf []byte) []byte {
	var newBytes []byte
	for i := len(buf) - 1; i > -1; i-- {
		newBytes = append(newBytes, buf[i])
	}
	return newBytes
}

// Insert 在index处前插入一个字节切片, 生成一个新字节切片
func Insert(buf []byte, index int, b []byte) ([]byte, error) {
	l := len(buf)
	if index > l || index < 0 {
		return nil, errors.New(fmt.Sprintf("insert index out of byt range, excepted 0 - %d", l))
	}
	var newBytes []byte
	newBytes = append(newBytes, buf[:index]...)
	newBytes = append(newBytes, b...)
	newBytes = append(newBytes, buf[index:]...)
	return newBytes, nil
}

// PopInx 弹出字节切片index的值, 生成一个新字节切片
func PopInx(buf []byte, index int) ([]byte, error) {
	l := len(buf)
	if index > l-1 || index < 0 {
		return nil, errors.New(fmt.Sprintf("pop index out of byt range, excepted 0 - %d", l))
	}
	var newBytes []byte
	if index < l-1 {
		newBytes = append(newBytes, buf[:index]...)
		newBytes = append(newBytes, buf[index+1:]...)
	} else {
		newBytes = append(newBytes, buf[:index]...)
	}
	return newBytes, nil
}

// Pop 弹出字节切首位的值, 生成一个新字节切片
func Pop(buf []byte) ([]byte, error) {
	if len(buf) == 0 {
		return nil, errors.New(fmt.Sprintf("pop error ,the byt is nil"))
	}
	return buf[1:], nil
}

// Del 删除切片中的一段类容，生成一个新字节切片
func Del(buf []byte, startIndex int, endIndex int) ([]byte, error) {
	l := len(buf)
	if startIndex < 0 {
		return nil, errors.New("startIndex out of byt range, min is 0")
	} else if startIndex > endIndex {
		return nil, errors.New("startIndex bigger than endIndex")
	} else if endIndex > l {
		return nil, errors.New(fmt.Sprintf("endIndex out of byt range ,max is %d", l))
	}
	var newBytes []byte
	newBytes = append(newBytes, buf[:startIndex]...)
	newBytes = append(newBytes, buf[endIndex:]...)
	return newBytes, nil
}

// Combine 合并多个切片
func Combine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

// Dec 字节拼接取整
func Dec(buf []byte) int {
	l := len(buf)
	sum := 0
	for i := 0; i < len(buf); i++ {
		sum += int(buf[i]) << ((l - i - 1) * 8)
	}
	return sum
}

// Hex 字节切片转16进制字符串
func Hex(buf []byte) string {
	return hex.EncodeToString(buf)
}

// Sum 切片求和
func Sum(buf []byte) int {
	sum := 0
	for i := 0; i < len(buf); i++ {
		sum += int(buf[i])
	}
	return sum
}

// Sum8 切片求Sum8值
func Sum8(buf []byte) byte {
	var sum byte
	for i := 0; i < len(buf); i++ {
		sum += buf[i]
	}
	return sum
}

// CheckSum8 sum8校验
func CheckSum8(buf []byte, target byte) bool {
	return target == Sum8(buf)
}

// CheckSum8Merge sum8校验值追加在切片末尾
func CheckSum8Merge(buf []byte) []byte {
	return Combine(buf, []byte{Sum8(buf)})
}

// Sum16 切片求Sum16值
func Sum16(buf []byte) int {
	sum := 0
	for i := 0; i < len(buf); i++ {
		sum += int(buf[i])
	}
	return sum & 0xFFFF
}

// Sum16Be sum16校验结果按大端显示
func Sum16Be(buf []byte) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(Sum16(buf)))
	return b
}

// Sum16Le sum16校验结果按小端显示
func Sum16Le(buf []byte) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(Sum16(buf)))
	return b
}

// CheckSum16Be sum16大端校验
func CheckSum16Be(buf []byte, target []byte) bool {
	return bytes.Equal(Sum16Be(buf), target)
}

// CheckSum16Le sum16小端校验
func CheckSum16Le(buf []byte, target []byte) bool {
	return bytes.Equal(Sum16Le(buf), target)
}

// CheckSum16BeMerge sum16大端校验并追加在切片末尾
func CheckSum16BeMerge(buf []byte) []byte {
	return Combine(buf, Sum16Be(buf))
}

// CheckSum16LeMerge sum16小端校验
func CheckSum16LeMerge(buf []byte) []byte {
	return Combine(buf, Sum16Le(buf))
}

// CRCModbus 计算CRC校验值返回切片
func CRCModbus(buf []byte) []byte {
	var num1 uint16 = 0xFFFF
	var num2 uint16 = 0xA001
	for _, v := range buf {
		num1 ^= uint16(v)
		for i := 0; i < 8; i++ {
			last := num1 % 2
			num1 >>= 1
			if last == 1 {
				num1 ^= num2
			}
		}
	}
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, num1)
	return b
}

// CheckCRCModbus 进行CRC校验对比
func CheckCRCModbus(buf []byte, target []byte) bool {
	return bytes.Equal(CRCModbus(buf), target)
}

// CheckCRCModbusMerge 将CRC校验值添加在切片末尾
func CheckCRCModbusMerge(buf []byte) []byte {
	return Combine(buf, CRCModbus(buf))
}

// CRCXmodem 计算CRC Xmodem校验值返回切片
func CRCXmodem(buf []byte) []byte {
	var crc16tab = []uint16{
		0x0000, 0x1021, 0x2042, 0x3063, 0x4084, 0x50a5, 0x60c6, 0x70e7,
		0x8108, 0x9129, 0xa14a, 0xb16b, 0xc18c, 0xd1ad, 0xe1ce, 0xf1ef,
	}
	var crc uint16 = 0
	var ch byte = 0
	size := len(buf)
	for i := 0; i < size; i++ {
		ch = byte(crc >> 12)
		crc <<= 4
		crc ^= crc16tab[ch^(buf[i]/16)]
		ch = byte(crc >> 12)
		crc <<= 4
		crc ^= crc16tab[ch^(buf[i]&0x0f)]
	}
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, crc)
	return b
}

// CheckCRCXmodem 进行CRC Xmodem校验对比
func CheckCRCXmodem(buf []byte, target []byte) bool {
	return bytes.Equal(CRCXmodem(buf), target)
}

// CheckCRCXmodemMerge 将CRC Xmodem校验值添加在切片末尾
func CheckCRCXmodemMerge(buf []byte) []byte {
	return Combine(buf, CRCXmodem(buf))
}

// BCC BCC异或和校验
func BCC(buf []byte) byte {
	sum := buf[0]
	for i := 1; i < len(buf); i++ {
		sum ^= buf[i]
	}
	return sum
}

// CheckBCC 进行BCC校验对比
func CheckBCC(buf []byte, target byte) bool {
	return BCC(buf) == target
}

// CheckBCCMerge 将BCC校验值添加在切片末尾
func CheckBCCMerge(buf []byte) []byte {
	return Combine(buf, []byte{BCC(buf)})
}

// Bytes32ToFloatBe 4字节切片转浮点数,大端
func Bytes32ToFloatBe(buf []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(buf))
}

// Bytes32ToFloatLe 4字节切片转浮点数,小端
func Bytes32ToFloatLe(buf []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(buf))
}

// DeBuff 读字节切换进行拆分，返回切片数组
func DeBuff(buf, subPre, subSuf []byte) [][]byte {
	subPreLens := len(subPre)
	subSufLens := len(subSuf)
	sub := append(subSuf, subPre...)
	var rt [][]byte
	if subPreLens != 0 && subSufLens != 0 {
		for {
			if bytes.Index(buf, sub) == -1 {
				rt = append(rt, buf)
				break
			} else {
				index := bytes.Index(buf, sub)
				data := buf[:index+subSufLens]
				rt = append(rt, data)
				buf = buf[index+subSufLens:]
			}
		}
		return rt
	} else if subPreLens != 0 && subSufLens == 0 {
		for {
			i := bytes.Index(buf, subPre)
			if i == -1 {
				rt = append(rt, buf)
				break
			} else if i == 0 {
				inx := bytes.Index(buf[subPreLens:], subPre)
				if inx != -1 {
					rt = append(rt, buf[:inx+subPreLens])
					buf = buf[inx+subPreLens:]
				} else {
					rt = append(rt, buf)
					break
				}
			} else {
				rt = append(rt, buf[:i])
				buf = buf[i:]
			}
		}
		return rt
	} else if subPreLens == 0 && subSufLens != 0 {
		for {
			i := bytes.Index(buf, subSuf)
			if i == -1 {
				rt = append(rt, buf)
				break
			} else {
				rt = append(rt, buf[:i+subSufLens])
				buf = buf[i+subSufLens:]
				if len(buf) == 0 {
					break
				}
			}
		}
		return rt
	} else {
		return [][]byte{buf}
	}
}

// ByteNeighbor3A3BToAB 将相邻的两个字节进行合并
func ByteNeighbor3A3BToAB(buf []byte) ([]byte, error) {
	l := len(buf)
	if l%2 != 0 {
		return nil, errors.New(fmt.Sprintf("The byte lens is %d。must be double ", l))
	}
	var newBytes []byte
	for i := 0; i < l; i += 2 {
		by := buf[i : i+2]
		newBytes = append(newBytes, (by[0]&0x0f)<<4|(by[1]&0x0f))
	}
	return newBytes, nil
}

// Decode 字节解码
//你可以使用GBK、gb18030、gb2312、utf8、utf16le、utf16be, abc 进行解码
func Decode(byt []byte, code string) (string, error) {
	switch code {
	case "GBK", "gbk":
		gbkData, err := simplifiedchinese.GBK.NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(gbkData), nil
		}
	case "GB18030", "gb18030":
		gbkData, err := simplifiedchinese.GB18030.NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(gbkData), nil
		}
	case "gb2312", "GB2312":
		gbkData, err := simplifiedchinese.HZGB2312.NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(gbkData), nil
		}
	case "ABC-BE", "abc-be", "abc", "ABC":
		l := len(byt)
		if l < 2 {
			return "", fmt.Errorf("cannot decode")
		}
		s := ""
		for i := 0; i < l; i = i + 2 {
			if i+1 >= l {
				break
			} else if byt[i] == 0 || byt[i] == 0xff || byt[i+1] == 0 || byt[i+1] == 0xff {
				break
			}
			ss, err := simplifiedchinese.GBK.NewDecoder().Bytes([]byte{byt[i] + 0xA0, byt[i+1] + 0xA0})
			if err != nil {
				break
			}
			s = s + string(ss)
		}
		if len(s) == 0 {
			return "", fmt.Errorf("cannot decode")
		} else {
			return s, nil
		}

	case "ABC-LE", "abc-le":
		l := len(byt)
		if l < 2 {
			return "", fmt.Errorf("cannot decode")
		}
		s := ""
		for i := 0; i < l; i = i + 2 {
			if i+1 >= l {
				break
			} else if byt[i] == 0 || byt[i] == 0xff || byt[i+1] == 0 || byt[i+1] == 0xff {
				break
			}
			ss, err := simplifiedchinese.GBK.NewDecoder().Bytes([]byte{byt[i+1] + 0xA0, byt[i] + 0xA0})
			if err != nil {
				break
			}
			s = s + string(ss)
		}
		if len(s) == 0 {
			return "", fmt.Errorf("cannot decode")
		} else {
			return s, nil
		}
	case "utf8", "utf-8", "UTF8", "UTF-8":
		return string(byt), nil
	case "utf-16-le", "utf16le", "UTF-16-LE", "UTF16LE":
		data, err := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(data), nil
		}
	case "utf-16-be", "utf16be", "UTF-16-BE", "UTF16BE", "UTF16":
		data, err := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(data), nil
		}
	default:
		panic(fmt.Sprintf("unknown decoding: %s", code))
	}
}

// NowTimeBCD 获取当前时间BCD码进行切片显示
func NowTimeBCD() []byte {
	b, _ := hex.DecodeString(time.Now().Format("060102150405"))
	return b
}

// NowTimeYS 获取当前时间YY:MM:DD:HH:MM:SS
func NowTimeYS() []byte {
	now := time.Now()
	return []byte{byte(now.Year() % 100), byte(now.Month()), byte(now.Day()), byte(now.Hour()), byte(now.Minute()), byte(now.Second())}
}

// NowTimeSY 获取当前时间SS:MM:HH:DD:MM:YY
func NowTimeSY() []byte {
	now := time.Now()
	return []byte{byte(now.Second()), byte(now.Minute()), byte(now.Hour()), byte(now.Day()), byte(now.Month()), byte(now.Year() % 100)}
}
