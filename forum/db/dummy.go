package db

import (
	"db-test/models"
	"fmt"
	"math/rand"
	"strings"

	"github.com/go-faker/faker/v4"
)

func GetRandomTitle() string {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	
	p := faker.Paragraph()
	return strings.Join(strings.Split(p, " ")[1:rand.Intn(9)], " ")
}

func GeneratePosts() {
	for i := 0; i < 1000; i++ {
		title := GetRandomTitle()
		if title == "" {
			continue
		}
		
		err := CreatePost(models.Post{
			Title:      title,
			Content:    faker.Paragraph(),
			UserID:     rand.Intn(100),
			Categories: SelectRandomCatag(),
		})
		if err != nil {
			continue
		}
	}
}


func SelectRandomCatag() []models.Category {
	categoryMap := make(map[int]bool)
	result := []models.Category{}
	numCategories := rand.Intn(5) + 1 // Ensure at least one category

	for i := 0; i < numCategories; i++ {
		catID := rand.Intn(10)
		if !categoryMap[catID] {
			categoryMap[catID] = true
			result = append(result, models.Category{
				ID: catID,
			})
		}
	}

	return result
}

func GenerateCatagories() {
	CreateCategory(1, "sport", "sport category")
	CreateCategory(1, "general", "general category")
	CreateCategory(1, "technology", "technology category")
	CreateCategory(1, "entertainment", "entertainment category")
	CreateCategory(1, "health", "health category")
	CreateCategory(1, "business", "business category")
	CreateCategory(1, "science", "science category")
	CreateCategory(1, "education", "education category")
	CreateCategory(1, "travel", "travel category")
}

func GenerateComments() {
	for range 1000 {
		// daEmail := fmt.Sprintf("%v-%v@gmail.com", faker.FirstName(), rand.Intn(1000))
		_, err := CreateComment(models.Comment{
			UserID:  rand.Intn(100),
			PostID:  rand.Intn(100),
			Content: faker.Paragraph(),
		})
		// _, err := CreateUser(faker.FirstName(), daEmail, faker.Password())
		if err != nil {
			continue
		}
	}
}

func GenerateCommentInteractions() {
	for range 1000 {
		_, err := CreateCommentInteraction(models.Interaction(rand.Intn(2)), rand.Intn(100), rand.Intn(100))
		if err != nil {
			continue
		}
	}
}

func GeneratePostInteractions() {
	for range 1000 {
		err := CreatePostInteraction(models.Vote(rand.Intn(2)), rand.Intn(100), rand.Intn(100))
		if err != nil {
			continue
		}
	}
}

func GenerateUsers() {
	for range 1000 {
		daEmail := fmt.Sprintf("%v-%v@gmail.com", faker.FirstName(), rand.Intn(1000))
		_, err := CreateUser(faker.FirstName(), daEmail, faker.Password())
		if err != nil {
			continue
		}
	}
}

func GenerateDummyData() {
	GenerateUsers()
	fmt.Println("User added")

	GenerateCatagories()
	fmt.Println("Categories generated")

	GeneratePosts()
	fmt.Println("Posts generated")

	GenerateComments()
	fmt.Println("Comments generated")

	GeneratePostInteractions()
	fmt.Println("Post interactions generated")

	GenerateCommentInteractions()
	fmt.Println("Comment interactions generated")
}
