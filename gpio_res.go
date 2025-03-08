/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  gpio_res
 * @Version: 1.0.0
 * @Date: 2025/02/18 09:12
 * @Update liwei 2025/2/18 09:12
 */

package gpio

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type gpioRes struct {
	pin int
}

func newGpioRes(pin *Gpio) *gpioRes {
	return &gpioRes{
		pin: pin.Pin,
	}
}

func (res *gpioRes) Export() error {
	filePath := fmt.Sprintf("/sys/class/gpio/gpio%d/value", res.pin)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		pinStr := strconv.Itoa(res.pin)
		err := ioutil.WriteFile("/sys/class/gpio/export", []byte(pinStr), 0644)
		if err != nil {
			if strings.Contains(err.Error(), "Device or resource busy") {
				// 引脚已经导出，忽略错误
				return nil
			}
			return fmt.Errorf("failed to export GPIO pin %d: %w", res.pin, err)
		}
	}

	return nil
}

func (res *gpioRes) Unexport(pin int) error {
	pinStr := strconv.Itoa(pin)
	err := ioutil.WriteFile("/sys/class/gpio/unexport", []byte(pinStr), 0644)
	if err != nil {
		return fmt.Errorf("failed to unexport GPIO pin %d: %w", pin, err)
	}
	return nil
}

func (res *gpioRes) SetOutDirection() error {
	direction := "out"
	pinStr := strconv.Itoa(res.pin)
	path := fmt.Sprintf("/sys/class/gpio/gpio%s/direction", pinStr)
	err := ioutil.WriteFile(path, []byte(direction), 0644)
	if err != nil {
		return fmt.Errorf("failed to set direction of GPIO pin %d: %w", res.pin, err)
	}
	return nil
}

func (res *gpioRes) SetInDirection() error {
	direction := "in"
	pinStr := strconv.Itoa(res.pin)
	path := fmt.Sprintf("/sys/class/gpio/gpio%s/direction", pinStr)
	err := ioutil.WriteFile(path, []byte(direction), 0644)
	if err != nil {
		return fmt.Errorf("failed to set direction of GPIO pin %d: %w", res.pin, err)
	}
	return nil
}

func (res *gpioRes) SetValue(value int) error {
	pinStr := strconv.Itoa(res.pin)
	path := fmt.Sprintf("/sys/class/gpio/gpio%s/value", pinStr)
	valueStr := strconv.Itoa(value)
	err := ioutil.WriteFile(path, []byte(valueStr), 0644)
	if err != nil {
		return fmt.Errorf("failed to set value of GPIO pin %d: %w", res.pin, err)
	}
	return nil
}

func (res *gpioRes) SetHigh() error {
	err := res.Export()
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	err = res.SetOutDirection()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	return res.SetValue(1)
}

func (res *gpioRes) SetLow() error {
	err := res.Export()
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	err = res.SetOutDirection()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	return res.SetValue(0)
}

func (res *gpioRes) Cleanup(pins []int) {
	for _, pin := range pins {
		if err := res.Unexport(pin); err != nil {
			fmt.Printf("Unexpected error unexporting pin %d: %v\n", pin, err)
			continue
		}
	}
}

func (res *gpioRes) Read(pin int) (int, error) {
	//检查 GPIO 引脚是否已经导出
	err := res.Export()
	if err != nil {
		return 0, err
	}
	filePath := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read value from GPIO pin %d: %w", pin, err)
	}
	value, err := strconv.Atoi(string(data[:1]))
	if err != nil {
		return 0, fmt.Errorf("failed to convert value from GPIO pin %d: %w", pin, err)
	}
	return value, nil
}

func (res *gpioRes) SetMode(pin int, mode string) error {
	filePath := fmt.Sprintf("/sys/class/gpio/gpio%d/active_low", pin)
	// 这里只是示例，不同系统对模式设置的文件和值可能不同
	// 实际中可能需要根据具体硬件和系统修改
	var value string
	switch mode {
	case "UP":
		value = "0"
	case "DOWN":
		value = "1"
	default:
		return fmt.Errorf("unsupported GPIO mode: %s", mode)
	}
	err := ioutil.WriteFile(filePath, []byte(value), 0644)
	if err != nil {
		return fmt.Errorf("failed to set mode for GPIO pin %d: %w", pin, err)
	}
	return nil
}
