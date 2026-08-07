package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelDAO struct {
	ID   int64  `bson:"id"`
	Name string `bson:"name"`
	City string `bson:"city"`
}

type HotelsRepo interface {
	GetHotelByID(ctx context.Context, id int64) (HotelDAO, error)
}

type HotelsMongoDB struct {
	client *mongo.Client
}

func (repo HotelsMongoDB) GetHotelByID(ctx context.Context, id int64) (HotelDAO, error) {
	var hotel HotelDAO

	err := repo.client.
		Database("hotel-platform").
		Collection("hotels").
		FindOne(ctx, bson.M{"id": id}).
		Decode(&hotel)
	if err != nil {
		return HotelDAO{}, err
	}

	return hotel, nil
}

type HotelsMock struct{}

func (repo HotelsMock) GetHotelByID(ctx context.Context, id int64) (HotelDAO, error) {
	return HotelDAO{
		ID:   id,
		Name: "Mock Hotel",
		City: "Cordoba",
	}, nil
}

func LoadHotel(ctx context.Context, repo HotelsRepo, id int64) error {
	hotel, err := repo.GetHotelByID(ctx, id)
	if err != nil {
		return err
	}

	fmt.Printf("hotel: %+v\n", hotel)
	return nil
}
