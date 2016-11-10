package main
import (
    "strconv"
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


