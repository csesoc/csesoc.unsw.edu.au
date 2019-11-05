package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPost - Retrieve a post from the database
func GetPost(collection *mongo.Collection, id int, category string) *Post {
	var result *Post

	// Search for post by id and category
	filter := bson.D{{"postID", id}, {"category", category}}
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
		filter := bson.D{{"post_category", cat}}
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

// NewPost - Add a new post
func NewPost(collection *mongo.Collection, id int, category int, showInMenu bool, title string, subtitle string, postType string, content string, github string, fb string) {
	post := Post{
		postID:           id,
		postTitle:        title,
		postSubtitle:     subtitle,
		postType:         postType,
		postCategory:     category,
		createdOn:        time.Now(),
		lastEditedOn:     time.Now(),
		postContent:      content,
		postLinkGithub:   github,
		postLinkFacebook: fb,
		showInMenu:       showInMenu,
	}

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdatePost - Update a post with new information
func UpdatePost(collection *mongo.Collection, id int, category int, showInMenu bool, title string, subtitle string, postType string, content string, github string, fb string) {
	filter := bson.D{{"postID", id}}
	update := bson.D{
		{"$set", bson.D{
			{"postTitle", title},
			{"postSubtitle", subtitle},
			{"postType", postType},
			{"postCategory", category},
			{"lastEditedOn", time.Now()},
			{"postContent", content},
			{"postLinkGithub", github},
			{"postLinkFacebook", fb},
			{"showinMenu", showInMenu},
		}},
	}

	// Find a post by id and update it
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// DeletePost - Delete a post from the database
func DeletePost(collection *mongo.Collection, id int) {
	filter := bson.D{{"postID", id}}

	// Find a post by id and delete it
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
