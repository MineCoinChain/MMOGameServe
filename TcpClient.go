package main

import (
	"fmt"
	"net"
	"io"
	"log"
	"os"
	"os/signal"
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"MMOGameServe/pb"
	"time"
	"math/rand"
)

//客户端测试软件
type Message struct {
	Len   int32
	MsgId int32
	Data  []byte
}

type TcpClient struct {
	conn net.Conn
	Pid  int32 //玩家ID
	X    float32
	Y    float32
	Z    float32
	V    float32
}

func NewTcpClient(ip string, port int) *TcpClient {
	addrStr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", addrStr)
	if err != nil {
		log.Println("conn err:", err)
	}
	//connection 请求成功
	client := &TcpClient{
		conn: conn,
		Pid:  0,
		X:    0,
		Y:    0,
		Z:    0,
		V:    0,
	}
	return client
}

//当前客户端打包数据
func (this *TcpClient) Pack(MsgID int32, Data []byte) ([]byte, error) {
	outbuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(outbuff, binary.LittleEndian, len(Data)); err != nil {
		log.Println("binary write err len(data):", err)
		return nil, err
	}
	if err := binary.Write(outbuff, binary.LittleEndian, MsgID); err != nil {
		log.Println("binary write err len(MsgID):", err)
		return nil, err
	}
	if err := binary.Write(outbuff, binary.LittleEndian, Data); err != nil {
		log.Println("binary write err len(MsgID):", err)
		return nil, err
	}
	return outbuff.Bytes(), nil
}

//当前客户端拆包数据
func (this *TcpClient) UnPack(headData []byte) (*Message, error) {
	headBufReader := bytes.NewReader(headData)
	head := &Message{}
	//读取Len
	if err := binary.Read(headBufReader, binary.LittleEndian, head.Len); err != nil {
		log.Println("UnPack Len err:", err)
		return nil, err
	}
	//读取MsgID
	if err := binary.Read(headBufReader, binary.LittleEndian, head.MsgId); err != nil {
		log.Println("Unpack MsgId err:", err)
		return nil, err
	}
	return head, nil
}

//当前客户端发送数据的方法
func (this *TcpClient) SendMsg(msgID int32, data proto.Message) {
	binary_data, err := proto.Marshal(data)
	if err != nil {
		log.Println("proto Marshal err:", err)
		return
	}
	//将数据打包成LTV格式
	sendData, err := this.Pack(msgID, binary_data)
	if err != nil {
		log.Println("pack err:", err)
		return
	}
	this.conn.Write(sendData)
}

//处理服务器返回的数据,根据不同的消息ID对消息进行处理
func (this *TcpClient) DoMsg(msg *Message) {
	fmt.Printf("msgId:%d,data:%v\n", msg.MsgId, msg.Data)
	if msg.MsgId == 1 {
		//服务器回执给客户端：分配ID
		var msgID = &pb.SyncPid{}
		proto.Unmarshal(msg.Data, msgID)
		this.Pid = msgID.Pid

	} else if msg.MsgId == 200 {
		var broadcast = &pb.BroadCast{}
		proto.Unmarshal(msg.Data, broadcast)
		if broadcast.Tp == 2 && broadcast.Pid == this.Pid {
			this.X=broadcast.GetP().X
			this.Y=broadcast.GetP().Y
			this.Z=broadcast.GetP().Z
			this.V=broadcast.GetP().V
			fmt.Printf("Player ID: %d online.. at(%f,%f,%f,%f)\n",this.Pid,this.X,this.Y,this.Z,this.V)
			go func() {
				for{
					this.AIRobotAction()//自动完成一个AI动作
					time.Sleep(3*time.Second)
				}
			}()
		}else if broadcast.Tp==1{
			fmt.Println(fmt.Sprintf("世界聊天: 玩家:%d 说的话是 %s", broadcast.Pid, broadcast.GetContent()))
		}

	}
}

func (this *TcpClient)AIRobotAction(){
	//聊天 或者 移动

	tp := rand.Intn(2)

	if tp == 0 {
		//自动发送一个聊天信息
		content := fmt.Sprintf("hello 我是 player :%d, 你是谁！？", this.Pid)
		msg := &pb.Talk{
			Content:content,
		}

		//将数据发送给对应的服务端
		this.SendMsg(2, msg)

	} else {
		//自动移动
		x := this.X
		z := this.Z

		randpos := rand.Intn(2)
		if randpos == 0 {
			x -= float32(rand.Intn(10))
			z -= float32(rand.Intn(10))
		} else {
			x += float32(rand.Intn(10))
			z += float32(rand.Intn(10))
		}

		//纠正坐标
		if x > 410 {
			x = 410
		} else if x < 85 {
			x = 85
		}

		if z > 400 {
			z = 400
		} else if z < 75 {
			z = 75
		}

		randv := rand.Intn(2)
		v := this.V
		if randv == 0 {
			v = 25
		} else {
			v = 350
		}

		//打包一个proto结构
		msg := &pb.Position{
			X:x,
			Y:this.Y,
			Z:z,
			V:v,
		}

		fmt.Println("Player Id = ", this.Pid, " walking...")

		this.SendMsg(3, msg)
	}
}


func (this *TcpClient) Start() {
	go func() {
		for {
			fmt.Println("deal server msg read and write...")
			//根据zinx框架的LTV  先解析头部8个字节，再得到包体
			headData := make([]byte, 8)
			if _, err := io.ReadFull(this.conn, headData); err != nil {
				log.Println(err)
				return
			}
			msg, err := this.UnPack(headData)
			if err != nil {
				log.Println("start Unpack err", err)
				return
			}
			//解析数据部分
			if msg.Len > 0 {
				msg.Data = make([]byte, 512)
				if _, err := io.ReadFull(this.conn, msg.Data); err != nil {
					log.Println("io readfull err:", err)
				}
				//根据不同的消息部分对消息进行处理
				this.DoMsg(msg)
			}

		}
	}()
}
func main() {
	client := NewTcpClient("127.0.0.1", 8999)
	client.Start()
	//阻塞
	c := make(chan os.Signal, 1)
	//对当前os.singnal os.kill信号进行收集
	signal.Notify(c, os.Kill, os.Interrupt)
	sig := <-c
	fmt.Println("==>recv signal,", sig)
	return
}
