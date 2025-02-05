package kubeclient

import (
	"bytes"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) GetPodList(ns string, ops metav1.ListOptions) *corev1.PodList {

	pods, err := c.KubeClient.CoreV1().Pods(ns).List(ops)
	if err != nil {
		logger.Errorf(err.Error())
	}
	for _, pod := range pods.Items {
		logger.Infof("Pod：", pod.Name, pod.Status.PodIP)
	}
	return pods
}

func (c *Clients) CreatePod(pod *corev1.Pod) *corev1.Pod {

	newPod, err := c.KubeClient.CoreV1().Pods(pod.Namespace).Create(pod)
	if err != nil {
		logger.Errorf(err.Error())
	}
	logger.Infof("Created pod %q \n", newPod.GetObjectMeta().GetName())
	return newPod
}

func (c *Clients) DeletePod(pod *corev1.Pod, ops *metav1.DeleteOptions) {
	err := c.KubeClient.CoreV1().Pods(pod.Namespace).Delete(pod.Name, ops)
	if err != nil {
		logger.Errorf(err.Error())
	}
	logger.Infof("Delete pod %q \n", pod.GetObjectMeta().GetName())
}

func (c *Clients) GetPod(ns string, podName string) *corev1.Pod {
	logger.Infof("GetPod: ns=%s, podName=%s", ns, podName)
	podInfo, err := c.KubeClient.CoreV1().Pods(ns).Get(podName, metav1.GetOptions{})
	if err != nil {
		logger.Errorf("error in getting Pod")
	}
	return podInfo
}
func (c *Clients) PrintPodLogs(ns string, podName string) string {
	podLogOpts := corev1.PodLogOptions{}
	logger.Infof("PrintPodLogs: ns=%s, podName=%s", ns, podName)
	req := c.KubeClient.CoreV1().Pods(ns).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream()
	//if err != nil {
	//	logger.Errorf("error in opening stream")
	//}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		logger.Errorf("error in copy information from podLogs to buf")
	}
	str := buf.String()

	logger.Infof("Pod loggers :", str)
	return str
}
