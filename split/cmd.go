package split

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	ERR_INVALID_ARGUMENT = 1
	ERR_OPEN_FILE = 2
	ERR_READ_FILE = 3
	ERR_WRITE_FILE = 4
	ERR_ONLY_CSV = 5
)


func OptionN(n int, args []string) int {
	// 値の範囲のバリデーション
	if n < 0 {
		return ERR_INVALID_ARGUMENT
	}

	// suffix指定有り無し
	switch len(args) {
		case 0:
			return ERR_INVALID_ARGUMENT
		case 1:
			// suffix指定無し
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			count, err := io.ReadAll(file)
			if err != nil {
				return ERR_READ_FILE
			}
		
			// 行を50行ずつに分割する
			lines := strings.Split(string(count), "\n")
			splitNum := len(lines) / n
		
			// 各行を書き込む
			start := 0
			for i := 0; i < int(n); i++ {
				bytes := []byte(strings.Join(lines[start:start+int(splitNum)], "\n"))
				err := os.WriteFile(fmt.Sprintf("output_%d.txt",i+1), bytes, 0644)
				if err != nil {
					return ERR_WRITE_FILE
				}
				start += int(splitNum)
			}
		case 2:
			// suffix指定有り
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			count, err := io.ReadAll(file)
			if err != nil {
				return ERR_READ_FILE
			}
		
			// 行を50行ずつに分割する
			lines := strings.Split(string(count), "\n")
			splitNum := len(lines) / n
		
			// 各行を書き込む
			start := 0
			for i := 0; i < int(n); i++ {
				bytes := []byte(strings.Join(lines[start:start+int(splitNum)], "\n"))
				err := os.WriteFile(fmt.Sprintf("%s_output_%d.txt",args[1],i+1), bytes, 0644)
				if err != nil {
					return ERR_WRITE_FILE
				}
				start += int(splitNum)
			}
		default :
			return ERR_INVALID_ARGUMENT
	}
	return 0
}

func OptionL(l int, args []string) int {
	// 値の範囲のバリデーション
	if l <= 0 {
		return ERR_INVALID_ARGUMENT
	}

	// suffix指定有り無し
	switch len(args) {
		case 0:
			return ERR_INVALID_ARGUMENT
		case 1:
			// suffix指定無し
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			count, err := io.ReadAll(file)
			if err != nil {
				return ERR_READ_FILE
			}
		
			// 行を50行ずつに分割する
			lines := strings.Split(string(count), "\n")
		
			// 各行を書き込む
			suffix := 0
			col := len(lines)
			var bytes []byte
			for i := 0; i < len(lines); i+=l {
				if i+l > len(lines) {
					bytes = []byte(strings.Join(lines[i:col], "\n"))	
				} else {
					bytes = []byte(strings.Join(lines[i:i+l], "\n"))
				}
				err := os.WriteFile(fmt.Sprintf("output_%d.txt",suffix+1), bytes, 0644)
				if err != nil {
					return ERR_WRITE_FILE
				}
				suffix++
			}
		case 2:
			// suffix指定有り
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			count, err := io.ReadAll(file)
			if err != nil {
				return ERR_READ_FILE
			}
		
			// 行を50行ずつに分割する
			lines := strings.Split(string(count), "\n")
		
			// 各行を書き込む
			suffix := 0
			col := len(lines)
			var bytes []byte
			for i := 0; i < len(lines); i+=l {
				if i+l > len(lines) {
					bytes = []byte(strings.Join(lines[i:col], "\n"))	
				} else {
					bytes = []byte(strings.Join(lines[i:i+l], "\n"))
				}
				err := os.WriteFile(fmt.Sprintf("%s_output_%d.txt",args[1],suffix+1), bytes, 0644)
				if err != nil {
					return ERR_WRITE_FILE
				}
				suffix++
			}
		default :
			return ERR_INVALID_ARGUMENT
	}
	return 0
}

func OptionB(b int, args []string) int {
	// 値の範囲のバリデーション
	if b <= 0 {
		return ERR_INVALID_ARGUMENT
	}

	// suffix指定有り無し
	switch len(args) {
		case 0:
			return ERR_INVALID_ARGUMENT
		case 1:
			// suffix指定無し
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			count, err := io.ReadAll(file)
			if err != nil {
				return ERR_READ_FILE
			}
		
			// 行を50行ずつに分割する
			lines := strings.Split(string(count), "\n")
		
			// 各行を書き込む
			suffix := 0
			var block []byte
			bytes := []byte(strings.Join(lines, "\n"))
			cob := len(bytes)
			for i := 0; i < len(bytes); i += b {
				if i+b > len(bytes) {
					block = bytes[i:cob]
				} else {
					block = bytes[i:i+b]
				}
				err := os.WriteFile(fmt.Sprintf("output_%d.txt",suffix+1), block, 0644)
				if err != nil {
					return ERR_WRITE_FILE
				}
				suffix++
			}
		case 2:
			// suffix指定有り
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			count, err := io.ReadAll(file)
			if err != nil {
				return ERR_READ_FILE
			}
		
			// 行を50行ずつに分割する
			lines := strings.Split(string(count), "\n")
		
			// 各行を書き込む
			suffix := 0
			var block []byte
			bytes := []byte(strings.Join(lines, "\n"))
			cob := len(bytes)
			for i := 0; i < len(bytes); i += b {
				if i+b > len(bytes) {
					block = bytes[i:cob]
				} else {
					block = bytes[i:i+b]
				}
				err := os.WriteFile(fmt.Sprintf("%s_output_%d.txt",args[1],suffix+1), block, 0644)
				if err != nil {
					return ERR_WRITE_FILE
				}
				suffix++
			}
		default :
			return ERR_INVALID_ARGUMENT
	}
	return 0
}

func OptionC(c int, args []string) int {
	// 値の範囲のバリデーション
	if c < 0 || len(args) == 0 {
		return ERR_INVALID_ARGUMENT
	}

	// csv限定機能
	if !strings.Contains(args[0],".csv") {
		return ERR_ONLY_CSV
	}

	// suffix指定有り無し
	switch len(args) {
		case 1:
			// suffix指定無し
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			reader := csv.NewReader(file)

			// 出力ファイルを作成する
			out_file_1 := "output_1.csv"
			out_file_2 := "output_2.csv"

			f1,err1 := os.Create(out_file_1)
			f2,err2 := os.Create(out_file_2)
			if err1 != nil || err2 != nil {
				return ERR_WRITE_FILE
			}
			defer f1.Close()
			defer f2.Close()
		
			w1 := csv.NewWriter(f1)
			w2 := csv.NewWriter(f2)
			for  {
				record, err := reader.Read()
				var w_record_1 []string
				var w_record_2 []string
				if err == io.EOF {
					break
				}

				if err != nil {
					return ERR_READ_FILE
				}

				for i := 0 ; i < len(record); i++ {
					if i < c {
						w_record_1 = append(w_record_1,record[i])
					} else {
						w_record_2 = append(w_record_2,record[i])
					}
				}
				w1.Write(w_record_1)
				w2.Write(w_record_2)
			}
			w1.Flush()
			w2.Flush()
		case 2:
			// suffix指定有り
			// ファイルを開く
			file, err := os.Open(args[0])
			if err != nil {
				return ERR_OPEN_FILE
			}
		
			// ファイルの行数を取得する
			reader := csv.NewReader(file)

			// 出力ファイルを作成する
			out_file_1 := fmt.Sprintf("%s_output_1.csv",args[1])
			out_file_2 := fmt.Sprintf("%s_output_2.csv",args[1])

			f1,err1 := os.Create(out_file_1)
			f2,err2:= os.Create(out_file_2)
			if err1 != nil || err2 != nil {
				return ERR_WRITE_FILE
			}
			defer f1.Close()
			defer f2.Close()
		
			w1 := csv.NewWriter(f1)
			w2 := csv.NewWriter(f2)
			for  {
				record, err := reader.Read()
				var w_record_1 []string
				var w_record_2 []string
				if err == io.EOF {
					break
				}

				if err != nil {
					return ERR_READ_FILE
				}

				for i := 0 ; i < len(record); i++ {
					if i < c {
						w_record_1 = append(w_record_1,record[i])
					} else {
						w_record_2 = append(w_record_2,record[i])
					}
				}
				w1.Write(w_record_1)
				w2.Write(w_record_2)
			}
			w1.Flush()
			w2.Flush()
		default :
			return ERR_INVALID_ARGUMENT
	}
	return 0
} 