package handlers

import (
	"context"
	"time"

	m "github.com/dinofei/gomicro-server/models"
	pb "github.com/dinofei/gomicro-server/proto/user"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceHandler struct{}

func (u *UserServiceHandler) Create(c context.Context, req *pb.User, rsp *pb.Response) error {

	// 从微服务上下文中获取追踪信息
	md, ok := metadata.FromContext(c)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// 创建新的 Span 并将其绑定到微服务上下文
	sp = opentracing.StartSpan("user.create", opentracing.ChildOf(wireContext))
	// 记录请求
	sp.SetTag("req", req)
	defer func() {
		// 记录响应
		sp.SetTag("res", rsp)
		// 在函数返回 stop span 之前，统计函数执行时间
		sp.Finish()
	}()

	user := &m.User{}
	user.Username = req.Username
	bcrypPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user.Password = string(bcrypPass)
	user.Age = req.Age
	now := int32(time.Now().Unix())
	user.CreatedAt = now
	user.UpdatedAt = now
	newUser, err := user.InsertUser()
	rsp.User = new(pb.User)
	rsp.User.Username = newUser.Username
	rsp.User.Password = newUser.Password
	rsp.User.Age = newUser.Age
	return err
}

func (u *UserServiceHandler) Get(c context.Context, req *pb.User, rsp *pb.Response) error {
	user := &m.User{}
	user.Username = req.Username
	info, err := user.GetByUser()
	if err == nil {
		puser := pb.User{
			Username: info.Username,
			Password: info.Password,
			Age:      info.Age,
		}
		rsp.User = &puser
	}
	return err
}

func (u *UserServiceHandler) GetAll(c context.Context, req *pb.Request, rsp *pb.Response) error {
	user := &m.User{}
	list, err := user.ListUser()
	if err == nil {
		pusers := make([]*pb.User, 0)
		for _, v := range list {
			pusers = append(pusers, &pb.User{
				Username: v.Username,
				Password: v.Password,
				Age:      v.Age,
			})
		}
		rsp.Users = pusers
	}
	return err
}
