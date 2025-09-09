package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/oatsmoke/20250905/internal/domain"
	"github.com/oatsmoke/20250905/internal/lib/err_msg"
	"github.com/oatsmoke/20250905/internal/lib/logger"
	"github.com/oatsmoke/20250905/internal/model"
)

type SubscriptionHandler struct {
	subscriptionService domain.Subscription
}

func NewSubscriptionHandler(subscriptionService domain.Subscription) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

// Create
// @Summary Create subscription
// @Tags subscription
// @Produce json
// @Param request body model.Subscription true "Subscription"
// @Success 201
// @Failure 400 {object} string "bad request"
// @Failure 405 {object} string "method not allowed"
// @Failure 500 {object} string "internal server error"
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.HttpError(w, err_msg.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	subscription := new(model.Subscription)
	if err := json.NewDecoder(r.Body).Decode(subscription); err != nil {
		logger.HttpError(w, err, http.StatusBadRequest)
		return
	}

	if *subscription == (model.Subscription{}) {
		logger.HttpError(w, err_msg.RequestBodyIsEmpty, http.StatusBadRequest)
		return
	}

	if err := h.subscriptionService.Create(r.Context(), subscription); err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Read
// @Summary Read subscription
// @Tags subscription
// @Produce json
// @Param id path int true "ID subscription"
// @Success 200 {object} model.Subscription
// @Failure 400 {object} string "bad request"
// @Failure 405 {object} string "method not allowed"
// @Failure 500 {object} string "internal server error"
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.HttpError(w, err_msg.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/subscriptions/"), 10, 64)
	if err != nil {
		logger.HttpError(w, err, http.StatusBadRequest)
		return
	}

	subscription, err := h.subscriptionService.Read(r.Context(), id)
	if err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(subscription); err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}
}

// Update
// @Summary Update subscription
// @Tags subscription
// @Produce json
// @Param id path int true "ID subscription"
// @Param request body model.Subscription true "Subscription"
// @Success 204
// @Failure 400 {object} string "bad request"
// @Failure 405 {object} string "method not allowed"
// @Failure 500 {object} string "internal server error"
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		logger.HttpError(w, err_msg.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/subscriptions/"), 10, 64)
	if err != nil {
		logger.HttpError(w, err, http.StatusBadRequest)
		return
	}

	subscription := new(model.Subscription)
	if err := json.NewDecoder(r.Body).Decode(subscription); err != nil {
		logger.HttpError(w, err, http.StatusBadRequest)
		return
	}

	if *subscription == (model.Subscription{}) {
		logger.HttpError(w, err_msg.RequestBodyIsEmpty, http.StatusBadRequest)
		return
	}

	subscription.ID = id
	if err := h.subscriptionService.Update(r.Context(), subscription); err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete
// @Summary Delete subscription
// @Tags subscription
// @Produce json
// @Param id path int true "ID subscription"
// @Success 204
// @Failure 400 {object} string "bad request"
// @Failure 405 {object} string "method not allowed"
// @Failure 500 {object} string "internal server error"
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		logger.HttpError(w, err_msg.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/subscriptions/"), 10, 64)
	if err != nil {
		logger.HttpError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.subscriptionService.Delete(r.Context(), id); err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// List
// @Summary List subscriptions
// @Tags subscription
// @Produce json
// @Success 200 {array} model.Subscription
// @Failure 405 {object} string "method not allowed"
// @Failure 500 {object} string "internal server error"
// @Router /subscriptions [get]
func (h *SubscriptionHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.HttpError(w, err_msg.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	list, err := h.subscriptionService.List(r.Context())
	if err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}
}

// Total
// @Summary Total subscriptions
// @Tags subscription
// @Produce json
// @Param user_id query string true "user ID"
// @Param service_name query string true "service name"
// @Param start_date query string true "start date"
// @Param end_date query string true "end date"
// @Success 200 {object} int
// @Failure 400 {object} string "bad request"
// @Failure 405 {object} string "method not allowed"
// @Failure 500 {object} string "internal server error"
// @Router /subscriptions/total [get]
func (h *SubscriptionHandler) Total(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.HttpError(w, err_msg.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	userId := query.Get("user_id")
	serviceName := query.Get("service_name")
	startDate := query.Get("start_date")
	endDate := query.Get("end_date")

	if userId == "" || serviceName == "" || startDate == "" || endDate == "" {
		logger.HttpError(w, err_msg.RequestBodyIsEmpty, http.StatusBadRequest)
		return
	}

	total, err := h.subscriptionService.Total(r.Context(), &model.Subscription{
		UserId:      userId,
		ServiceName: serviceName,
		StartDate:   startDate,
		EndDate:     endDate,
	})
	if err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(total); err != nil {
		logger.HttpError(w, err, http.StatusInternalServerError)
		return
	}
}
