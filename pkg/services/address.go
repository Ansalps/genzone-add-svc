package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Ansalps/genzone-add-svc/pkg/db"
	"github.com/Ansalps/genzone-add-svc/pkg/responsemodels"
	"github.com/Ansalps/genzone-add-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedAddressServiceServer
}

func (s *Server) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.CreateAddressResponse, error) {
	log.Printf("Received CreateAddress request: %+v\n", req)
	var address responsemodels.Address
	// u64, err := strconv.ParseUint(req.Userid, 10, 32) // base 10, 32-bit integer
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	//u := uint(u64) // Convert uint64 to uint if needed
	address.UserID = req.Userid
	address.Country = req.Country
	address.State = req.State
	address.District = req.District
	address.StreetName = req.Streetname
	address.PinCode = req.Pincode
	address.Phone = req.Phone
	address.Default = req.Default
	if result := s.H.DB.Create(&address); result.Error != nil {
		return &pb.CreateAddressResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.CreateAddressResponse{
		Status: http.StatusCreated,
	}, nil
}
func (s *Server) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error) {
	log.Printf("Received CreateAddress request: %+v\n", req)
	var address []responsemodels.Address
	u64, err := strconv.ParseUint(req.Userid, 10, 32) // base 10, 32-bit integer
	if err != nil {
		fmt.Println("Error:", err)
	}
	u := uint(u64) // Convert uint64 to uint if needed
	if result := s.H.DB.Where("user_id=?", u).Find(&address); result.Error != nil {
		fmt.Println(result)
		return &pb.GetAddressResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	 // Convert addresses to the response format
	 var response pb.GetAddressResponse
	 for _, addr := range address {
		// Create a new pb.Address instance and map the fields
		pbAddr := &pb.Address{
			// Assuming Address fields in responsemodels.Address map to pb.Address fields
			Id: int64(addr.ID),
			UserId: addr.UserID,
			Country: addr.Country,
			State: addr.State,
			District: addr.District,
			StreetName: addr.StreetName,
			Pincode: addr.PinCode,
			Phone: addr.Phone,
			Default: addr.Default,
			// Add any other fields you need to map
		}
		response.Addresses = append(response.Addresses, pbAddr)
	}
	return &pb.GetAddressResponse{
		Status: http.StatusOK,
		Addresses: response.Addresses,
	}, nil
}
