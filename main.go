package main

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {

	var(

		out string="C:/Users/Soucer/OneDrive - DPSP"
		in string="C:/Users/Soucer/OneDrive/Interface/05"
	)



	var f []string
	err := IOReadDir(&f,out)

	if err != nil {

		println(err.Error())
	}

	for _, a := range f {

		aux:=strings.Replace(a,out,in,1)

		println("Copiando: "+strings.Replace(a,out,"",1))

		Copy(a,aux)
	}

}
func IOReadDir(files *[]string, root string) error {

	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return  err
	}

	for _, file := range fileInfo {

		if file.IsDir() {

			IOReadDir(files, path.Join(root,file.Name()))
		} else {


		*files = append(*files, path.Join(root,file.Name()))
	}
	}

	return nil
}

func Copy(src, dst string){

	aux,err:=os.Stat(src)
	if err!=nil{

		println(err.Error())
	}

	mk:=strings.Replace(dst,aux.Name(),"",1)

	os.MkdirAll(mk,os.ModePerm)

	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}