# orangepi-gpio
orangepi.gpio action
````azure
   //Export 导出指定的 GPIO 引脚
     Export() error

	//Unexport 取消导出指定的 GPIO 引脚
	Unexport(pin int) error

	//SetDirection 设置 GPIO 引脚的方向（输出）
	SetOutDirection() error

	//SetDirection 设置 GPIO 引脚的方向（输入）
	SetInDirection() error

	SetValue(value int) error

	SetHigh() error

	SetLow() error

	Cleanup(pins []int)

````

### DEMO

```azure
func main(){
    pin := 123
	gpioRes := gpio.NewGpio(pin)
    if err := gpioRes.MakeGpio().SetHigh(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("gpio set high success")
}
```
