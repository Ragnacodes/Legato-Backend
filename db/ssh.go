package legatoDb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
)

const sshType string = "sshes"

type Ssh struct {
	gorm.Model
	Username     string
	Host         string
	Password     string
	SshKey       string
	ConnectionID uint
	Connection   *Connection `gorm:"foreignkey:id;references:ConnectionID"`
	Service      Service     `gorm:"polymorphic:Owner;"`
}
type loginWithPasswordData struct {
	Commands []string `json:"commands"`
	Host     string   `json:"host"`
	Password string   `json:"password"`
	Username string   `json:"username"`
}
type loginWithSshKeyData struct {
	Username string   `json:"username"`
	Host     string   `json:"host"`
	SshKey   string   `json:"sshKey"`
	Commands []string `json:"commands"`
}

type updateData struct {
	ConnectionId uint     `json:"connectionid"`
	Username     string   `json:"username"`
	Host         string   `json:"host"`
	SshKey       string   `json:"sshKey"`
	Commands     []string `json:"commands"`
}

func (ss *Ssh) String() string {
	return fmt.Sprintf("(@SSH: %+v)", *ss)
}

func (ldb *LegatoDB) CreateSshForScenario(s *Scenario, ssh Ssh) (*Ssh, error) {
	ssh.Service.UserID = s.UserID
	ssh.Service.ScenarioID = &s.ID
	var dataWithPass loginWithPasswordData
	err1 := json.Unmarshal([]byte(ssh.Service.Data), &dataWithPass)
	if err1 == nil {
		ssh.Host = dataWithPass.Host
		ssh.Username = dataWithPass.Username
		ssh.Password = dataWithPass.Password
	}

	var dataWithkey loginWithSshKeyData
	err := json.Unmarshal([]byte(ssh.Service.Data), &dataWithkey)
	if err == nil {
		ssh.Host = dataWithkey.Host
		ssh.Username = dataWithkey.Username
		ssh.SshKey = dataWithkey.SshKey
	}
	ldb.db.Create(&ssh)
	ldb.db.Save(&ssh)
	return &ssh, nil

}

func (ldb *LegatoDB) GetSshByID(id uint, u *User) (Ssh, error) {
	var s Ssh
	err := ldb.db.Where(&Connection{UserID: u.ID}).Where("ID = ?", id).Find(&s).Error
	if err != nil {
		return Ssh{}, err
	}
	return s, nil
}
func (ldb *LegatoDB) GetUserSsh(u *User) ([]Ssh, error) {
	var services []Service
	err := ldb.db.Select("id").Where(&Service{UserID: u.ID}).Find(&services).Error
	if err != nil {
		return nil, err
	}
	var serviceIds []uint
	serviceIds = []uint{}
	for _, srv := range services {
		serviceIds = append(serviceIds, srv.ID)
	}

	var sshs []Ssh
	err = ldb.db.Where(serviceIds).Preload("Service").Find(&sshs).Error
	if err != nil {
		return nil, err
	}
	return sshs, nil
}

func (ldb *LegatoDB) GetSshByService(serv Service) (*Ssh, error) {
	var s Ssh
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&s).Error
	if err != nil {
		return nil, err
	}
	if s.ID != uint(serv.OwnerID) {
		return nil, errors.New("the ssh service is not in this scenario")
	}

	return &s, nil
}
func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
func ConnectWithUserPass(myssh Ssh, commands []string) {
	// SSH client config
	config := &ssh.ClientConfig{
		User: myssh.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(myssh.Password),
		},
		// Non-production only
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to host
	client, err := ssh.Dial("tcp", myssh.Host+":"+"22", config)
	if err != nil {
		log.Println("unable to authenticate username or password is incorrect")
	}
	defer client.Close()

	if client == nil {
		log.Println("client is null")
		return
	}

	// Create session
	session, err := client.NewSession()
	if err != nil {
		log.Print("Failed to create session: ", err)
	}
	defer session.Close()

	// Std in Pipe for commands
	stdIn, err := session.StdinPipe()
	if err != nil {
		log.Print(err)
	}
	commandsInOneLine := ""
	reverseAny(commands)
	for i, con := range commands {
		if i < len(commands)-1 {
			commandsInOneLine += con + "&"
		} else {
			commandsInOneLine += con
		}

	}
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(commandsInOneLine); err != nil {
		log.Print("Failed to run: " + err.Error())
	}
	log.Print("\n", b.String())
	_ = stdIn.Close()
	_ = session.Close()
	_ = client.Close()
	// Uncomment to store in variable

}
func ConnectWithSShKey(myssh Ssh, commands []string) {
	signer, err := ssh.ParsePrivateKey([]byte(myssh.SshKey))

	if err != nil {
		log.Print(err)

	}

	config := &ssh.ClientConfig{
		User: myssh.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to host
	client, err := ssh.Dial("tcp", myssh.Host+":"+"22", config)
	if err != nil {
		log.Println("unable to authenticate username or sshkey is incorrect")
	}
	defer client.Close()

	if client == nil {
		log.Println("client is null")
		return
	}

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		log.Print("Failed to create session: ", err)
	}
	defer session.Close()

	// Std in Pipe for commands
	stdIn, err := session.StdinPipe()
	if err != nil {
		log.Print(err)
	}
	commandsInOneLine := ""
	reverseAny(commands)
	for i, con := range commands {
		if i < len(commands)-1 {
			commandsInOneLine += con + "&"
		} else {
			commandsInOneLine += con
		}

	}
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(commandsInOneLine); err != nil {
		log.Println("Failed to run: " + err.Error())
	}
	log.Print("\n", b.String())
	_ = stdIn.Close()
	_ = session.Close()
	_ = client.Close()

}

// Service Interface for ssh
func (ss Ssh) Execute(...interface{}) {
	err := legatoDb.db.Preload("Service").Find(&ss).Error
	if err != nil {
		log.Println(err)
	}

	SendLogMessage("*******Starting SSH Service*******", *ss.Service.ScenarioID, nil)

	logData := fmt.Sprintf("Executing type (%s) : %s\n", sshType, ss.Service.Name)
	SendLogMessage(logData, *ss.Service.ScenarioID, nil)

	var dataWithPass loginWithPasswordData
	mySsh := Ssh{}
	err = json.Unmarshal([]byte(ss.Service.Data), &dataWithPass)

	hasPassword := strings.Contains(ss.Service.Data, "password") == true

	var dataWithkey loginWithSshKeyData
	err1 := json.Unmarshal([]byte(ss.Service.Data), &dataWithkey)
	if err1 != nil {
		log.Print(err1)
	}
	var commands []string
	if hasPassword {
		SendLogMessage("Connecting with Password ...", *ss.Service.ScenarioID, &ss.Service.ID)
		commands = dataWithPass.Commands
		mySsh.Username = dataWithPass.Username
		mySsh.Password = dataWithPass.Password
		mySsh.Host = dataWithPass.Host
		ConnectWithUserPass(mySsh, dataWithPass.Commands)

	} else {
		SendLogMessage("Connecting with SSH KEY ...", *ss.Service.ScenarioID, &ss.Service.ID)
		commands = dataWithkey.Commands
		mySsh.Username = dataWithkey.Username
		mySsh.SshKey = dataWithkey.SshKey
		mySsh.Host = dataWithkey.Host
		ConnectWithSShKey(mySsh, dataWithkey.Commands)

	}

	logData = fmt.Sprintf("username: %s host: %s", mySsh.Username, mySsh.Host)
	SendLogMessage(logData, *ss.Service.ScenarioID, &ss.Service.ID)

	logData = fmt.Sprintf("Executing following commands on remote server: %s", commands)
	SendLogMessage(logData, *ss.Service.ScenarioID, &ss.Service.ID)

	ss.Next()
}

func (ss Ssh) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", sshType, ss.Service.Name)
}

func (ss Ssh) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&ss).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		return
	}

	log.Printf("Executing \"%s\" Children \n", ss.Service.Name)

	for _, node := range ss.Service.Children {
		go func(n Service) {
			serv, err := n.Load()
			if err != nil {
				log.Println("error in loading services in Next()")
				return
			}

			serv.Execute()
		}(node)
	}

	log.Printf("*******End of \"%s\"*******", ss.Service.Name)
}

func (ldb *LegatoDB) UpdateSsh(s *Scenario, servId uint, ssh Ssh) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var ss Ssh
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&ss).Error
	if err != nil {
		return err
	}
	if ss.Service.ID != servId {
		return errors.New("the ssh service is not in this scenario")
	}
	var a updateData
	err = json.Unmarshal([]byte(ssh.Service.Data), &a)
	if err != nil {
		log.Println("con not update ssh")
	}

	if a.ConnectionId != 0 {
		ssh.ConnectionID = a.ConnectionId
	}

	ldb.db.Model(&serv).Updates(ssh.Service)
	ldb.db.Model(&ss).Updates(ssh)

	if ssh.Service.ParentID == nil {
		legatoDb.db.Model(&serv).Select("parent_id").Update("parent_id", nil)
	}

	return nil
}
