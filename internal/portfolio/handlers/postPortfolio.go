package portfolio

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sajeevany/portfolio-service/internal/datastore/as"
	"github.com/sajeevany/portfolio-service/pkg/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

//@Summary Creates portfolio a unique ID
//@Description Insert portfolio. Returns the portfolio ID.
//@Produce json
//@Success 200 {object} model.PortfolioID
//@Failure 404 {string} model.Error
//@Router /portfolio/:portfolioID [post]
//@Tags portfolio
func PostPortfolio(logger *logrus.Logger, client *aerospike.Client, setMetadata config.SetMD) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Bind body to portfolio object
		var portfolio model.PortfolioCreateModel
		if bErr := ctx.ShouldBindJSON(&portfolio); bErr != nil {
			logger.Errorf("Unable to bind request body to portfolio object %v", bErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": bErr.Error()})
			return
		}

		//Validate portfolio
		if valid, vErr := portfolio.IsValid(); (vErr != nil) || !valid {
			logger.WithFields(portfolio.GetFields()).Errorf("Input portfolio is invalid")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": vErr})
			return
		}

		//Get unused ID
		id, err := as.GetUniqueID(logger, client, setMetadata)
		if err != nil {
			logger.WithFields(setMetadata.GetFields()).Errorf("Unable to get unique id %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Infof("Generated ID %v", id)

		//Default response for testing
		response := model.AllPortfoliosViewModel{
			Portfolios: []model.PortfolioViewModel{{
				Metadata: model.MetadataViewModel{
					ID: "1", CreateTime: "2", LastUpdated: "3"},
				Stocks: []model.StockViewModel{}}},
		}
		ctx.JSON(http.StatusOK, response)
	}
}
