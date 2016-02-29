package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var isoFilename = filepath.Join(os.Getenv("HOME"), ".docker/machine/cache/boot2docker.iso")

	log.Println("Mounting %s", isoFilename)
	if err := hdiutil("attach", isoFilename); err != nil {
		log.Fatal(err)
	}

	p, _ := os.Open("/Volumes")
	defer p.Close()

	l, _ := ioutil.ReadDir(p.Name())
	vSlice := make([]string, 0)
	var version string

	for _, f := range l {
		re := regexp.MustCompile(`(?P<first>.*)-(?P<second>.*)`)
		re2 := regexp.MustCompile(`(^v.*)`)
		vSlice = re.FindStringSubmatch(f.Name())
		for _, v := range vSlice {
			version = v
			if re2.MatchString(version) {
				break
			}
		}
	}

	isoRootDir := "/Volumes/Boot2Docker-" + version
	vmlinuz64 := isoRootDir + "/boot/vmlinuz64"
	initrd := isoRootDir + "/boot/initrd.img"

	log.Println("Extract vmlinuz64")
	if err := CopyFile(vmlinuz64, "./vmlinuz64"); err != nil {
		log.Fatal(err)
	}
	log.Println("Extract initrd.img")
	if err := CopyFile(initrd, "./initrd.img"); err != nil {
		log.Fatal(err)
	}
	log.Println("Unmounting %s", isoFilename)
	if err := hdiutil("detach", "/Volumes/Boot2Docker-v1.9/"); err != nil {
		log.Fatal(err)
	}
}

func hdiutil(args ...string) error {
	cmd := exec.Command("hdiutil", args...)

	log.Println("executing: %v %v", cmd, strings.Join(args, " "))

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.Chmod(dst, fi.Mode()); err != nil {
		return err
	}

	return nil
}
