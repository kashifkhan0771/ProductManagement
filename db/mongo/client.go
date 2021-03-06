package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kashifkhan0771/ProductManagement/config"
	"github.com/kashifkhan0771/ProductManagement/db"
	domainErr "github.com/kashifkhan0771/ProductManagement/errors"
	"github.com/kashifkhan0771/ProductManagement/models"
)

const (
	productCollection = "products"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mongo database connection.
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/?connect=direct", viper.GetString(config.DBHost), viper.GetString(config.DBPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

func (c *client) AddProduct(product *models.Product) (string, error) {
	if product.ID != "" {
		return "", errors.New("id is not empty")
	}
	product.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(productCollection)
	if _, err := collection.InsertOne(context.TODO(), product); err != nil {
		return "", errors.Wrap(err, "failed to add product")
	}

	return product.ID, nil
}

func (c *client) GetProductByID(id string) (*models.Product, error) {
	var product *models.Product
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(productCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("product: %s not found", id))
		}
	}

	return product, nil
}

func (c *client) ListProducts(ctx context.Context, filter map[string]interface{}, sortBy map[string]bool, offset, limit int) ([]*models.Product, error) {
	productList := make([]*models.Product, 0)
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(productCollection)
	//filter[models.OffSelling] = bson.M{"$ne": models.OffSelling}
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(sortOrder(sortBy)).SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve list of tasks")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product *models.Product
		if err = cursor.Decode(&product); err != nil {
			return nil, err
		}
		productList = append(productList, product)
	}
	log().Infof("tasks returned: %d", len(productList))

	return productList, nil
}

func (c *client) DeleteProduct(id string) error {
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(productCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete product")
	}

	return nil
}

func (c *client) UpdateProduct(product *models.Product) error {
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(productCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": product.ID}, bson.M{"$set": product}); err != nil {
		return errors.Wrap(err, "failed to update product")
	}

	return nil
}

func sortOrder(spec map[string]bool) map[string]int {
	sort := make(map[string]int)
	for key, value := range spec {
		if value {
			sort[key] = 1
		}
	}

	return sort
}
