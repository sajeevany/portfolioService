package endpoints

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/config"
	portfolio "github.com/sajeevany/portfolio-service/internal/portfolio/handlers"
	"github.com/sirupsen/logrus"
)

const (
	PortfolioGroup      = "/portfolio"
	GetAllPortfolios_   = "/"
	PostPortfolio       = "/"
	GetPortfolio        = "/:portfolioID"
	DeletePortfolio     = "/:portfolioID"
	AddStock            = "/:portfolioID/:ticker"
	ReplaceStockEntry   = "/:portfolioID/:inventoryID"
	ReplaceStockEntries = "/:portfolioID"
	DeleteStockEntry    = "/:portfolioID/:inventoryID"
	DeleteStockEntries  = "/:portfolioID"
)

//BuildGetAllPortfoliosEndpoint - Returns all portfolios
func BuildGetAllPortfoliosEndpoint(logger *logrus.Logger, asClient *aerospike.Client, setMetadata config.SetMD, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      GetAllPortfolios_,
		Handlers: append(handlers, portfolio.GetAllPortfolios(logger, asClient, setMetadata)),
	}
}

//BuildPostPortfolioEndpoint - Inserts new portfolio with a unique key
func BuildPostPortfolioEndpoint(logger *logrus.Logger, asClient *aerospike.Client, setMetadata config.SetMD, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      PostPortfolio,
		Handlers: append(handlers, portfolio.PostPortfolio(logger, asClient, setMetadata)),
	}
}