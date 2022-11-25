/*
Copyright ApeCloud Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apecloud/kubeblocks/internal/dbctl/util/fake"
)

var _ = Describe("cluster util", func() {
	client := fake.NewClientSet(
		fake.Pods(3, fake.Namespace, fake.ClusterName),
		fake.Node(),
		fake.Secrets(fake.Namespace, fake.ClusterName),
		fake.Services())

	dynamic := fake.NewDynamicClient(
		fake.Cluster(fake.ClusterName, fake.Namespace),
		fake.ClusterDef(),
		fake.Appversion())

	It("get cluster objects", func() {
		clusterName := fake.ClusterName
		objs := NewClusterObjects()
		getter := ObjectsGetter{
			ClientSet:      client,
			DynamicClient:  dynamic,
			Name:           clusterName,
			Namespace:      fake.Namespace,
			WithConfigMap:  true,
			WithAppVersion: true,
		}

		Expect(getter.Get(objs)).Should(Succeed())
		Expect(objs.Cluster.Name).Should(Equal(clusterName))
		Expect(objs.ClusterDef.Name).Should(Equal(fake.ClusterDefName))
		Expect(objs.AppVersion.Name).Should(Equal(fake.AppVersionName))

		Expect(len(objs.Pods.Items)).Should(Equal(3))
		Expect(len(objs.Nodes)).Should(Equal(1))
		Expect(len(objs.Secrets.Items)).Should(Equal(1))
		Expect(objs.Services).ShouldNot(BeNil())
	})
})
