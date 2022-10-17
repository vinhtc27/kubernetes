# Metric server Kubernetes

## Install metric server
+ Chạy lệnh: ```kubectl apply -f components.yaml```
+ Xác nhận metric server đã chạy: ```kubectl get deployment metrics-server -n kube-system```
+ Kiểm tra metric server: ```kubectl top pod```