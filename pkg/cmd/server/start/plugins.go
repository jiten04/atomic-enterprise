package start

import (

	// Admission control plug-ins used by OpenShift
	_ "github.com/GoogleCloudPlatform/kubernetes/plugin/pkg/admission/admit"
	_ "github.com/GoogleCloudPlatform/kubernetes/plugin/pkg/admission/limitranger"
	_ "github.com/GoogleCloudPlatform/kubernetes/plugin/pkg/admission/namespace/exists"
	_ "github.com/GoogleCloudPlatform/kubernetes/plugin/pkg/admission/namespace/lifecycle"
	_ "github.com/GoogleCloudPlatform/kubernetes/plugin/pkg/admission/resourcequota"
	_ "github.com/GoogleCloudPlatform/kubernetes/plugin/pkg/admission/serviceaccount"
	_ "github.com/projectatomic/atomic-enterprise/pkg/build/admission"
	_ "github.com/projectatomic/atomic-enterprise/pkg/project/admission/lifecycle"
	_ "github.com/projectatomic/atomic-enterprise/pkg/project/admission/nodeenv"
	_ "github.com/projectatomic/atomic-enterprise/pkg/security/admission"
)
