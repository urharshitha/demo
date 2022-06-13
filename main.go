package main

import (
	"fmt"
	"strings"
    "k8s.io/kubectl/pkg/cmd/get"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/meta"
)

func createPodSpecResource(memReq, memLimit, cpuReq, cpuLimit string) corev1.PodSpec {
	
	podSpec := corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{},
					Limits:   corev1.ResourceList{},
				},
			},
		},
	}

	req := podSpec.Containers[0].Resources.Requests
	if memReq != "" {
		memReq, err := resource.ParseQuantity(memReq)
		req["memory"] = memReq
	}
	if cpuReq != "" {
		cpuReq, err := resource.ParseQuantity(cpuReq)
		req["cpu"] = cpuReq
	}
	limit := podSpec.Containers[0].Resources.Limits
	if memLimit != "" {
		memLimit, err := resource.ParseQuantity(memLimit)
		limit["memory"] = memLimit
	}
	if cpuLimit != "" {
		cpuLimit, err := resource.ParseQuantity(cpuLimit)
		limit["cpu"] = cpuLimit
	}

	return podSpec
}

func main(){
	type Object interface {
		Live() runtime.Object
		Merged() (runtime.Object, error)
	
		Name() string
	}


    T := []struct {
	obj         runtime.Object
	sort        runtime.Object
	field       string
	name        string
	expectedErr string
    }{
	    {
		    name: "empty",
		    obj: &corev1.PodList{
			    Items: []corev1.Pod{},
		    },
		    sort: &corev1.PodList{
			    Items: []corev1.Pod{},
		    },
		    field: "{.metadata.name}",
	    },
	    {
		    name: "in-order-already",
		    obj: &corev1.PodList{
			    Items: []corev1.Pod{
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "a",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "b",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "c",
					    },
				    },
			    },
		    },
		    sort: &corev1.PodList{
			    Items: []corev1.Pod{
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "a",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "b",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "c",
					    },
				    },
			    },
		    },
		    field: "{.metadata.name}",
	    },
	    {
		    name: "reverse-order",
		    obj: &corev1.PodList{
			    Items: []corev1.Pod{
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "b",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "c",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "a",
					    },
				    },
			    },
		    },
		    sort: &corev1.PodList{
			    Items: []corev1.Pod{
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "a",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "b",
					    },
				    },
				    {
					    ObjectMeta: metav1.ObjectMeta{
						    Name: "c",
					    },
				    },
			    },
		    },
		    field: "{.metadata.name}",
	    },
	

    }
            
            obj:= &corev1.PodList{
	            Items: []corev1.Pod{
		            {
			            Spec: createPodSpecResource("", "", "0.5", ""),
		            },
		            {
			            Spec: createPodSpecResource("", "", "10", ""),
		            },
		            {
			            Spec: createPodSpecResource("", "", "100m", ""),
		            },
		            {
			            Spec: createPodSpecResource("", "", "", ""),
		            },
	                },
            },
            objs, err := meta.ExtractList(obj)
            fieldName := "{.metadata.name}"
	    runtimeSortName := NewRuntimeSort(fieldName, objs)

}
