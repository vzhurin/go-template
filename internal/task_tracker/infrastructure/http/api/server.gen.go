package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (PUT /task_tracker/{taskID}/estimate)
	EstimateTask(w http.ResponseWriter, r *http.Request, taskID UUID)

	// (POST /tasks)
	CreateTask(w http.ResponseWriter, r *http.Request)

	// (GET /tasks/{taskID})
	GetTask(w http.ResponseWriter, r *http.Request, taskID UUID)

	// (PUT /tasks/{taskID})
	UpdateTask(w http.ResponseWriter, r *http.Request, taskID UUID)

	// (PUT /tasks/{taskID}/assignParticipant)
	AssignParticipantToTask(w http.ResponseWriter, r *http.Request, taskID UUID)

	// (PUT /tasks/{taskID}/transitStatus)
	TransitTaskStatus(w http.ResponseWriter, r *http.Request, taskID UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// EstimateTask operation middleware
func (siw *ServerInterfaceWrapper) EstimateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if ctx == nil {
		
	}

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.EstimateTask(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTask(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetTask operation middleware
func (siw *ServerInterfaceWrapper) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTask(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateTask operation middleware
func (siw *ServerInterfaceWrapper) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateTask(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AssignParticipantToTask operation middleware
func (siw *ServerInterfaceWrapper) AssignParticipantToTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AssignParticipantToTask(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// TransitTaskStatus operation middleware
func (siw *ServerInterfaceWrapper) TransitTaskStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TransitTaskStatus(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/task_tracker/{taskID}/estimate", wrapper.EstimateTask)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/tasks", wrapper.CreateTask)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/tasks/{taskID}", wrapper.GetTask)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/tasks/{taskID}", wrapper.UpdateTask)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/tasks/{taskID}/assignParticipant", wrapper.AssignParticipantToTask)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/tasks/{taskID}/transitStatus", wrapper.TransitTaskStatus)
	})

	return r
}
