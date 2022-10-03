# Kubernetes Services

## Sử dụng ClusterIP
+ Tạo file tên hello-service.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-service.yaml```
+ Kiểm tra service đã chạy hay chưa: ```kubectl get svc```
+ Ta có thể connect redis với địa chỉ sau redis://redis:6379 với host name là tên của Service chúng ta đặt trong trường metadata
+ Test thử service: ```kubectl run hello-redis --image=080196/hello-redis```
+ Kiểm tra log của pod, nếu in ra Connect redis success thì chúng ta đã kết nối được với redis host bằng dns: kubectl logs hello-redis
+ Xóa resources: ```kubectl delete pod hello-redis```

## Sử dụng Nodeport
+ Tạo file tên hello-nodeport.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-nodeport.yaml```
+ Kiểm tra service đã chạy hay chưa: ```kubectl get svc```
+ Test gửi request tới Pod với địa chỉ: ```http://<worker_node_ip>:<node_port>```
+ Xóa resources: ```kubectl delete rs hello-rs```

## Sử dụng Load Balancer
+ Chỉ dùng được khi chạy Kubernetes cloud

## Sử dụng Ingress Resource
+ Gán 1 domain với service trong cluster
