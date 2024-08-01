package devicehandler

import (
	"Api-GateWay/protos/deviceproto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	ClientDevice deviceproto.DeviceServiceClient
}

// CreateDevice godoc
// @Summary Create a new device
// @Description Create a new device
// @Tags device
// @Accept json
// @Produce json
// @Param device body deviceproto.DeviceRequest true "Device request body"
// @Success 200 {object} deviceproto.DeviceResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /device [post]
func (d *DeviceHandler) CreateDevice(c *gin.Context) {
	var req deviceproto.DeviceRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := d.ClientDevice.CreateDevice(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetDeviceById godoc
// @Summary Get a device by ID
// @Description Get a device by ID
// @Tags device
// @Accept json
// @Produce json
// @Param id path string true "Device ID"
// @Success 200 {object} deviceproto.Device
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /device/{id} [get]
func (d *DeviceHandler) GetDeviceById(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := d.ClientDevice.GetDeviceById(ctx, &deviceproto.DeviceResponse{DeviceId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateDevice godoc
// @Summary Update a device
// @Description Update a device
// @Tags device
// @Accept json
// @Produce json
// @Param device body deviceproto.Device true "Device update request body"
// @Success 200 {object} deviceproto.DeviceOperationResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /device [put]
func (d *DeviceHandler) UpdateDevice(c *gin.Context) {
	var req deviceproto.Device
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := d.ClientDevice.UpdateDevice(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteDevice godoc
// @Summary Delete a device
// @Description Delete a device
// @Tags device
// @Accept json
// @Produce json
// @Param device body deviceproto.DeviceResponse true "Delete device request body"
// @Success 200 {object} deviceproto.DeviceOperationResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /device [delete]
func (d *DeviceHandler) DeleteDevice(c *gin.Context) {
	var req deviceproto.DeviceResponse
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := d.ClientDevice.DeleteDevice(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// CreateCommand godoc
// @Summary Create a new device
// @Description Create a new device
// @Tags device
// @Accept json
// @Produce json
// @Param device body deviceproto.CommandRequest true "Device request body"
// @Success 200 {object} deviceproto.CommandRequest
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /Command [post]
func (d *DeviceHandler) CreateCommand(c *gin.Context){
	var req deviceproto.CommandRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	
	resp, err := d.ClientDevice.CreateCommand(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
