package app

import (
	"net/http"
	"strconv"
	"time"
	"twitter/internal/model"
	"twitter/pkg/bcrypt"
	"twitter/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func (s *Server) SignupHandler(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	u, err := s.userService.GetUser(user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if u != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "username is already busy",
		})
	}

	hash, err := bcrypt.Generate(user.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	user.Password = hash

	id, err := s.userService.CreateUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	t, err := getToken(id, user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   t,
	})
}

func (s *Server) SigninHandler(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	u, err := s.userService.GetUser(user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if u == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "user not found",
		})
	}

	err = bcrypt.Compare(u.Password, user.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	t, err := getToken(u.Id, u.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   t,
	})
}

func getToken(id int, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(config.SECRET))
	return t, err
}

func (s *Server) GetUsersPostsHandler(c *fiber.Ctx) error {
	userIdString := c.Params("userId")
	userId, _ := strconv.Atoi(userIdString)

	posts, err := s.postsService.GetUsersPosts(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"posts": posts,
	})
}

func (s *Server) GetPostHandler(c *fiber.Ctx) error {
	userIdString := c.Params("userId")
	userId, _ := strconv.Atoi(userIdString)
	postIdString := c.Params("postId")
	postId, _ := strconv.Atoi(postIdString)

	post, err := s.postsService.GetPost(userId, postId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"post": post,
	})
}

func (s *Server) CreatePostHandler(c *fiber.Ctx) error {
	post := new(model.Posts)

	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err := s.postsService.CreatePosts(post)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (s *Server) UpdatePostHandler(c *fiber.Ctx) error {
	post := new(model.Posts)

	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err := s.postsService.UpdatePost(post)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (s *Server) DeletePostHandler(c *fiber.Ctx) error {
	postIdString := c.Params("postId")
	postId, _ := strconv.Atoi(postIdString)

	err := s.postsService.DeletePost(postId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
