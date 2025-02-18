/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  gpio
 * @Version: 1.0.0
 * @Date: 2025/02/18 09:12
 * @Update liwei 2025/2/18 09:12
 */

package gpio

type Gpio struct {
	Pin int
}

func NewGpio(pin int) *Gpio {
	return &Gpio{
		Pin: pin,
	}
}

func (p *Gpio) MakeGpio() IsGpio {
	return newGpioRes(p)
}
