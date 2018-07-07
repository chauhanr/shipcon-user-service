package main

import (
	"context"
	pb "github.com/chauhanr/shipcon-user-service/proto/user"
	"log"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"fmt"
)
type service struct{
	repo Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error{
	user, err := srv.repo.Get(req.Id)
	if err != nil{
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error{
	users, err := srv.repo.GetAll()
	if err != nil{
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error{
	log.Println("Logging in with: ", req.Email, req.Password)
	reqPass := req.Password
	user,err := srv.repo.GetByEmailAndPassword(req)
	if err != nil{
		return err
	}
	// compare password with hashed password
	// log.Printf("User password %s, DB password %s", user.Password, reqPass)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqPass)); err != nil{
		return err
	}
	token, err := srv.tokenService.Encode(user)
	if err != nil{
		return err
	}
	res.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	log.Printf("Creating user %v\n", req)
	// get hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		return errors.New(fmt.Sprintf("error hashing password %v", err))
	}
	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil{
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := srv.tokenService.Encode(req)
	if err != nil{
		return err
	}

	res.User = req
	res.Token = &pb.Token{Token: token}
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}