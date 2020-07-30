package main

import (
	//"os"
	//"bufio"
	"fmt"
	"net/http"
	"io/ioutil"
)

var list []string

func main(){
	//stdin := bufio.NewReader(os.Stdin)
	//http.HandleFunc("/",rootHandler)
	http.HandleFunc("/hoge",hogeHandler)

	/*list = append(list,"hogehoge")
	list = append(list,"aiueo")*/
	http.HandleFunc("/todo",todoHandler)

	//8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

func hogeHandler(w http.ResponseWriter, r *http.Request){
	//fmt.Print("アクセスされたよ\n")
	fmt.Fprint(w, "hoge")
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello world")
}

func todoHandler(w http.ResponseWriter, r *http.Request){
	//fmt.Fprint(w, "todo list")
	switch r.Method{

	case "GET":
		fmt.Fprint(w, "GET hello!\n")
		/*todo listを出力*/
		for _, s:=range list{
			fmt.Fprintf(w,"%s\n",s)
		}

	case "POST":
		fmt.Fprint(w, "POST hello!\n")
		b, err:=ioutil.ReadAll(r.Body)
		if err!=nil {
			fmt.Fprint(w,"Error!!!!!!!!\n")
		}
		
		/*todo listに追加*/
		list = append(list, string(b))

	case "DELETE":
		fmt.Fprint(w,"DELETE hello!\n")
		b, err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			fmt.Fprint(w, "Error!!!!!!\n")
		}

		var i int
		i=0
		for _, s:=range list{
			if s==string(b){
				list=unset(list,i)
			}
			i++
		}
		/*todo listを消去*/

	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func unset(s []string, index int) []string {
	if index>=len(s){
		return s
	}
	return append(s[:index],s[index+1:]...)
}