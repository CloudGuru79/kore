// +build !ignore_autogenerated

/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKS":                  schema_pkg_apis_eks_v1alpha1_EKS(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSCredentials":       schema_pkg_apis_eks_v1alpha1_EKSCredentials(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSCredentialsSpec":   schema_pkg_apis_eks_v1alpha1_EKSCredentialsSpec(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSCredentialsStatus": schema_pkg_apis_eks_v1alpha1_EKSCredentialsStatus(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSNodeGroup":         schema_pkg_apis_eks_v1alpha1_EKSNodeGroup(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSNodeGroupSpec":     schema_pkg_apis_eks_v1alpha1_EKSNodeGroupSpec(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSNodeGroupStatus":   schema_pkg_apis_eks_v1alpha1_EKSNodeGroupStatus(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSSpec":              schema_pkg_apis_eks_v1alpha1_EKSSpec(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSStatus":            schema_pkg_apis_eks_v1alpha1_EKSStatus(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSVPC":               schema_pkg_apis_eks_v1alpha1_EKSVPC(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSVPCSpec":           schema_pkg_apis_eks_v1alpha1_EKSVPCSpec(ref),
		"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSVPCStatus":         schema_pkg_apis_eks_v1alpha1_EKSVPCStatus(ref),
	}
}

func schema_pkg_apis_eks_v1alpha1_EKS(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKS is the Schema for the eksclusters API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSSpec", "github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSCredentials(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSCredentials is the Schema for the ekscredentials API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSCredentialsSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSCredentialsStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSCredentialsSpec", "github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSCredentialsStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSCredentialsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSCredentialsSpec defines the desired state of EKSCredential",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"accountID": {
						SchemaProps: spec.SchemaProps{
							Description: "AccountID is the AWS account these credentials reside within",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretAccessKey": {
						SchemaProps: spec.SchemaProps{
							Description: "SecretAccessKey is the AWS Secret Access Key",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessKeyID": {
						SchemaProps: spec.SchemaProps{
							Description: "AccessKeyID is the AWS Access Key ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"credentialsRef": {
						SchemaProps: spec.SchemaProps{
							Description: "CredentialsRef is a reference to the credentials used to create clusters",
							Ref:         ref("k8s.io/api/core/v1.SecretReference"),
						},
					},
				},
				Required: []string{"accountID"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretReference"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSCredentialsStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSCredentialsStatus defines the observed state of EKSCredential",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a collection of potential issues",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Condition"),
									},
								},
							},
						},
					},
					"verified": {
						SchemaProps: spec.SchemaProps{
							Description: "Verified checks that the credentials are ok and valid",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status provides a overall status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Condition"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSNodeGroup(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSNodeGroup is the Schema for the eksnodegroups API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSNodeGroupSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSNodeGroupStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSNodeGroupSpec", "github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSNodeGroupStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSNodeGroupSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSNodeGroupSpec defines the desired state of EKSNodeGroup",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"amiType": {
						SchemaProps: spec.SchemaProps{
							Description: "AMIType is the AWS Machine Image type. We use a sensible default.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cluster": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster refers to the cluster this object belongs to",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
					"diskSize": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"instanceType": {
						SchemaProps: spec.SchemaProps{
							Description: "InstanceType is the EC2 machine type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels are any custom kubernetes labels to apply to nodes",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version is the Kubernetes version to run for the kubelet",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"releaseVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "ReleaseVersion is release version of the managed node ami",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"desiredSize": {
						SchemaProps: spec.SchemaProps{
							Description: "DesiredSize is the number of nodes to attempt to use",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"maxSize": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxSize is the most nodes the nodegroups can grow to",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"minSize": {
						SchemaProps: spec.SchemaProps{
							Description: "MinSize is the least nodes the nodegroups can shrink to",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"subnets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Subnets is the VPC networks to use for the nodes",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"tags": {
						SchemaProps: spec.SchemaProps{
							Description: "Tags are the AWS metadata to apply to the node group",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Description: "Region is the AWS location to launch node group within, must match the region of the cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sshSourceSecurityGroups": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "SSHSourceSecurityGroups is the security groups that are allowed SSH access (port 22) to the worker nodes",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"eC2SSHKey": {
						SchemaProps: spec.SchemaProps{
							Description: "EC2SSHKey is the Amazon EC2 SSH key that provides access for SSH communication with the worker nodes in the managed node group https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"amiType", "diskSize", "desiredSize", "maxSize", "minSize", "subnets", "region", "eC2SSHKey"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Ownership"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSNodeGroupStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSNodeGroupStatus defines the observed state of EKSNodeGroup",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is the status of the components",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Component"),
									},
								},
							},
						},
					},
					"nodeIAMRole": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeIAMRole is the IAM role assumed by the worker nodes themselves",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status provides a overall status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Component"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSSpec defines the desired state of EKSCluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"authorizedMasterNetworks": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "AuthorizedMasterNetworks is the network ranges which are permitted to access the EKS control plane endpoint i.e the managed one (not the authentication proxy)",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"cluster": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster refers to the cluster this object belongs to",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version is the Kubernetes version to use",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Description: "Region is the AWS region to launch this cluster within",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"subnetIDs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "SubnetIds is a list of subnet IDs",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"securityGroupIDs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "SecurityGroupIds is a list of security group IDs",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"region", "subnetIDs"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Ownership"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSStatus defines the observed state of EKS cluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is the status of the components",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Component"),
									},
								},
							},
						},
					},
					"caCertificate": {
						SchemaProps: spec.SchemaProps{
							Description: "CACertificate is the certificate for this cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint is the endpoint of the cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"roleARN": {
						SchemaProps: spec.SchemaProps{
							Description: "RoleARN is the role ARN which provides permissions to EKS",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status provides a overall status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Component"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSVPC(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSVPC is the Schema for the eksvpc API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSVPCSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSVPCStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSVPCSpec", "github.com/appvia/kore/pkg/apis/eks/v1alpha1.EKSVPCStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSVPCSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSVPCSpec defines the desired state of EKSVPC",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cluster": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster refers to the cluster this object belongs to",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
					"privateIPV4Cidr": {
						SchemaProps: spec.SchemaProps{
							Description: "PrivateIPV4Cidr is the private range used for the VPC",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Description: "Region is the AWS region of the VPC and any resources created",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"privateIPV4Cidr", "region"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Ownership"},
	}
}

func schema_pkg_apis_eks_v1alpha1_EKSVPCStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EKSVPCStatus defines the observed state of a VPC",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is the status of the components",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Component"),
									},
								},
							},
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status provides a overall status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"infra": {
						SchemaProps: spec.SchemaProps{
							Description: "Infra provides a cache of values discovered from infrastructure k8s:openapi-gen=false",
							Ref:         ref("github.com/appvia/kore/pkg/apis/eks/v1alpha1.Infra"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Component", "github.com/appvia/kore/pkg/apis/eks/v1alpha1.Infra"},
	}
}
