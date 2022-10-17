# Kubernetes Volume

## Volumn emptyDir
+ Tạo file tên emptydir.yaml (đã có)
+ Tạo pod: ```kubectl apply -f emptydir.yaml```
+ Kiểm tra Pod đã chạy hay chưa: ```kubectl get pod```
+ Chạy câu lệnh để expose port của pod: ```kubectl port-forward fortune 8080:80```
+ Ta test thử: ```curl http://localhost:8080```
+ Xóa resources: ```kubectl delete pod fortune```


## Volumn gitRepo
+ Tạo file tên gitrepo.yaml (đã có)
+ Tạo pod: ```kubectl apply -f gitrepo.yaml```
+ Kiểm tra Pod đã chạy hay chưa: ```kubectl get pod```
+ Chạy câu lệnh để expose port của pod: ```kubectl port-forward gitrepo-volume-pod 8080:80```
+ Ta test thử: ```curl http://localhost:8080```
+ Xóa resources: ```kubectl delete pod gitrepo-volume-pod```

## Volumn hostPath
+ Tạo file tên hostpath.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hostpath.yaml```
+ Kiểm tra Pod đã chạy hay chưa: ```kubectl get pod```
+ Chạy câu lệnh để expose port của pod: ```kubectl port-forward hostpath-volume 8080:80```
+ Ta test thử: ```curl http://localhost:8080```
+ Xóa resources: ```kubectl delete pod hostpath-volume```







