package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"sen1or/letslive/user/dto"
	servererrors "sen1or/letslive/user/errors"
	"sen1or/letslive/user/pkg/logger"
	"sen1or/letslive/user/services"
	"sen1or/letslive/user/types"
	"strconv"

	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	ErrorHandler
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUserByIdPublicHandler(w http.ResponseWriter, r *http.Request) {
	authenticatedUserId, _ := getUserIdFromCookie(r)
	userId := r.PathValue("userId")
	if len(userId) == 0 {
		h.WriteErrorResponse(w, servererrors.ErrInvalidPath)
		return
	}

	userUUID, err := uuid.FromString(userId)
	if err != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidInput)
		return
	}

	user, serviceErr := h.userService.GetUserById(userUUID, authenticatedUserId)
	if serviceErr != nil {
		h.WriteErrorResponse(w, serviceErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsersPublicHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil || page < 0 {
		h.WriteErrorResponse(w, servererrors.ErrInvalidInput)
		return
	}

	users, serviceErr := h.userService.GetAllUsers(page)
	if serviceErr != nil {
		h.WriteErrorResponse(w, serviceErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) SearchUsersPublicHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	users, err := h.userService.SearchUserByUsername(username)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByStreamAPIKeyInternalHandler(w http.ResponseWriter, r *http.Request) {
	streamAPIKeyString := r.URL.Query().Get("streamAPIKey")
	if len(streamAPIKeyString) == 0 {
		h.WriteErrorResponse(w, servererrors.ErrUnauthorized)
		return
	}

	streamAPIKey, err := uuid.FromString(streamAPIKeyString)
	if err != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidInput)
		return
	}

	user, err := h.userService.GetUserByStreamAPIKey(streamAPIKey)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetCurrentUserPrivateHandler(w http.ResponseWriter, r *http.Request) {
	userUUID, cookieErr := getUserIdFromCookie(r)
	if cookieErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrUnauthorized)
		return
	}
	user, err := h.userService.GetUserFullInformation(*userUUID)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// INTERNAL
func (h *UserHandler) CreateUserInternalHandler(w http.ResponseWriter, r *http.Request) {
	var body dto.CreateUserRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidPayload)
		return
	}

	createdUser, err := h.userService.CreateNewUser(body)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&createdUser)
}

// INTERNAL
func (h *UserHandler) SetUserVerifiedInternalHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	if len(userId) == 0 {
		h.WriteErrorResponse(w, servererrors.ErrInvalidInput)
		return
	}

	userUUID, err := uuid.FromString(userId)
	if err != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidInput)
		return
	}

	if err := h.userService.UpdateUserVerified(userUUID); err != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidInput)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) UpdateCurrentUserPrivateHandler(w http.ResponseWriter, r *http.Request) {
	userUUID, cookieErr := getUserIdFromCookie(r)
	if cookieErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrUnauthorized)
		return
	}
	defer r.Body.Close()

	var requestBody dto.UpdateUserRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidPayload)
		return
	}

	requestBody.Id = uuid.FromStringOrNil(userUUID.String())
	updatedUser, err := h.userService.UpdateUser(requestBody)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func (h *UserHandler) GenerateNewAPIStreamKeyPrivateHandler(w http.ResponseWriter, r *http.Request) {
	userUUID, cookieErr := getUserIdFromCookie(r)
	if cookieErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrUnauthorized)
		return
	}
	defer r.Body.Close()

	newKey, err := h.userService.UpdateUserAPIKey(*userUUID)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newKey))
}

func (h *UserHandler) UpdateUserProfilePicturePrivateHandler(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 10 * 1024 * 1024
	userUUID, cookieErr := getUserIdFromCookie(r)
	if cookieErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrUnauthorized)
		return
	}
	defer r.Body.Close()

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	if err := r.ParseMultipartForm(0); err != nil {
		var maxByteError *http.MaxBytesError
		if errors.As(err, &maxByteError) {
			h.WriteErrorResponse(w, servererrors.ErrImageTooLarge)
			return
		}

		h.WriteErrorResponse(w, servererrors.ErrInvalidPayload)
		return
	}

	file, fileHeader, formErr := r.FormFile("profile-picture")
	if formErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidPayload)
		return
	}

	savedPath, err := h.userService.UpdateUserProfilePicture(file, fileHeader, *userUUID)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(savedPath))
}

func (h *UserHandler) UpdateUserBackgroundPicturePrivateHandler(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 10 * 1024 * 1024
	userUUID, cookieErr := getUserIdFromCookie(r)
	if cookieErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrUnauthorized)
		return
	}
	defer r.Body.Close()

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	if err := r.ParseMultipartForm(0); err != nil {
		var maxByteError *http.MaxBytesError
		if errors.As(err, &maxByteError) {
			h.WriteErrorResponse(w, servererrors.ErrImageTooLarge)
			return
		}

		h.WriteErrorResponse(w, servererrors.ErrInternalServer)
		return
	}

	file, fileHeader, formErr := r.FormFile("background-picture")
	if formErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrInternalServer)
		return
	}

	savedPath, err := h.userService.UpdateUserBackgroundPicture(file, fileHeader, *userUUID)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(savedPath))
}

func (h *UserHandler) UpdateUserInternalHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	defer r.Body.Close()

	if len(userID) == 0 {
		h.WriteErrorResponse(w, servererrors.ErrInvalidPath)
		return
	}

	var requestBody dto.UpdateUserRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidPayload)
		return
	}
	requestBody.Id = uuid.FromStringOrNil(userID)

	updatedUser, err := h.userService.UpdateUserInternal(requestBody)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func getUserIdFromCookie(r *http.Request) (*uuid.UUID, *servererrors.ServerError) {
	accessTokenCookie, err := r.Cookie("ACCESS_TOKEN")
	if err != nil || len(accessTokenCookie.Value) == 0 {
		logger.Debugf("missing credentials")
		return nil, servererrors.ErrUnauthorized
	}

	myClaims := types.MyClaims{}

	// the signature should already been checked from the api gateway before going to this
	_, _, err = jwt.NewParser().ParseUnverified(accessTokenCookie.Value, &myClaims)
	if err != nil {
		logger.Debugf("invalid access token: %s", err)
		return nil, servererrors.ErrUnauthorized
	}

	userUUID, err := uuid.FromString(myClaims.UserId)
	if err != nil {
		logger.Debugf("userId not valid")
		return nil, servererrors.ErrUnauthorized
	}

	return &userUUID, nil
}

func (h *UserHandler) UploadSingleFileToMinIOHandler(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 10 * 1024 * 1024
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(0); err != nil {
		var maxByteError *http.MaxBytesError
		if errors.As(err, &maxByteError) {
			h.WriteErrorResponse(w, servererrors.ErrImageTooLarge)
			return
		}

		h.WriteErrorResponse(w, servererrors.ErrInvalidPayload)
		return
	}

	file, fileHeader, formErr := r.FormFile("file")
	if formErr != nil {
		h.WriteErrorResponse(w, servererrors.ErrInvalidPayload)
		return
	}

	savedPath, err := h.userService.UploadFileToMinIO(file, fileHeader)
	if err != nil {
		h.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(savedPath))
}
