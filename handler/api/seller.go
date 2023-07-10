package api

// import (
// 	"kmipn-2023/model"
// 	"kmipn-2023/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type SellerAPI interface {
// 	Register(c *gin.Context)
// 	Login(c *gin.Context)
// }

// type sellerAPI struct {
// 	sellerService service.SellerService
// }

// func NewSellerAPI(sellerService service.SellerService) *sellerAPI {
// 	return &sellerAPI{sellerService}
// }

// func (u *sellerAPI) Register(c *gin.Context) {
// 	var seller model.SellerRegister

// 	if err := c.BindJSON(&seller); err != nil {
// 		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
// 		return
// 	}

// 	if seller.Email == "" || seller.Password == "" || seller.Username == "" {
// 		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
// 		return
// 	}

// 	var recordSeller = model.Seller{
// 		Username: seller.Username,
// 		Email:    seller.Email,
// 		Password: seller.Password,
// 	}

// 	recordSeller, err := u.sellerService.Register(&recordSeller)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
// 		return
// 	}

// 	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
// }

// func (u *sellerAPI) Login(c *gin.Context) {
// 	var seller model.SellerLogin

// 	if err := c.BindJSON(&seller); err != nil {
// 		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
// 		return
// 	}

// 	if seller.Email == "" || seller.Password == "" {
// 		c.JSON(http.StatusBadRequest, model.NewErrorResponse("email or password is empty"))
// 		return
// 	}

// 	loginSeller := &model.Seller{
// 		Email:    seller.Email,
// 		Password: seller.Password,
// 	}

// 	token, err := u.sellerService.Login(loginSeller)
// 	if err != nil {
// 		if err.Error() == "user not found" || err.Error() == "wrong email or password" {
// 			c.JSON(http.StatusBadRequest, model.NewErrorResponse("email or password is incorrect"))
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
// 		return
// 	}

// 	c.SetCookie("session_token", *token, 3600, "/", "localhost", false, true)

// 	c.JSON(http.StatusOK, gin.H{
// 		"seller_id": loginSeller.ID,
// 		"message":   "login success",
// 	})
// }
