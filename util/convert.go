package util

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"unsafe"
)

func StringToByte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		cap int
	}{s, len(s)}))
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ProjectToNS(project string) string {
	return project + "-workspace"
}

func ProjectToSit(project string) string {
	return project + "-sit"
}

func RawJsonToUnstructured(raw []byte) (*unstructured.Unstructured, error) {
	obj, _, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(raw, nil, nil)
	if err != nil {
		Logger.Errorf("util.RawJsonToUnstructured err: %s", err)
		return nil, err
	}

	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		Logger.Errorf("util.RawJsonToUnstructured err: %s", err)
		return nil, err
	}

	return &unstructured.Unstructured{Object: m}, nil
}
