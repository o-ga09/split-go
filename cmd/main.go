package main

import (
	"flag"
	"fmt"
	"main/split"
	"strconv"
	"strings"
)

var (
	lineNum          string
	fileNum          string
	byteNum          string
	columnNum        string
	suffix           string
)

func init() {
	flag.StringVar(&fileNum, "n", "default", "file split number")
    flag.StringVar(&lineNum, "l", "default", "line split number")
	flag.StringVar(&byteNum, "b", "default", "byte split number")
	flag.StringVar(&columnNum, "c", "default", "column split number")
	flag.StringVar(&suffix,  "a", "default", "output file suffix")
	flag.Usage = func() {
		fmt.Println("usage: split [-l line_count] [-a suffix_length] [file [prefix]]")
		fmt.Println("       split -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]")
		fmt.Println("       split -n chunk_count [-a suffix_length] [file [prefix]]")
		fmt.Println("       split -c column number [-a suffix_length] [file [prefix]]")
		fmt.Println("notice: option -c is only csv file")
	}
}

func main () {
	flag.Parse()

	// flagの組み合わせのバリデーションチェック
	if !Validate(){
		flag.Usage()
	}

	// メイン処理
	res := run()
	if res != 0 {
		handleError(res)
		flag.Usage()
	}
}

func Validate() bool {
	// -n の指定がある
	if strings.Compare(fileNum,"default") != 0 {
		if strings.Compare(lineNum,"default") != 0 || 
		   strings.Compare(columnNum,"default") != 0 ||
		   strings.Compare(byteNum,"default") != 0 {
			return false
		   }
	}
	// -l の指定がある
	if strings.Compare(lineNum,"default") != 0 {
		if strings.Compare(fileNum,"default") != 0 || 
		   strings.Compare(columnNum,"default") != 0 ||
		   strings.Compare(byteNum,"default") != 0 {
			return false
		   }
	}
	// -b の指定がある
	if strings.Compare(byteNum,"default") != 0 {
		if strings.Compare(lineNum,"default") != 0 || 
		   strings.Compare(columnNum,"default") != 0 ||
		   strings.Compare(fileNum,"default") != 0 {
			return false
		   }
	}
	// -c の指定がある
	if strings.Compare(columnNum,"default") != 0 {
		if strings.Compare(lineNum,"default") != 0 || 
		   strings.Compare(fileNum,"default") != 0 ||
		   strings.Compare(byteNum,"default") != 0 {
			return false
		   }
	}
	// バリデーションOK
	return true
}

func handleError(e int) {
	switch e {
		case 1 :
			fmt.Println("invalid argument")
		case 2 :
			fmt.Println("can not open file")
		case 3 :
			fmt.Println("can not read file")
		case 4:
			fmt.Println("can not write file")
		case 5:
			fmt.Println("option -c is only csv file")
	}
}

func run() int {
	args := flag.Args()
	// flag 指定なし
	if strings.Compare(lineNum,"default") == 0 &&
	   strings.Compare(fileNum,"default") == 0 &&
	   strings.Compare(byteNum,"default") == 0 &&
	   strings.Compare(columnNum,"default") == 0 {
		return split.OptionN(1,args)
	}
	// -n 指定
	if strings.Compare(fileNum,"default") != 0 {
		n,_ := strconv.Atoi(fileNum)
		return split.OptionN(n,args)
	}
	// -l 指定
	if strings.Compare(lineNum,"default") != 0 {
		l,_ := strconv.Atoi(lineNum)
		return split.OptionL(l,args)
	}
	// -b 指定
	if strings.Compare(byteNum,"default") != 0 {
		b,_ := strconv.Atoi(byteNum)
		return split.OptionB(b,args)
	}
	// -c 指定
	if strings.Compare(columnNum,"default") != 0 {
		c,_ := strconv.Atoi(columnNum)
		return split.OptionC(c,args)
	}

	return 0
}