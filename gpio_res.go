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
	pinStr := strconv.Itoa(res.pin)
	err := ioutil.WriteFile("/sys/class/gpio/export", []byte(pinStr), 0644)
	if err != nil {
		if strings.Contains(err.Error(), "Device or resource busy") {
			// 引脚已经导出，忽略错误
			return nil
		}
		return fmt.Errorf("failed to export GPIO pin %d: %w", res.pin, err)
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
