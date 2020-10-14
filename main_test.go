package main

import "testing"

func TestAbs(t *testing.T) {
	got := int(-1)
	if got != -1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
//conn, err := net.Dial("tcp", "jsonplaceholder.typicode.com:80")
//if err != nil {
//	log.Panic(err)
//}
//_, err = conn.Write([]byte("GET /todos HTTP/1.0\r\nHost: jsonplaceholder.typicode.com\r\nContent-Type: application/*\r\n\r\n"))
//checkError(err)
//result, err := ioutil.ReadAll(conn)
//checkError(err)
//fmt.Fprintf(os.Stdout,"%s", string(result))