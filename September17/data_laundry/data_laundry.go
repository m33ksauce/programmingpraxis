package datalaundry

import (
	"bytes"
	"io/ioutil"
	"os"
)

type iface struct {
	name   string
	inet   string
	status string
}

func readFile(f string, i *iface) error {
	i.name = "lo0"
	i.inet = "127.0.0.1"

	return nil
}

func cleanFile(in string, out string) error {
	inFile, err := ioutil.ReadFile(in)

	lines := bytes.Split(inFile, []byte("\n"))

	var interfaces []iface
	var output []byte
	var firstLine bool
	var i iface
	for _, k := range lines {
		if k[0] != byte('\t') {
			firstLine = true
		}
		if firstLine {
			firstLine = false
			if i.name != "" {
				interfaces = append(interfaces, i)
			}
			i = iface{}
			t := bytes.Index(k, []byte(":"))
			i.name = string(k[:t])
		}
		if !firstLine {
			if bytes.Index(k, []byte("inet ")) > -1 {
				t := k[bytes.Index(k, []byte("inet "))+5:]
				i.inet = string(t[:bytes.Index(t, []byte(" "))])
			}
			if bytes.Index(k, []byte("status: ")) > -1 {
				t := k[bytes.Index(k, []byte("status: "))+8:]
				i.status = string(t)
			}
		}
	}
	interfaces = append(interfaces, i)

	output = []byte("interface,inet,status")
	for _, n := range interfaces {
		output = append(output, []byte("\n")...)
		output = append(output, []byte(n.name)...)
		output = append(output, []byte(",")...)
		output = append(output, []byte(n.inet)...)
		output = append(output, []byte(",")...)
		output = append(output, []byte(n.status)...)
	}

	err = ioutil.WriteFile(out, []byte(output), os.ModeAppend)

	return err
}
