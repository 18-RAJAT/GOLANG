package main
import "fmt"

func main(){
    fmt.Print("Enter the number ")
    var n int

    fmt.Scanln(&n)
    if(n%2==0){
        fmt.Println(n,"Even");
    }else{
        fmt.Println(n,"Odd");
    }
}