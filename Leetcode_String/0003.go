package Leetcode_String

import (
  "bufio"
  "os"
  //"fmt"
  "strings"
 )

/*
题目：
思路：
时间：2023/12/10
链接：
 */

func AcmMode() {
  // 多行输入
  input := bufio.NewScanner(os.Stdin)
  output := bufio.NewWriter(os.Stdout)
  // 测试用例过长的时候，要先指定bufio的bufer，否则会读取失败
  // bufer := make([]byte,2000*1024)
  // input.Buffer(bufer,len(bufer))
  for input.Scan(){ // 每调用一次Scan,从输入流中读取一行
    // 获取这一行的输入
    str := strings.Split(input.Text()," ") 
    //TODO
    // 将结果写入缓冲输出区
    output.WriteString(str[0]) 
    //TODO
  }
  // 输出答案
  output.Flush() 
}


