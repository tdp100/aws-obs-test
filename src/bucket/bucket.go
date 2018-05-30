package bucket

type bucketInfo struct {
	//Name  桶名
	Name   string
	//Region 所属区域
	Region string
	//Size 桶的存储占用大小
	Size int64
	//ObjectNumb 桶内对象数
	ObjectNumb int64
	//FragmentNumb 桶碎片数
	FragmentNumb int64
}


