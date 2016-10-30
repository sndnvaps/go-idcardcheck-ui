//Create by sndnvaps <sndnvaps@gmail.com>

package main

import (
    "strconv"
    "github.com/andlabs/ui"
)

/*
 * ai -> a1 , a2, a3, a4, a5, a6... a17 (a18 是校验码) 身份证前17位对应(ai)
 * wi -> 7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2 (17位)
 *
 *  int  res = 0;
 *  for (i = 1; i < 17; i++) {
 *          res += (a[i] * w[i])
 *          }
 *     int  y = res % 11；
 *
 *
 * y 与 a18的对应关系
 *
 *  y    { 0, 1,  2,  3, 4, 5, 6, 7, 8, 9, 10}
 *  a18  { 1, 0, 'X', 9, 8, 7, 6, 5, 4, 3,  2 } -> vefiry[18] = { 1, 0, 'X', 9, 8, 7, 6, 5, 4, ,3, 2};
 */

//测试身份证号码：34052419800101001X
//测试身份证号码：511028199507215915


//covert byte to int 
func byte2int(x byte) byte {
    if x == 88 {
        return 'X'
    }
    return (x - 48) // 'X' - 48 = 40;
}


//calculate the latest num of the id card 
func check_id(id [17]byte) int {
    arry := make([]int, 17)

    //强制类型转换，将[]byte转换成[]int ,变化过程
    // []byte -> byte -> string -> int
    //将通过range 将[]byte转换成单个byte,再用强制类型转换string()，将byte转换成string
    //再通过strconv.Atoi()将string 转换成int 类型
    for index , value := range id {
        arry[index], _ = strconv.Atoi(string(value))
    }

    var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
    var res int
    for i := 0; i < 17; i++ {
        //fmt.Println("id =", i, byte2int(id[i]), wi[i])
        res += arry[i] * wi[i]
    }

    return (res % 11)
}

//compare the calculate key with the latest num of the id card
func verify_id(verify int, id_v byte) (bool, string) {
    var temp byte
    var i int
    a18 := [11]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2}

    for i = 0; i < 11; i++ {
        if i == verify {
            temp = a18[i]
            break
        }
    }
    if temp == id_v {
        return true, "验证成功"
    }
    return false, "验证失败"
}


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

