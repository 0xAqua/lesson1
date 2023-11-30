// //This package contains all controllers used in the application
package controllers

// )

// Import the necessary packages to make an HTTP request, manipulate strings, and track time.
import (
	"net/http"
	"strings"
	"time"

// // Importing packages for use in the project:
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

// // This code imports all the necessary packages for the backend application.
	"backend/pkg/errors"
	"backend/pkg/models"
	"backend/pkg/utils"
)

// // Create a filter to get only documents with a valid status
	filter := bson.D{{"status", bson.D{{"$ne", "deleted"}}}}

	// Retrieve all documents from the
func ListTodos(c *gin.Context, collection *mongo.Collection) {
	ctx, cancel := utils.CreateContext(5)
	defer cancel()

// // Check for errors when fetching todos from the collection
	todoList, err := utils.FetchTodos(ctx, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errors.ErrFetchTodos})
		return
	}

// // Send the response with a status code of 200 and the list of todos as a JSON object
	c.JSON(http.StatusOK, gin.H{"todos": todoList})
}

// if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
func AddTodo(c *gin.Context, collection *mongo.Collection) {
	var todo models.ToDoItem

// // Check if the JSON data is valid before binding it to the "todo" variable
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors.ErrInvalidJSON})
		return
	}

// // Sets the CreatedAt field to the current time

// Sets the CreatedAt field to the current time
	todo.CreatedAt = time.Now()

// // Create a context with a timeout of 5 seconds
	ctx, cancel := utils.CreateContext(5)
	defer cancel()

// // Insert the ToDoItem into the database
	// ToDoItemをデータベースに追加
	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errors.ErrServerError})
		return
	}

// // Send a JSON response to the client indicating that the todo was added successfully
	c.JSON(http.StatusCreated, gin.H{"message": "Todo added successfully"})
}

// // Delete the item
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.
func DeleteTodo(c *gin.Context, collection *mongo.Collection) {
	itemTodoDelete := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(itemTodoDelete)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors.ErrInvalidIDFormat})
		return
	}

// // Create a context with a timeout of 5 seconds
	ctx, cancel := utils.CreateContext(5)
	defer cancel()

// // Check for errors when deleting a todo by ID, and handle any errors that occur
	if err := utils.DeleteTodoByID(ctx, collection, objID); err != nil {
		utils.HandleMongoError(c, err)
		return
	}

// //Todo deleted successfully message returned with status OK
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

// // Get the data from the body
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest,
func UpdateTodo(c *gin.Context, collection *mongo.Collection) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors.ErrInvalidIDFormat})
		return
	}

// // Check if the request body is valid and bind it to the todo model, otherwise return an error status code
	var todo models.ToDoItem
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors.ErrInvalidRequestBody})
		return
	}

// // Check if the todo text is empty before creating a new todo
	if strings.TrimSpace(todo.Text) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors.ErrTodoTextEmpty})
		return
	}

// // Update the todo item's timestamp to the current time
	todo.UpdateAt = time.Now()

// // Create a context with a timeout of 5 seconds
	ctx, cancel := utils.CreateContext(5)
	defer cancel()

// // Update the todo document in the collection, return an error if it fails
	if err := utils.UpdateTodoByID(ctx, collection, objID, todo.Text, todo.Deadline, todo.Priority, todo.UpdateAt); err != nil {
		utils.HandleMongoError(c, err)
		return
	}

// // Send a JSON response with a message indicating that the Todo was successfully updated
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}


