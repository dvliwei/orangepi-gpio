/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  gpio_fac
 * @Version: 1.0.0
 * @Date: 2025/02/18 09:12
 * @Update liwei 2025/2/18 09:12
 */

package gpio

type GpioFactory interface {
	MakeGpio() IsGpio
}

type IsGpio interface {
	//Export 导出指定的 GPIO 引脚
	Export() error

	//Unexport 取消导出指定的 GPIO 引脚
	Unexport(pin int) error

	//SetDirection 设置 GPIO 引脚的方向（输出）
	SetOutDirection() error

	//SetDirection 设置 GPIO 引脚的方向（输入）
	SetInDirection() error

	SetValue(value int) error

	SetOutHigh() error

	SetOutLow() error

	SetInHigh() error

	SetInLow() error

	Cleanup(pins []int)

	Read(pin int) (int, error)

	SetMode(pin int, mode string) error
}
