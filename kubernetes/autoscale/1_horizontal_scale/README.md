# Horizontal Scaling in Kubernetes

## Scale by CPU usage
+ Tạo file tên deployment.yaml (đã có)
+ Tạo file tên hpa-cpu.yaml (đã có)
+ Tạo Deployment: ```kubectl apply -f deployment.yaml```
+ Tạo HPA: ```kubectl apply -f hpa-cpu.yaml```
+ Chạy lệnh đến khi số replicas dc scale down: ```kubectl get hpa,deployment```
+ Mở terminal khác và tạo requests vào pod: ```kubectl run -it --rm --restart=Never request-gen1 --image=busybox -- sh -c "while true; do wget -O - -q http://192.168.1.5:30001; done"```
+ Kiểm tra quá trình scale up: ```kubectl get hpa,deployment``` hoặc ```watch -n 1 kubectl get hpa,deployment``` nếu cài watch
+ Kiểm tra thông tin chi tiết của HPA: ```kubectl describe hpa kubia``` hoặc ```watch -n 1 kubectl describe hpa kubia``` nếu cài watch
+ Xóa resources: ```kubectl delete deployment kubia```
+ Xóa HPA: ```kubectl delete hpa kubia```

## Scale by Memory usage
+ Tạo file tên hpa-cpu.yaml (đã có)

