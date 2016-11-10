//Create by sndnvaps <sndnvaps@gmail.com>

package main

import (
    "github.com/andlabs/ui"
)


// paramer :val string; 输入18位的身份证号码
//         w *ui.Window ; 主界面函数调用
// return is_valid; 返回身份证验证结果 
func getInputVal(w *ui.Window, val string) string {
    var id_card [18]byte // 'X' == byte(88)， 'X'在byte中表示为88

    if len(val) != 18 {
        ui.MsgBox(w, "输入错误！", "必须要输入18位的身份证号码!\n" + "你输入的号码为 = " + val)
        return "error"
    }

    // 将字符串，转换成[]byte,并保存到id_card[]数组当中
    for k, v := range []byte(val) {
        id_card[k] = byte(v)
    }

    //把身份证前17位，复制到id_card_copy里面
    var id_card_copy [17]byte
    for j := 0; j < 17; j++ {
        id_card_copy[j] = id_card[j]
    }

   _, is_valid :=  verify_id(check_id(id_card_copy), byte2int(id_card[17]))
   return is_valid

}

func createMainWindow() {
    err := ui.Main(func() {
        windows := ui.NewWindow("验证身份证号码正确性", 150, 300, false)
        id_card_num_input := ui.NewEntry()
        button := ui.NewButton("检测")
        output := ui.NewLabel("验证结果:")
        box := ui.NewVerticalBox()
        box.Append(id_card_num_input, false)
        box.Append(button, false)
        box.Append(output, false)

        windows.SetChild(box)
        button.OnClicked(func(*ui.Button) {
            check_valid :=  getInputVal(windows, id_card_num_input.Text())
            output.SetText("验证结果:" +  check_valid )
        })
        windows.OnClosing(closeMainWindow)
        windows.Show()

    })

    if err != nil {
        panic(err)
    }
}

func closeMainWindow(*ui.Window) bool {
	ui.Quit()
	return true
}


func main() {
	createMainWindow()
}

