package main

import(
	"fmt"
	"os"
)

func CreateFile(filename string, data string) (error){
	SData := []byte(data)
	err := os.WriteFile(filename,SData,0600)
	if err != nil{
		fmt.Println("Не удалось создать файл")
		return err
	} else{
		return nil
	}

}

func ReadFile(filename string) (error, string){
	data, err := os.ReadFile(filename)
	if err != nil{
		fmt.Println("Не удалось прочитать файл")
		return err, "" 
	} else{
		return nil, string(data)
	}
}

func DeleteFile(filename string) error{
	err := os.Remove(filename)
	if err != nil{
		fmt.Println("Не удалось удалить файл")
		return err
	} else{
		fmt.Println("Файл успешно удален")
		return nil
	}
}

func main(){
	err1 := CreateFile("test.txt", "hui")
	if err1 !=nil{
		fmt.Println(err1)
	}
	err, data := ReadFile("test.txt")
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(data)
	}
	err = DeleteFile("test.txt")
	if err != nil{
		fmt.Println(err)
	}

}
