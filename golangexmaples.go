package golangexamples
import (
"fmt"
"github.com/ehteshamz/greetings"
)
func ConcatSlice string(slicetoconcat []byte) {
  var s string
  var b string
  var del="-"``
  for i:=0;i<len(slicetocontact);i++{
    b=string(slicetoconcat[i])
    s=append(s,b)
    s=append(s,del)
  }
  return s
}
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
  for i:=0;i<len(sliceToEncrypt);i++{
    for j:=0;j<ceaserCount;j++{
      sliceToEncrypt[i++]
    }
  }


}

func EZGreetings string(name string){

  a=greetings.PrintGreetings(name)
  prntf(a)

}
