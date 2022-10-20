# StatefulSets: deploying replicated stateful applications
+ StatefulSets: một resource sẽ giúp ta deploy một stateful application.

## Hạn chế của việc sử dụng ReplicaSet để tạo replicated stateful app
+ giống Pod trừ tên và id.
+ ta config volume trong Pod template -> tất cả các Pod được replicated đều lưu trữ dữ liệu chung một storage.
==> không thể sử dụng một ReplicaSet rồi set thuộc tính replicas của nó để chạy một ứng dụng distributed data store được.

## Tạo nhiều ReplicaSet chỉ có một Pod mỗi ReplicaSet

## Cung cấp stable identity cho mỗi Pod

## StatefulSets
+ Mỗi Pod được tạo ra bởi StatefulSet được gán với một index, index này được sử dụng để định danh cho mỗi Pod. Tên của Pod được đặt theo kiểu <statefulset name>-<index>, chứ không phải random như của ReplicaSet.

## Cách StatefulSets thay thế một Pod bị mất
+ khi 1 Pod trong 1 StatefulSets bị mất -> StatefulSets sẽ tạo 1 Pod mới(tên và hostname giống pod cũ), ReplicaSet thì tạo ra Pod mới hoàn toàn khác với Pod cũ.

## Cách StatefulSets scale Pod
+ scaleup: tạo ra 1 pod mới có index = index(cũ)+1
+ scale down: xóa Pod với index lớn nhất.

## Cung cấp storage riêng biệt cho mỗi Pod

## Tạo một StatefulSets
+ Tạo file kubia-statefulset.yaml với image luksa/kubia-pet(da co)
+ Tạo StatefulSet: ```kubectl create -f kubia-statefulset.yaml```
