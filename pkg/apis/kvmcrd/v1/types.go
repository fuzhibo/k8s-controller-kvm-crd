package v1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KVM describes a ecs resource
type KVM struct{
    // TypeMeta is the metadata for the resource, like kind and apiversion
    metav1.TypeMeta `json:",inline"`
    // ObjectMeta contains the metadata for the particular object, including
    // things like...
    // - name
    // - namespace
    // - self link
    // - labels
    // - ... etc ...
    metav1.ObjectMeta `json:"metadata,omitempty"`

    // Spec is the custom resource spec
    Spec KVMSpec `json:"spec"`
    Status KVMStatus `json:"status"`
}

// KVMSpec is the spec for a KVM resource
type KVMSpec struct {
    // processes, name, command and args are example custom spec fields
    //
    // this is where you would put your custom resource data
    Processes []KVMProcess `json:"process"`
}

// KVMProcess is the spec for one process will run in KVM virtual machine
type KVMProcess struct {
    // name, command and args are example custom spec fields
    //
    // this is where you would put your custom resource data
    Name string `json:"name"`
    Command string `json:"command"`
    Args []string `json:"args"`
}

type KVMList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata"`

    Items []KVM `json:"items"`
}

