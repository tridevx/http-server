package router

import (
	"fmt"
	"net/http"
)

type Route struct {
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Get(pattern string, controller interface{}, methodName string) {
	r.addRoute("GET", pattern, controller, methodName)
}

func (r *Router) Post(pattern string, controller interface{}, methodName string) {
	r.addRoute("POST", pattern, controller, methodName)
}

func (r *Router) Put(pattern string, controller interface{}, methodName string) {
	r.addRoute("PUT", pattern, controller, methodName)
}

func (r *Router) Patch(pattern string, controller interface{}, methodName string) {
	r.addRoute("PATCH", pattern, controller, methodName)
}

func (r *Router) Delete(pattern string, controller interface{}, methodName string) {
	r.addRoute("DELETE", pattern, controller, methodName)
}

func (r *Router) Resource(pattern string, controller interface{}) {
	r.Get(pattern, controller, "index")
	r.Post(pattern, controller, "store")
	r.Get(fmt.Sprintf("%s/:id", pattern), controller, "show")
	r.Put(fmt.Sprintf("%s/:id", pattern), controller, "update")
	r.Delete(fmt.Sprintf("%s/:id", pattern), controller, "destroy")
}

func (r *Router) addRoute(method, pattern string, controller interface{}, methodName string) {
	key := fmt.Sprintf("%s-%s", pattern, method)
	r.routes[key] = func(w http.ResponseWriter, req *http.Request) {
		callControllerMethod(controller, methodName, w, req)
	}
}


func callControllerMethod(controller interface{}, methodName string, w http.ResponseWriter, req *http.Request) {
	controllerValue := reflect.ValueOf(controller)
	method := controllerValue.MethodByName(methodName)

	if !method.IsValid() {
		fmt.Fprintf(w, "Invalid controller method: %s", methodName)
		return
	}

	method.Interface().(func(http.ResponseWriter, *http.Request))(w, req)
}