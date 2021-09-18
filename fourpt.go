package main 
import  (
	"fmt"
)

func And(x1,x2 bool) (bool){
	return (x1 && x2)
}

func Or(x1,x2 bool) (bool){
	return (x1||x2)
}

func Xor(x1,x2 bool) (bool){
	return  (x1!=x2)
}

func Nand(x1,x2 bool) (bool){
	return !And(x1,x2)
}


func main(){

	fmt.Println("And 0,0",And(false,false))
	fmt.Println("And 1,0",And(true,false))
	fmt.Println("And 0,1",And(false,true))
	fmt.Println("And 1,1",And(true,true))

	fmt.Println("Or 0,0",Or(false,false))
	fmt.Println("Or 1,0",Or(true,false))
	fmt.Println("Or 0,1",Or(false,true))
	fmt.Println("Or 1,1",Or(true,true))

	fmt.Println("Xor 0,0",Xor(false,false))
	fmt.Println("Xor 1,0",Xor(true,false))
	fmt.Println("Xor 0,1",Xor(false,true))
	fmt.Println("Xor 1,1",Xor(true,true))

	fmt.Println("Nand 0,0",Nand(false,false))
	fmt.Println("Nand 1,0",Nand(true,false))
	fmt.Println("Nand 0,1",Nand(false,true))
	fmt.Println("Nand 1,1",Nand(true,true))
}
