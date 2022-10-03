# Kubernetes Deployment

## Deployment Update
+ Build lại Docker image mới: ```docker build . -t golang-docker-example-v2```
+ Tạo file tên hello-deploy.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-deploy.yaml```
+ Kiểm tra RS đã chạy hay chưa: ```kubectl get rs```
+ Ta test thử deployment type service NodePort với port 31000: ```curl localhost:31000```
+ Update lại ứng dụng trong Pod với Deployment: ```kubectl set image deployment hello-deploy hello-deploy=golang-docker-example-v2```
+ Cấu trúc câu lệnh kubectl set image deployment: ```<deployment-name> <container-name>=<new-image>```
+ Kiểm tra qua trình update đã xong chưa: ```kubectl rollout status deploy hello-deploy```
+ Xóa resources: ```kubectl delete deployment hello-deploy```

## Deployment Rollback
+ Trước tiên bạn có thể kiểm tra lịch sử các lần ứng dụng của chúng ta đã cập nhật: ```kubectl rollout history deploy hello-deploy```
+ Lùi deployment lại version cũ: ```kubectl rollout undo deployment hello-deploy --to-revision=<revision>```
+ Test thử xem có nhận được version cũ hay không: ```curl localhost:31000```
+ Về revision thì default là 10, điều chỉnh bằng tham số revisionHistoryLimit của Deployment
+ Xóa resources: ```kubectl delete deployment hello-deploy```






