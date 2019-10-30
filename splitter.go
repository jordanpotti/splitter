package main

import (
    "bufio"
    "strconv"
    "os"
    "flag"
    "bytes"
    "io"
)


// This tool re-implements the Linux split tool, but it'll probably be bettter and waay more buggy. And cross platform.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func split(target string, number int, pre string, post string, total int) {
    dat, err := os.Open(target)
    defer dat.Close()
    check(err)
    scanner := bufio.NewScanner(dat)
    i := 0
    j := 0
	split := total / number
	split += 1
	var txtlines []string
    for scanner.Scan() {
        if i < split {
           txtlines = append(txtlines, scanner.Text())
           i += 1
        } else {
            f, _ := os.Create (pre + strconv.Itoa(j) + post)
            defer f.Close()
            for _, element := range txtlines {
				element = element + "\r\n"
                f.WriteString(element)
			 }
			 txtlines = nil
			 j += 1
			 i = 0
         }
	 }
	 f, _ := os.Create (pre + strconv.Itoa(j) + post)
	 defer f.Close()
	 for _, element := range txtlines {
		element = element + "\r\n"
		f.WriteString(element)
	  }
	  txtlines = nil
	  j += 1
	  split += split
}

func lineCounter(file string) (int, error) {
    var r io.Reader
    //var err error
    r, _ = os.Open(file)
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}

func main() {
    wordPtr := flag.String("target", "", "Target File to split")
    numbPtr := flag.Int("numb", 2, "Number of files to split the target file into")
	prePtr := flag.String("pre", "", "String to prepend to output")
	postPtr := flag.String("post", "", "String to append to the output")
    flag.Parse()
    total, _ := lineCounter(*wordPtr)
    split(*wordPtr, *numbPtr, *prePtr, *postPtr, total)
}