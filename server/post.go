package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPosts - Retrieve a post from the database
func GetPosts(collection *mongo.Collection, id int, category string) Post {
	var result Post

	// Search for post by id and category
	filter := bson.D{{Key: "postid", Value: id}, {Key: "category", Value: category}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// GetAllPosts - Retrieve all posts
func GetAllPosts(collection *mongo.Collection, count int, cat string) []*Post {
	findOptions := options.Find()
	if count != 10 {
		findOptions.SetLimit(int64(count))
	} else {
		findOptions.SetLimit(10)
	}

	var posts []*Post
	var cur *mongo.Cursor
	var err error

	if cat == "" { // No specified category
		cur, err = collection.Find(context.TODO(), bson.D{{}}, findOptions)
	} else {
		filter := bson.D{{Key: "postcategory", Value: cat}}
		cur, err = collection.Find(context.TODO(), filter, findOptions)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Iterate through all results
	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, &elem)
	}

	return posts
}

// NewPosts - Add a new post
func NewPosts(collection *mongo.Collection, id int, category int, showInMenu bool, title string, subtitle string, postType string, content string, github string, fb string) {
	currTime := time.Now()
	post := Post{
		PostID:           id,
		PostTitle:        title,
		PostSubtitle:     subtitle,
		PostType:         postType,
		PostCategory:     category,
		CreatedOn:        currTime.Unix(),
		LastEditedOn:     currTime.Unix(),
		PostContent:      content,
		PostLinkGithub:   github,
		PostLinkFacebook: fb,
		ShowInMenu:       showInMenu,
	}

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdatePosts - Update a post with new information
func UpdatePosts(collection *mongo.Collection, id int, category int, showInMenu bool, title string, subtitle string, postType string, content string, github string, fb string) {
	filter := bson.D{{Key: "postid", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "posttitle", Value: title},
			{Key: "postsubtitle", Value: subtitle},
			{Key: "posttype", Value: postType},
			{Key: "postcategory", Value: category},
			{Key: "lasteditedon", Value: time.Now()},
			{Key: "postcontent", Value: content},
			{Key: "postlinkgithub", Value: github},
			{Key: "postlinkfacebook", Value: fb},
			{Key: "showinmenu", Value: showInMenu},
		}},
	}

	// Find a post by id and update it
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// DeletePosts - Delete a post from the database
func DeletePosts(collection *mongo.Collection, id int) {
	filter := bson.D{{Key: "postid", Value: id}}

	// Find a post by id and delete it
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
