package service

import (
	"github.com/fogray/baasmanager/baas-core/common/httputil"
	"github.com/fogray/baasmanager/baas-core/core/model"
	"github.com/fogray/baasmanager/baas-fabricengine/config"
	"github.com/fogray/baasmanager/baas-fabricengine/constant"
)

type KubeService struct {
	baseFiles         []string //模板文件
	kafkaFiles        []string //模板文件
	baasKubeEngineUrl string
}

func (k KubeService) deployData(datas *model.K8sData) []byte {
	return httputil.PostJson(k.baasKubeEngineUrl+"/deployData", datas)
}

func (k KubeService) deleteData(datas *model.K8sData) []byte {
	return httputil.PostJson(k.baasKubeEngineUrl+"/deleteData", datas)
}

func (k KubeService) getChainDomain(nss string) []byte {
	return httputil.Get(k.baasKubeEngineUrl + "/getChainDomain?namesapces=" + nss)
}

func (k KubeService) getChainPods(nss string) []byte {
	return httputil.Get(k.baasKubeEngineUrl + "/getChainPods?namesapces=" + nss)
}

func (k KubeService) changeDeployResources(datas *model.Resources) []byte {
	return httputil.PostJson(k.baasKubeEngineUrl+"/changeDeployResources", datas)
}

func (k KubeService) getPod(ns string, podName string) []byte {
	return httputil.Get(k.baasKubeEngineUrl + "/getPod?ns=" + ns + "&podName=" + podName)
}
func (k KubeService) printPodLogs(ns string, podName string) string {
	logger.Infof("kube_service printPodLogs: ns=%s, podName=%s", ns, podName)
	bts := httputil.Get(k.baasKubeEngineUrl + "/printPodLogs?ns=" + ns + "&podName=" + podName)
	return string(bts[:])
}

func newKubeService() KubeService {
	return KubeService{
		baseFiles:         []string{constant.K8sNamespaceYaml, constant.K8sNfsYaml, constant.K8sOrdererYaml, constant.K8sPeerYaml, constant.K8sCaYaml},
		kafkaFiles:        []string{constant.K8sZookeeperYaml, constant.K8sKafkaYaml},
		baasKubeEngineUrl: config.Config.GetString("BaasKubeEngine"),
	}
}
