package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateRSAKey(pripath, pubpath, passwd string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	//使用pem格式对x509输出的内容进行编码
	privateFile, err := os.Create(pripath)
	if err != nil {
		return err
	}
	defer privateFile.Close()

	//构建一个pem.Block结构体对象
	//privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	privateBlock, err := x509.EncryptPEMBlock(rand.Reader, "RSA Private Key", x509PrivateKey, []byte(passwd), x509.PEMCipherAES256)
	if err != nil {
		return err
	}

	//将数据保存到文件
	err = pem.Encode(privateFile, privateBlock)
	if err != nil {
		return err
	}

	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	//pem格式编码
	publicFile, err := os.Create(pubpath)
	if err != nil {
		return err
	}
	defer publicFile.Close()

	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//将数据保存到文件
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}

	return nil
}

func genkey(name, passwd string) error {
	err := generateRSAKey(name+".pri", name+".pub", passwd)
	if err != nil {
		fmt.Println("Rsa key gen failed", err)
		return err
	}

	fmt.Printf("Gen privatekey: %s, publickey: %s\n", name+".pri", name+".pub")
	return nil
}

func main() {
	genkey("rsa", "password")
}
