package main

import ("fmt" 
"encoding/hex" 
"encoding/base64")

func Sum(C, E int) int {
	C += E
	return C
}

func Substract(C, E int) int {
	C -= E
	return C
}

func Multiply(C, E int) int {
	C *= E
	return C
}

func Divided(C, E int) int {
	if E != 0 {
		C /= E
		return C
	} else {
		return 0
	}
}

func Modulous(C, E int) int {
	C %= E
	return C
}

func And_Operator(C, E int) {
	if C > 0 && E > 0 {
		fmt.Println("Both Integer are positive!")
	}

}

func OR_Operator(C, E int) {
	if C > E || E > C {
		fmt.Println("One is greater than other.")
	}
}

func Not_Operator(C, E int) {
	if C != E {
		fmt.Println("They are not equal to each other")
	} else {
		fmt.Println("THey are equal to each other")
	}
}


func My_AND(C, E int) int {
	R := C & E
	return R
}


func My_OR(C, E int) int {
	R := C | E
	return R
}

// func My_XOR(C, E int) int {
	// R := C ^ E
	// return R
// }



func My_XOR(message string, key byte) string{
	mes_byte := []byte(message)
	for i := 0; i < len(mes_byte); i++{
		mes_byte[i] ^= key
	}
	return string(mes_byte)
}

func My_Not(C, E int) (int, int) {
	return ^C, ^E
}


func Left_Shift(C, E , shift int) (int, int){
	return C << shift, E << shift
}

func Right_Shift(C, E , shift int) (int, int){

	return C >> shift, E >> shift

}



func Hex_encoding(Text string) string{
	Enc_Text := hex.EncodeToString([]byte(Text))
	return Enc_Text
}


func Base64_encoding(Text string) string {
	Enc_Text := base64.StdEncoding.EncodeToString([]byte(Text))
	return Enc_Text
}



func Binary_encoding(Text string) string {
	Enc_Text := ""
	for i := 0; i < len(Text); i++ {
		Enc_Text += fmt.Sprintf("%08b ", Text[i])
	}
	return Enc_Text
}



func main() {

	// fmt.Print("Hello World!")

	// // Lab1: Task 1: Assignment Operator Demonstration

	//  = Operator
	// var C int
	// var D int
	// fmt.Printf("Input Value C: ")
	// fmt.Scan(&C)
	// D = C
	// fmt.Printf("D = C = %d\n", D)

	// var A int
	// fmt.Print("Input Value A: ")
	// fmt.Scan(&A)
	// var B int
	// fmt.Print("Input Value B: ")
	// fmt.Scan(&B)

	// // += Operator
	// Result1 := Sum(A,B)
	// fmt.Println("A += B:",Result1)

	// // -= Operator
	// Result2 := Substract(A,B)
	// fmt.Println("A -= B:",Result2)

	// // *= Operator
	// Result3 := Multiply(A,B)
	// fmt.Println("A *= B:",Result3)

	// // /= Operator
	// Result4 := Divided(A,B)
	// fmt.Println("A /= B:",Result4)

	// // %= Operator
	// Result5 := Modulous(A, B)
	// fmt.Println("A %= B:",Result5)

	// // Lab2 : Logical Operator
	// And_Operator(A, B)
	// OR_Operator(A, B)
	// Not_Operator(A, B)

	// // Lab3 : Bitwise and Assignment Operator

	// R1 := My_OR(A, B)
	// fmt.Println("A XOR B:", R1)

	// R2 := My_AND(A, B)
	// fmt.Println("A AND B:", R2)

	// R3 := My_OR(A, B)
	// fmt.Println("A OR B:", R3)

	// R4, R5 := My_Not(A, B)
	// fmt.Println("Not A:", R4)
	// fmt.Println("Not B:",R5)

	// R6, R7 := Left_Shift(A, B, 2)
	// fmt.Printf("Left shift of %d is: %d\n", A, R6)
	// fmt.Printf("Left shift of %d is: %d\n", B, R7)

	// R8, R9 := Right_Shift(A, B, 2)
	// fmt.Printf("Right shift of %d is: %d\n",A, R8)
	// fmt.Printf("Right shift of %d is: %d\n",B, R9)

	// // Lab 4: Mini Calculator

	// for {
	// 	var Option int
	// 	fmt.Printf("[1]. Add\n[2]. Sub\n[3]. Mult\n[4]. Div\n[5]. Mod\n[6]. Exit\n")
	// 	fmt.Printf("Enter the options: ")
	// 	fmt.Scan(&Option)
	// 	if Option == 6 {
	// 		fmt.Println("Bye")
	// 		break
	// 	} else {
	// 		var Int1 int
	// 		var Int2 int
	// 		fmt.Printf("Enter Integer 1: ")
	// 		fmt.Scan(&Int1)
	// 		fmt.Printf("Enter Integer 2: ")
	// 		fmt.Scan(&Int2)

	// 		switch Option {
	// 		case 1:
	// 			Result1 := Sum(Int1, Int2)
	// 			fmt.Println("Result:", Result1)
	// 		case 2:
	// 			Result2 := Substract(Int1, Int2)
	// 			fmt.Println("Result:", Result2)
	// 		case 3:
	// 			Result3 := Multiply(Int1, Int2)
	// 			fmt.Println("Result:", Result3)
	// 		case 4:
	// 			Result4 := Divided(Int1, Int2)
	// 			if Result4 == 0 {
	// 				fmt.Println("Error!!, Integer can't divided by 0!")
	// 			} else {
	// 				fmt.Println("Result:", Result4)
	// 			}
	// 		case 5:
	// 			Result5 := Modulous(Int1, Int2)
	// 			fmt.Println("Result:", Result5)
	// 		}

	// 	}

	// }

// Lab 5: Hex Binary and Base64 Encoding 
	// var input string
	// fmt.Printf("Enter the text: ")
	// fmt.Scan(&input)
	// // calling function
	// Encoded_Hex := Hex_encoding(input)
	// Encoded_Base64 := Base64_encoding(input)
	// Encoded_Bin := Binary_encoding(input)

	// fmt.Println("Hex:", Encoded_Hex)
	// fmt.Println("Base64:", Encoded_Base64)
	// fmt.Println("Binary:", Encoded_Bin)


// Lab 6: XOR Encryption
	var Plain_text string
	fmt.Printf("Enter Plain text: ")
	fmt.Scan(&Plain_text)
	var Key int
	fmt.Printf("Enter key: ")
	fmt.Scan(&Key)
	Cipher_Text := My_XOR(Plain_text, byte(Key))
	fmt.Println("Cipher Text:",Cipher_Text)
	Decrypt_Text := My_XOR(Cipher_Text, byte(Key))
	fmt.Println("Decrypted Text:",Decrypt_Text)
	


}
