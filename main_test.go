package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var clientset *kubernetes.Clientset

func TestRootHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp Response
	data, err := ioutil.ReadAll(rr.Body)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		t.Fatalf("Test failed\n")
	}

	if resp.Status != "OK" {
		t.Errorf("Response Code error: %s\n", resp.Status)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Response Status error: %v\n", resp.StatusCode)
	}
	if resp.Body.Message != "connected to scout" {
		t.Errorf("Response Body error: %v\n", resp.Body.Message)
	}
}

func TestGetClientsetExternal(t *testing.T) {
	clientset = getClientsetExternal()
	clientsetDefault := &kubernetes.Clientset{}

	if reflect.TypeOf(clientsetDefault) != reflect.TypeOf(clientset) {
		t.Errorf("Not equal typeOf")
	}
}

func TestGetPods(t *testing.T) {
	podListDefault := &v1.PodList{}
	podList := getPods(clientset)

	if reflect.TypeOf(podListDefault) != reflect.TypeOf(podList) {
		t.Errorf("Not equal typeOf")
	}
}

func TestCreateRequestObjInsideCluster(t *testing.T) {
	pod := v1.Pod{}
	req := createRequestObj(true, pod)
	if req.Host == "localhost:8080" {
		t.Errorf("Host error: %v\n", req.Host)
	}
}

func TestCreateRequestObjOutsideCluster(t *testing.T) {
	pod := &v1.Pod{}
	req := createRequestObj(false, *pod)
	if req.Host != "localhost:8080" {
		t.Errorf("Host error: %v\n", req.Host)
	}
}

func TestCreateRequestObject(t *testing.T) {
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
			Labels:    map[string]string{"foo": "bar"},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Name:  "nginx",
				Image: "nginx",
			},
			},
		},
		Status: v1.PodStatus{
			PodIP: "1.2.3.4",
		},
	}
	req := createRequestObj(false, *pod)
	if req.Host != "localhost:8080" {
		t.Errorf("Host error: %v\n", req.Host)
	}
}

func TestMakeRequest(t *testing.T) {
	handlers()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
			Labels:    map[string]string{"foo": "bar"},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Name:  "nginx",
				Image: "nginx",
			},
			},
		},
		Status: v1.PodStatus{
			PodIP: "1.2.3.4",
		},
	}
	req := createRequestObj(false, *pod)
	if req.Host != "localhost:8080" {
		t.Errorf("Host error: %v\n", req.Host)
	}
	err := makeRequest(false, *pod)
	if err != nil {
		t.Errorf("Error in makeRequest: %s\n", err)
	}
}
