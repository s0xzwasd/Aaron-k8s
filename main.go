package main

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func main() {
	metav1.FieldSelectorQueryParam("1")
}
