# Downward API
+ truyền metadata của pod vào container

## Truyền metadata bằng env
+ Tạo file  downward-api-env.yaml
+ Tạo Pod và kiếm tra thử:
```kubectl apply -f downward-api-env.yaml```
```kubectl exec downward main -- env```

## Truyền metadata bằng volume file
+ Tạo file downward-api-volume.yaml
+ Tạo Pod và kiếm tra thử:
```kubectl apply -f downward-api-volume.yaml```
```kubectl exec downward-volume -- ls -lL /etc/downward```

## Kubernetes API server
+ check: ```kubectl cluster-info```
+ gửi request tới API server: ```curl https://kubernetes.docker.internal:6443```
+ truyen thêm params --insecure (hoặc -k): ```curl -k https://kubernetes.docker.internal:6443```
+ test: ```kubectl proxy```
+ terminal khac: ```curl 127.0.0.1:8001```

## Tương tác với API server
+  ```curl 127.0.0.1:8001/api/v1```
hoac: ```curl curl 127.0.0.1:8001/api/v1/namespaces/default/pods```
+ Cấu trúc dường dẫn của API <api-server-url>/api/v1/namespaces/<namespace-name>/pods
+ lấy thông tin của 1 Pod: ```curl 127.0.0.1:8001/api/v1/namespaces/gitlab/pods/gitlab-webservice-defaul-ff4459cf5-bqklw```
+  cấu trúc: <api-server-url>/api/v1/namespaces/<namespace-name>/pods/<pod-name>

## Tương tác API server bên trong container của Pod
+ ```kubectl get svc```
+ Ta có một service tên là kubernetes, bên trong container, ta có thể gọi tới API server bằng URL https://kubernetes. tạo một pod và truy cập để test thử: ```kubectl run curl --image=curlimages/curl --command -- sleep 9999999```
```kubectl exec -it curl -- sh``` , ```curl https://kubernetes```, 
+ truy cập đc vào pod, gửi request tới API server( làm theo hướng dẫn trên vibo)


