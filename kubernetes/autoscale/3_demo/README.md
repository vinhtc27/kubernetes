# Kubernetes Autoscaling Demo

## Đề bài

## Tài nguyên cho trước
### Docker image demo-server
+ Tạo file Dockerfile cho file ```demo-server.go``` (đã có)
+ Run câu lệnh build image: ```docker build . -t demo-server```
+ Test thử container: ```docker run -d --name demo-server -p 3000:3000 demo-server```
+ Xóa container: ```docker rm -f demo-server```

### Resource 
+ Kubernetes single node cluster trên Docker desktop
+ Giới hạn phần cứng Docker desktop là 4 core cpu, 4gb ram, 32gb disk.

## Giải pháp

## Từng bước triển khai và config

### Setup metrics server 
+ Chạy lệnh: ```kubectl apply -f demo-metric.yaml```
+ Xác nhận metric server đã chạy: ```kubectl get deployment metrics-server -n kube-system```
+ Kiểm tra metric server: ```kubectl top pod```
### Setup scale bằng HPA theo CPU usage
+ Tạo file tên demo-deployment.yaml (đã có)
+ Tạo file tên demo-hpa.yaml (đã có)
+ Tạo Deployment + Service NodePort: ```kubectl apply -f demo-deployment.yaml```
+ Tạo HPA: ```kubectl apply -f demo-hpa.yaml```
+ Kiểm tra quá trình scale up: ```kubectl get hpa,deployment``` hoặc ```watch -n 1 kubectl get hpa,deployment``` nếu cài watch
+ Chạy file demo-client để tạo ra nhiều requests đến cluster, có thể thay đổi số threadNumber ở client để giả lập lượng request lớn hơn
+ Kiểm tra thông tin chi tiết của HPA: ```kubectl describe hpa demo``` hoặc ```watch -n 1 kubectl describe hpa demo``` nếu cài watch
+ Xóa Deployment: ```kubectl delete deployment demo```
+ Xóa Service: ```kubectl delete service demo```
+ Xóa HPA: ```kubectl delete hpa demo```

### Giải thích quá trình autoscaling

## Kết luận
