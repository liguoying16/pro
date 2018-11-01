package poker_hash


import(
        "fmt"
      )

type Poker_hash struct {
    M_hash string
    M_indexPoker int
    M_startBit int
    M_stepBit int

}

func PrintHash(poker Poker_hash){
    fmt.Println(poker.M_hash, poker.M_indexPoker, poker.M_startBit, poker.M_stepBit)
}

func pow(x int, n int) int {
    ret := 1

    if n == 0 {
        return 1
    }
    for i := 0; i < n; i++ {
       ret *= x
    }

    //fmt.Printf("x: %d, n: %d, pow: %d \n ", x, n, ret)
    return ret
}

func GetPoker(poker *Poker_hash) int {

    retNum := 0
    startByte := 0
    endByte := 0

    if (poker.M_indexPoker >= 0) && (poker.M_indexPoker < 22) {
        poker.M_stepBit = 6
    } else if (poker.M_indexPoker >= 22) && (poker.M_indexPoker < 37) {
        poker.M_stepBit = 5
    } else if (poker.M_indexPoker >= 38) && (poker.M_indexPoker < 45) {
        poker.M_stepBit = 4
    } else if (poker.M_indexPoker >= 46) && (poker.M_indexPoker < 49) {
        poker.M_stepBit = 3
    } else if (poker.M_indexPoker >= 50) && (poker.M_indexPoker < 51) {
        poker.M_stepBit = 2
    } else if (poker.M_indexPoker >= 52) && (poker.M_indexPoker < 53) {
        poker.M_stepBit = 1
    }


    var bNum []byte = []byte(poker.M_hash)

    if poker.M_startBit > 255 {
        poker.M_startBit = 0
    }

    if poker.M_indexPoker > 53 {
        return -1
    }

    startBit := poker.M_startBit
    endBit := poker.M_startBit + poker.M_stepBit - 1
    startByte = startBit/8
    endByte = endBit/8

    //fmt.Printf("pokerIndx:%d, startBit: %d, step: %d, endBit: %d, ", poker.M_indexPoker, startBit, poker.M_stepBit, endBit)

    if(startByte == endByte) {
        retNum = int(bNum[startByte])
        //fmt.Printf("retNum：%d \n", retNum)
        oprNum := pow(2, 8-(startBit%8))-1
        //fmt.Printf("oprNum：%d\n", oprNum)
        retNum &= oprNum
        //fmt.Printf("retNum：%d \n", retNum)
        for k := 0; k < (8 - (endBit%8)-1); k++ {
            retNum = retNum >> 1
        }

    } else if(startByte == (endByte-1)) {
        retNum = int(bNum[startByte])
        //fmt.Printf("retNum：%d \n", retNum)
        oprNum := pow(2, 8-(startBit%8))-1
        //fmt.Printf("oprNum：%d ", oprNum)
        retNum &= oprNum
        //fmt.Printf("retNum：%d \n", retNum)

        for j := 0; j < ((endBit%8)+1); j++ {
            retNum = retNum << 1
        }
        //fmt.Printf("retNum：%d \n", retNum)

        retNum2 := int(bNum[endByte])
        for k := 0; k < (8-(endBit%8)-1); k++ {
            retNum2 = retNum2 >> 1
        }

        retNum += retNum2
        //fmt.Printf("retNum：%d \n", retNum)



    } else {
        fmt.Printf("error: startByte: %d, endByte: %d\n", startByte, endByte);
        return -1
    }


    //fmt.Printf("retNum: %02d \n", retNum)

    poker.M_indexPoker += 1
    poker.M_startBit = endBit + 1


    for i := 0; i < 32; i++ {
        //fmt.Printf("%02x ", bNum[i])
    }

    return retNum
}


