// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package components

var (
	
	ComponentAPIServer = component{
		Version: "v2.7.0",
		Digest:  "sha256:33755103ca321c07e07e1b1302428c3a6d29bc40f0b71519be879a0a689868fd",
		Image:   "tigera/cnx-apiserver",
	}
	
	
	ComponentComplianceBenchmarker = component{
		Version: "v2.7.0",
		Digest:  "sha256:59886e14f8af2a9335df1666e65f6801a4b30f076872f3dad6cef14dd35e1838",
		Image:   "tigera/compliance-benchmarker",
	}
	
	
	ComponentComplianceController = component{
		Version: "v2.7.0",
		Digest:  "sha256:a18b73d69c6edef6c0c87e0b85dfd4f8977293f24d899c6c369be914106b1b93",
		Image:   "tigera/compliance-controller",
	}
	
	
	ComponentComplianceReporter = component{
		Version: "v2.7.0",
		Digest:  "sha256:030112b0dbd9d66fd557f6c118e8001ff8569661bcc6503ed29891281e0e0695",
		Image:   "tigera/compliance-reporter",
	}
	
	
	ComponentComplianceServer = component{
		Version: "v2.7.0",
		Digest:  "sha256:421c4ffe4b04bf91d4ae4840d65c3d42aa7edfa3e0fa6a222a7cb8297cc05465",
		Image:   "tigera/compliance-server",
	}
	
	
	ComponentComplianceSnapshotter = component{
		Version: "v2.7.0",
		Digest:  "sha256:b2a7ad0c8df9a7f51a9bf8865bdc357682309fa3ab0247092915304115c5be86",
		Image:   "tigera/compliance-snapshotter",
	}
	
	
	ComponentEckKibana = component{
		Version: "7.3.2",
		Digest:  "sha256:fbd81abc89aab14b99b463b32310052e4ea563870cd0eadb1dd1153e54769ed3",
		Image:   "tigera/kibana",
	}
	
	
	ComponentElasticTseeInstaller = component{
		Version: "v2.7.0-0.dev-38-gbf021d6",
		Digest:  "sha256:3fded7f9999c7b4547044da41c857e30e061b0680942efdd007dbd5214ab574b",
		Image:   "tigera/intrusion-detection-job-installer",
	}
	
	
	ComponentElasticsearch = component{
		Version: "7.3.2",
		Digest:  "sha256:9dd701842833e902c0ba2b5682e0b7b3cd32f08e94a637d87ebdcb730e0a101a",
		Image:   "elasticsearch/elasticsearch",
	}
	
	
	ComponentElasticsearchOperator = component{
		Version: "0.9.0",
		Digest:  "sha256:9f134a7647371fc4e627c592496ff1ef1c3d50d6206fa6fb3571aacdb6ade574",
		Image:   "eck/eck-operator",
	}
	
	
	ComponentEsCurator = component{
		Version: "v2.7.0",
		Digest:  "sha256:34a8486c7e0d646b17c788adb4bafd41d567a85be5c75f4bfe1d650ebc85d475",
		Image:   "tigera/es-curator",
	}
	
	
	ComponentEsProxy = component{
		Version: "v2.7.0",
		Digest:  "sha256:9da0f6f55431b9d6737bd915083cf51e09fe537b2072681e8b88a0cee5695259",
		Image:   "tigera/es-proxy",
	}
	
	
	ComponentFluentd = component{
		Version: "v2.7.0",
		Digest:  "sha256:89ab6e0e35d113f3a000916b4f03f0b811504cdc701f7c92d8d83f054b1f3bcc",
		Image:   "tigera/fluentd",
	}
	
	
	ComponentGuardian = component{
		Version: "v2.7.0",
		Digest:  "sha256:80db830bbac524904d6837ebdf66dadfe337c7ee26aabcf746fef1b7d82cb749",
		Image:   "tigera/guardian",
	}
	
	
	ComponentIntrusionDetectionController = component{
		Version: "v2.7.0-0.dev-38-gbf021d6",
		Digest:  "sha256:d91bc2a82527c14306e50c4682cf3ff681914c067547df5c500cb5760f5413d6",
		Image:   "tigera/intrusion-detection-controller",
	}
	
	
	ComponentKibana = component{
		Version: "7.3",
		Digest:  "sha256:4c8e458fa0327c477c9aabd3672cc294e4497173faa61c475c76c47748677bf0",
		Image:   "tigera/kibana",
	}
	
	
	ComponentManager = component{
		Version: "v2.7.0",
		Digest:  "sha256:439666b2f313341330ccece00cceff97b0c9cd77b4d834470060c519e51fccfc",
		Image:   "tigera/cnx-manager",
	}
	
	
	ComponentManagerProxy = component{
		Version: "v2.7.0",
		Digest:  "sha256:cd4540db662573a9f0383feff7686555f188233802476546e0b523104802e689",
		Image:   "tigera/voltron",
	}
	
	
	ComponentQueryServer = component{
		Version: "v2.7.0",
		Digest:  "sha256:3603e45257693d8112fee4244ed5ceeb09c6dd34df07eb5ddb1b2e882eb209ae",
		Image:   "tigera/cnx-queryserver",
	}
	
	
	ComponentTigeraKubeControllers = component{
		Version: "v2.7.0-0.dev-128-g3f4d4f5",
		Digest:  "sha256:489115ecd0cdd249a7d321f79fa5bbfe8df2de5b07b2ce11f25337283490d4a1",
		Image:   "tigera/kube-controllers",
	}
	
	
	ComponentTigeraNode = component{
		Version: "v2.7.0",
		Digest:  "sha256:205c5e15543ef4c0f166a77326a061f5d99a1e20f6c9e4a0a8227b620333d280",
		Image:   "tigera/cnx-node",
	}
	
	
	ComponentTigeraTypha = component{
		Version: "v2.7.0",
		Digest:  "sha256:262d9e73ac8842eeaa6219a1c1453295cd55124b5ded5ade2cc765b03a1006e7",
		Image:   "tigera/typha",
	}
	
)
