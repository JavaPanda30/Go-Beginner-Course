package routes

import (
	"net/http"
	"strconv"

	"example.com/eventbook/models"
	"github.com/gin-gonic/gin"
)

func GetEventsById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Events": events})
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	event.UserId = c.GetInt64("id")
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse given data", "error": err.Error()})
		return
	}

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created Successfully", "Event": event})
}

func UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse given data", "error": err.Error()})
		return
	}

	if err := event.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfully", "Event": event})
}

func DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found or could not be deleted"})
		return
	}
	err = models.DeleteEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found or could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event, "action": "Event Deleted Successfully"})
}

func AddRegistration(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found"})
		return
	}
	err = event.Register()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could Not Register", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registered User Successfully"})
}

func CancelRegistration(c *gin.Context) {
	userid := c.GetInt64("id")
	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	var event models.Event
	event.ID = eventid
	err = event.Deregister(userid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could Not Register", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "De-Registered User Successfully"})
}
