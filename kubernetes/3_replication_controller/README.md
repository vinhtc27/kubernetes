# Kubernetes Replication Controller & Other Controllers

## Sử dụng Replication Controller (RC)
+ Tạo file tên là hello-rc.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-rc.yaml```
+ Kiểm tra RC đã chạy hay chưa: ```kubectl get rc```
+ Kiểm tra số lượng pod được tạo ra bởi RC có đúng với số lượng chỉ định ở replicas như lý thuyết hay không: ```kubectl get pod```
+ Tên của pod được tạo ra bởi RC sẽ theo kiểu: <replicationcontroller name>-<random>
+ Xóa thử một thằng pod xem RC có tạo lại một thằng pod khác cho chúng ta như lý thuyết không. Nhớ chỉ định đúng tên pod của bạn: ```kubectl delete pod hello-rc-<random>```
+ Mở cửa sổ terminal khác và gõ câu lệnh: ```kubectl get pod```
=> Có một Pod cũ đang bị xóa đi, và cũng lúc đó, sẽ có một Pod mới được RC tạo ra.
+ Để xóa RC dùng câu lệnh: ```kubectl delete rc hello-rc```

## Sử dụng ReplicaSets (RS) thay thế RC
+ Tạo file tên là hello-rs.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-rs.yaml```
+ Kiểm tra RS đã chạy hay chưa: ```kubectl get rs```
=> Nếu có 2 pod tạo ra là chúng ta đã chạy RS thành công. 
+ Đề xóa RS ta dùng câu lệnh: ```kubectl delete rs hello-rs```